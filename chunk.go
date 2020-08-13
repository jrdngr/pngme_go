package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"hash/crc32"
)

// Chunk for a PNG file
type Chunk struct {
	length    uint32
	chunkType ChunkType
	data      []byte
	crc       uint32
}

// CreateChunkFromBytes does what it says it does
func CreateChunkFromBytes(chunkBytes []byte) Chunk {
	reader := bytes.NewReader(chunkBytes)
	fourBytes := make([]byte, 4)

	reader.Read(fourBytes)
	length := binary.BigEndian.Uint32(fourBytes)

	reader.Read(fourBytes)
	chunkTypeBytes := [4]byte{}
	copy(chunkTypeBytes[:], fourBytes[:])
	chunkType := CreateChunkTypeFromBytes(chunkTypeBytes)

	data := make([]byte, length)
	reader.Read(data)

	reader.Read(fourBytes)
	crc := binary.BigEndian.Uint32(fourBytes)

	var checksumBytes []byte
	checksumBytes = append(checksumBytes, chunkTypeBytes[:]...)
	checksumBytes = append(checksumBytes, data...)
	expectedCrc := crc32.ChecksumIEEE(checksumBytes)

	if crc != expectedCrc {
		fmt.Println("Calculated crc did not match data crc")
		fmt.Printf("Calculated crc: %v\n", expectedCrc)
		fmt.Printf("Data crc: %v\n", crc)
		panic("Invalid input data")
	}

	return Chunk{
		length,
		chunkType,
		data,
		crc,
	}
}

func (chunk Chunk) String() string {
	length := fmt.Sprintf("Length: %v", chunk.length)
	chunkType := fmt.Sprintf("Chunk Type: %v", chunk.chunkType)
	data := fmt.Sprintf("Data: %v bytes", len(chunk.data))
	crc := fmt.Sprintf("Crc: %v", chunk.crc)

	return fmt.Sprintf("%v\n%v\n%v\n%v", length, chunkType, data, crc)
}
