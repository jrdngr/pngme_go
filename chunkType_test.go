package main

import "testing"

func TestChunkTypeFromBytes(t *testing.T) {
	expected := [4]byte{71, 111, 89, 111};
	actualBytes := [4]byte{71, 111, 89, 111}
	actual := CreateChunkTypeFromBytes(actualBytes)

	if expected != actual.Bytes() {
		t.Errorf("CreateChunkTypeFromBytes(%v) = %v. Expected: %v", actualBytes, actual.Bytes(), expected)
	}
}

func TestChunkTypeFromString(t *testing.T) {
	expected := [4]byte{71, 111, 89, 111};
	actual := CreateChunkTypeFromString("GoYo")

	if expected != actual.Bytes() {
		t.Errorf("CreateChunkTypeFromString(%v) = %v. Expected: %v", "GoYo", actual.Bytes(), expected)
	}
}

func TestChunkTypeIsCritical(t *testing.T) {
	chunk := CreateChunkTypeFromString("GoYo")
	if !chunk.IsCritical() {
		t.Errorf("%v.IsCritical() should be true", chunk)
	}
}

func TestChunkTypeIsNotCritical(t *testing.T) {
	chunk := CreateChunkTypeFromString("goYo")
	if chunk.IsCritical() {
		t.Errorf("%v.IsCritical() should be false", chunk)
	}
}

func TestChunkTypeIsPublic(t *testing.T) {
	chunk := CreateChunkTypeFromString("GOYo")
	if !chunk.IsPublic() {
		t.Errorf("%v.IsPublic() should be true", chunk)
	}
}

func TestChunkTypeIsPrivate(t *testing.T) {
	chunk := CreateChunkTypeFromString("GoYo")
	if chunk.IsPublic() {
		t.Errorf("%v.IsPublic() should be false", chunk)
	}
}

func TestChunkTypeIsReservedBitValid(t *testing.T) {
	chunk := CreateChunkTypeFromString("GoYo")
	if !chunk.IsReservedBitValid() {
		t.Errorf("%v.IsReservedBitValid() should be true", chunk)
	}
}

func TestChunkTypeIsReservedBitInvalid(t *testing.T) {
	defer func() { recover() }()

	chunk := CreateChunkTypeFromString("Goyo")
	
	t.Errorf("%v did not panic", chunk)
}

func TestChunkTypeIsSafeToCopy(t *testing.T) {
	chunk := CreateChunkTypeFromString("GoYo")
	if !chunk.IsSafeToCopy() {
		t.Errorf("%v.IsSafeToCopy() should be true", chunk)
	}
}

func TestChunkTypeIsUnsafeToCopy(t *testing.T) {
	chunk := CreateChunkTypeFromString("GoYO")
	if chunk.IsSafeToCopy() {
		t.Errorf("%v.IsSafeToCopy() should be false", chunk)
	}
}

func TestValidChunkIsValid(t *testing.T) {
	chunk := CreateChunkTypeFromString("GoYo")
	if !chunk.IsValid() {
		t.Errorf("%v.IsValid() should be true", chunk)
	}
}

func TestInvalidChunkIsValid(t *testing.T) {
	defer func() { 
		recover() 

		defer func() { recover() }()
		badChunk := CreateChunkTypeFromString("Goyo")

		t.Errorf("%v did not panic", badChunk)
	}()

	chunk := CreateChunkTypeFromString("Goyo")

	t.Errorf("%v did not panic", chunk)
}

func TestChunkTypeString(t *testing.T) {
	chunk := CreateChunkTypeFromString("GoYo")
	if chunk.String() != "GoYo" {
		t.Errorf("Expected chunk string: %v. Got: %v", "GoYo", chunk.String())
	}
}
