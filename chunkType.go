package main

import (
	"fmt"
	"unicode"
)

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
		panic(fmt.Sprintf("String must have exactly 4 characters: %v", str))
	}


	strBytes := []byte(str)

	for _, b := range strBytes {
		if b < 'A' || (b > 'Z' && b < 'a') || b > 'z' {
			panic(fmt.Sprintf("String must only contain the letters A-Z and a-z: %v", rune(b)))
		}
	}

	var bytes [4]byte
	copy(bytes[:], strBytes[0:4])

	chunkType := ChunkType { bytes }

	if !chunkType.IsReservedBitValid() {
		panic(fmt.Sprintf("Reserved bit must be uppercase: %v", str))
	}

	return chunkType
}

func (ct ChunkType) String() string {
	return string(ct.bytes[:])
}

// Equals checks if two ChunkTypes are equal
func (ct *ChunkType) Equals(other *ChunkType) bool {
	i := 0
	for i < len(ct.bytes) {
		if ct.bytes[i] != other.bytes[i] {
			return false
		}
	}

	return true
}

// Bytes returns the ChunkType's bytes
func (ct *ChunkType) Bytes() [4]byte {
	return ct.bytes
}

// IsValid does stuff
func (ct *ChunkType) IsValid() bool {
	return ct.IsReservedBitValid()
}

// IsCritical does stuff
func (ct *ChunkType) IsCritical() bool {
	return unicode.IsUpper(rune(ct.bytes[0]))
}

// IsPublic does stuff
func (ct *ChunkType) IsPublic() bool {
	return unicode.IsUpper(rune(ct.bytes[1]))
}

// IsReservedBitValid does stuff
func (ct *ChunkType) IsReservedBitValid() bool {
	return unicode.IsUpper(rune(ct.bytes[2]))
}

// IsSafeToCopy does stuff
func (ct *ChunkType) IsSafeToCopy() bool {
	return unicode.IsLower(rune(ct.bytes[3]))
}
