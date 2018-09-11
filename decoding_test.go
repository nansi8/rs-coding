package rs_coding

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestDecode(t *testing.T) {
	const dataBlocks, checksumBlocks, degree = 6, 4, 8
	decoder := NewDecoder(dataBlocks, checksumBlocks, degree)
	encoder := NewEncoder(dataBlocks, checksumBlocks, degree)

	input := []byte{15, 199, 187, 129, 134, 57, 172, 72, 164, 198}
	output := encoder.Encode(input)
	for i := 0; i < checksumBlocks; i++ {
		removeIndex := rand.Intn(len(output))
		output = append(output[:removeIndex], output[removeIndex+1:]...)
	}
	decode, err := decoder.Decode(output)
	if err != nil {
		t.Fatal(err)
	}
	for i := range input {
		if input[i] != decode[i] {
			t.Errorf("Wrong decoding at index %d. Expected value: %d. Actual value %d", i, input[i], decode[i])
		}
	}
}

func TestString(t *testing.T) {
	const dataBlocks, checksumBlocks, degree = 6, 4, 8
	decoder := NewDecoder(dataBlocks, checksumBlocks, degree)
	encoder := NewEncoder(dataBlocks, checksumBlocks, degree)

	input := []byte("Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum.")
	output := encoder.Encode(input)
	for i := 0; i < checksumBlocks; i++ {
		removeIndex := rand.Intn(len(output))
		output = append(output[:removeIndex], output[removeIndex+1:]...)
	}
	decode, err := decoder.Decode(output)
	if err != nil {
		t.Fatal(err)
	}
	for i := range input {
		if input[i] != decode[i] {
			t.Errorf("Wrong decoding at index %d. Expected value: %d. Actual value %d", i, input[i], decode[i])
		}
	}
}

func TestFailed(t *testing.T) {
	const dataBlocks, checksumBlocks, degree = 6, 4, 8
	decoder := NewDecoder(dataBlocks, checksumBlocks, degree)
	encoder := NewEncoder(dataBlocks, checksumBlocks, degree)

	input := []byte{15, 199, 187, 129, 134, 57, 172, 72, 164, 198}
	output := encoder.Encode(input)
	for i := 0; i < checksumBlocks+2; i++ {
		removeIndex := rand.Intn(len(output))
		output = append(output[:removeIndex], output[removeIndex+1:]...)
	}
	_, err := decoder.Decode(output)
	if err == nil {
		t.Error("Error must be thrown")
	}
	expected := fmt.Sprintf("Can not decode as %d blocks required, but only %d are present", dataBlocks, checksumBlocks-2)
	if expected != err.Error() {
		t.Errorf("Wrong error message")
		t.Errorf("Expected: %s", expected)
		t.Errorf("Actual: %s", err.Error())
	}
}
