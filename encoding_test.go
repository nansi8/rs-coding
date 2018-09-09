package rs_coding

import (
	"testing"
)

func TestEncode(t *testing.T) {
	encoder := NewEncoder(3, 3, 4)
	input := []byte{7, 3, 2}
	output := encoder.Encode(input)
	expectedBlockData := []byte{7, 3, 2, 6, 7, 1}
	expectedBlockTypes := []BlockType{Data, Data, Data, Checksum, Checksum, Checksum}
	for i := range output {
		if output[i].w[0] != expectedBlockData[i] || output[i].blockType != expectedBlockTypes[i] {
			t.Errorf("Block [%d] is wrong. Expected value [%d]. Actual value [%d]", i, expectedBlockData[i], output[i].w[0])
			t.Errorf("Block [%d] is wrong. Expected type [%d]. Actual type [%d]", i, expectedBlockTypes[i], output[i].blockType)
		}
	}
}

func TestEncodeGF256(t *testing.T) {
	encoder := NewEncoder(6, 4, 8)
	input := []byte{78, 92, 94, 21, 12, 36}
	output := encoder.Encode(input)
	expectedBlockData := []byte{78, 92, 94, 21, 12, 36, 113, 164, 115, 109}
	expectedBlockTypes := []BlockType{Data, Data, Data, Data, Data, Data, Checksum, Checksum, Checksum, Checksum}
	for i := range output {
		if output[i].w[0] != expectedBlockData[i] || output[i].blockType != expectedBlockTypes[i] || output[i].index != i {
			t.Errorf("Block [%d] is wrong. Expected value [%d]. Actual value [%d]", i, expectedBlockData[i], output[i].w[0])
			t.Errorf("Block [%d] is wrong. Expected type [%d]. Actual type [%d]", i, expectedBlockTypes[i], output[i].blockType)
			t.Errorf("Block [%d] is wrong. Expected index [%d]. Actual index [%d]", i, i, output[i].index)
		}
	}
}

func TestEncode2(t *testing.T) {
	encoder := NewEncoder(3, 3, 4)
	input := []byte{7, 3, 2, 7, 3, 2}
	output := encoder.Encode(input)
	expectedBlockData := []byte{7, 3, 2, 6, 7, 1}
	expectedBlockTypes := []BlockType{Data, Data, Data, Checksum, Checksum, Checksum}
	for i := range output {
		if output[i].w[0] != expectedBlockData[i] || output[i].blockType != expectedBlockTypes[i] {
			t.Errorf("Block [%d] is wrong. Expected value [%d]. Actual value [%d]", i, expectedBlockData[i], output[i].w[0])
			t.Errorf("Block [%d] is wrong. Expected type [%d]. Actual type [%d]", i, expectedBlockTypes[i], output[i].blockType)
		}
		if output[i].w[1] != expectedBlockData[i] || output[i].blockType != expectedBlockTypes[i] {
			t.Errorf("Block [%d] is wrong. Expected value [%d]. Actual value [%d]", i, expectedBlockData[i], output[i].w[1])
			t.Errorf("Block [%d] is wrong. Expected type [%d]. Actual type [%d]", i, expectedBlockTypes[i], output[i].blockType)
		}
	}
}
