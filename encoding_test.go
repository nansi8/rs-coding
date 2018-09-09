package rs_coding

import (
	"github.com/nansi8/math"
	"testing"
)

var defaultByteAlgebra = new(math.ByteAlgebraImpl)
var galoisByteAlgebra = math.New(4)

func TestVandermore(t *testing.T) {
	vandermore := vandermore(3, 3, galoisByteAlgebra)
	if vandermore[0][0] != 1 || vandermore[0][1] != 1 || vandermore[0][2] != 1 ||
		vandermore[1][0] != 1 || vandermore[1][1] != 2 || vandermore[1][2] != 3 ||
		vandermore[2][0] != 1 || vandermore[2][1] != 4 || vandermore[2][2] != 5 {
		t.Error("Vandermore matrix is not generated correctly")
	}
}

func TestPow(t *testing.T) {
	var a, x byte = 2, 4
	result := pow(a, x, defaultByteAlgebra)
	if result != 16 {
		t.Errorf("2^4 must be equal to 16 in default algebra but it is %d", result)
	}
}

func TestPowGalois(t *testing.T) {
	var a, x byte = 3, 2
	result := pow(a, x, galoisByteAlgebra)
	if result != 5 {
		t.Errorf("2^4 must be equal to 5 in galois algebra but it is %d", result)
	}
}

func TestExtend(t *testing.T) {
	const length = 10
	data := []byte{0, 1, 2, 3, 4, 5}
	result := extend(data, length)
	if len(result) != length {
		t.Errorf("Length of extended slice must be %d but it is %d", length, len(result))
	}
	for i := range data {
		if data[i] != result[i] {
			t.Errorf("Values of slices in index %d are different. Original value %d. Copied value %d", i, data[i], result[i])
		}
	}
	for i := len(data); i < len(result); i++ {
		if result[i] != 0 {
			t.Errorf("Elements of extended slice must be equal to 0, but it is %d", result[i])
		}
	}
}

func TestExtendSmall(t *testing.T) {
	data := []byte{0, 1, 2, 3, 4, 5}
	result := extend(data, 1)
	if len(result) != len(data) {
		t.Errorf("Result length must be equal to original length, but it is %d", len(result))
	}
	for i := range data {
		if data[i] != result[i] {
			t.Errorf("Values of slices in index %d are different. Original value %d. Copied value %d", i, data[i], result[i])
		}
	}
}

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
