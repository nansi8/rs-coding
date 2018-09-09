package rs_coding

import "github.com/nansi8/math"

type BlockType byte
type BlockId int64

const (
	Data BlockType = iota
	Checksum
)

type Block struct {
	index     int
	w         []byte
	blockType BlockType
}

func getDataBlockMatrix(data []byte) [][]byte {
	result := make([][]byte, len(data))
	for i := 0; i < len(data); i++ {
		result[i] = make([]byte, 1)
		result[i][0] = data[i]
	}
	return result
}

func getCheckBlock(data [][]byte) []byte {
	result := make([]byte, len(data))
	for i := 0; i < len(data); i++ {
		result[i] = data[i][0]
	}
	return result
}

func extend(data []byte, length int) []byte {
	if length < len(data) {
		return data
	}
	result := make([]byte, length)
	copy(result, data)
	return result
}

func vandermore(rows, cols int, alg math.ByteAlgebra) [][]byte {
	result := make([][]byte, rows)
	for i := 0; i < rows; i++ {
		result[i] = make([]byte, cols)
		for j := 0; j < cols; j++ {
			result[i][j] = pow(byte(j+1), byte(i), alg)
		}
	}
	return result
}

// returns a^x
func pow(a, x byte, alg math.ByteAlgebra) byte {
	if x == 0 {
		return 1
	}
	if x == 1 {
		return a
	}
	return alg.Mul(a, pow(a, x-1, alg))
}
