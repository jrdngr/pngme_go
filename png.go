package main

import (
	"bytes"
	"encoding/binary"
	"io"
)

// Png is a PNG
type Png struct {
	signature [8]byte
	chunks []Chunk
}

// CreatePngFromBytes does what it says
func CreatePngFromBytes(pngBytes []byte) Png {
	reader := bytes.NewReader(pngBytes)
	
	signatureBuffer := make([]byte, 8)
	reader.Read(signatureBuffer)

	if !isSignatureValid(signatureBuffer) {
		panic("8 byte signature did not match PNG spec")
	}

	signature := [8]byte{}
	copy(signature[:], signatureBuffer[:])

	chunks := make([]Chunk, 0)

	lengthBuffer := make([]byte, 4)

	for {
		_, err := reader.Read(lengthBuffer)
		if err == io.EOF {
			break
		}

		length := binary.LittleEndian.Uint32(lengthBuffer)

		chunkBuffer := make([]byte, length + 8)
		reader.Read(chunkBuffer)

		
		data := make([]byte, length + 12)
		data = append(data, lengthBuffer...)
		data = append(data, chunkBuffer...)

		chunk := CreateChunkFromBytes(data)
		chunks = append(chunks, chunk)
	}

	return Png {
		signature,
		chunks,
	}
}

func isSignatureValid(signature []byte) bool {
	expectedSignature := []byte{ 137, 80, 78, 71, 13, 10, 26, 10 }

	for i, b := range expectedSignature {
		if signature[i] != b {
			return false
		}
	} 

	return true
}
