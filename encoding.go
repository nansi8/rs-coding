package rs_coding

import (
	"github.com/nansi8/math"
	gomath "math"
)

type BlockType byte
type BlockId int64

const (
	Data BlockType = iota
	Checksum
)

type Block struct {
	w         []byte
	blockType BlockType
}

type Encoder struct {
	blocks         []Block
	dataBlocks     int
	checksumBlocks int
	degree         byte
}

func NewEncoder(dataBlocks, checksumBlocks int, degree byte) *Encoder {
	encoder := new(Encoder)
	encoder.blocks = make([]Block, dataBlocks+checksumBlocks)
	encoder.dataBlocks = dataBlocks
	encoder.checksumBlocks = checksumBlocks
	encoder.degree = degree
	return encoder
}

func (e *Encoder) Encode(input []byte) []Block {
	blocksNumber := int(gomath.Ceil(float64(len(input)) / float64(e.dataBlocks)))

	data := extend(input, blocksNumber*e.dataBlocks)
	galoisAlgebra := math.New(e.degree)
	vandermore := vandermore(e.checksumBlocks, e.dataBlocks, galoisAlgebra)
	resultBlocks := make([][]byte, blocksNumber)
	blocks := make([]Block, e.dataBlocks+e.checksumBlocks)
	for i := 0; i < blocksNumber; i++ {
		dataBlock := data[i*e.dataBlocks : (i+1)*e.dataBlocks]
		mul := math.Mul(vandermore, getDataBlockMatrix(dataBlock), galoisAlgebra)
		checkBlock := getCheckBlock(mul)
		resultBlocks[i] = append(resultBlocks[i], dataBlock...)
		resultBlocks[i] = append(resultBlocks[i], checkBlock...)
	}
	for i := 0; i < e.dataBlocks+e.checksumBlocks; i++ {
		block := new(Block)
		blockData := make([]byte, blocksNumber)
		if i < e.dataBlocks {
			block.blockType = Data
		} else {
			block.blockType = Checksum
		}
		for j := 0; j < len(resultBlocks); j++ {
			blockData[j] = resultBlocks[j][i]
		}
		block.w = blockData
		blocks[i] = *block
	}
	return blocks
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
