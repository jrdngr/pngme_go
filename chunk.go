package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"hash/crc32"
)

// Chunk for a PNG file
type Chunk struct {
	length uint32
	chunkType ChunkType
	data []byte
	crc uint32
}

// CreateChunkFromBytes does what it says it does
func CreateChunkFromBytes(chunkBytes []byte) Chunk {
	reader := bytes.NewReader(chunkBytes)
	fourBytes := make([]byte, 4)

	reader.Read(fourBytes)
	length := binary.LittleEndian.Uint32(fourBytes)

	reader.Read(fourBytes)
	chunkTypeBytes := [4]byte{}
	copy(chunkTypeBytes[:], fourBytes[:])
	chunkType := CreateChunkTypeFromBytes(chunkTypeBytes)
	
	data := make([]byte, length)
	reader.Read(data)

	reader.Read(fourBytes)
	crc := binary.LittleEndian.Uint32(fourBytes)

	checksumBytes := make([]byte, len(data) + int(length))
	checksumBytes = append(checksumBytes, chunkTypeBytes[:]...)
	checksumBytes = append(checksumBytes, data...)
	expectedCrc := crc32.ChecksumIEEE(checksumBytes);

	if crc != expectedCrc {
		panic("Calculated crc did not match data crc")
	}

	return Chunk {
		length,
		chunkType,
		data,
		crc,
	}
}

func (chunk Chunk) String() string {
	return fmt.Sprintf("%v", chunk.chunkType)
}
