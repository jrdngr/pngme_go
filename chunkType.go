package main

// ChunkType for a PNG file
type ChunkType struct {
	bytes [4]byte
}

// CreateChunkTypeFromBytes creates a ChunkType from 4 bytes
func CreateChunkTypeFromBytes(bytes [4]byte) ChunkType {
	return ChunkType { bytes }
}

// CreateChunkTypeFromString creates a ChunkType from a string of 4 characters
func CreateChunkTypeFromString(str string) ChunkType {
	if len(str) != 4 {
		panic("String must have exactly 4 characters")
	}

	strBytes := []byte(str)
	var bytes [4]byte
	copy(bytes[:], strBytes[0:4])

	return ChunkType { bytes }
}

