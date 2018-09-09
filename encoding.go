package rs_coding

import (
	"github.com/nansi8/math"
	gomath "math"
)

type Encoder struct {
	dataBlocks     int
	checksumBlocks int
	degree         byte
}

func NewEncoder(dataBlocks, checksumBlocks int, degree byte) *Encoder {
	encoder := new(Encoder)
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
		block.index = i
		if i < e.dataBlocks {
			block.blockType = Data
		} else {
			block.blockType = Checksum
		}
		blockData := make([]byte, blocksNumber)
		for j := 0; j < len(resultBlocks); j++ {
			blockData[j] = resultBlocks[j][i]
		}
		block.w = blockData
		blocks[i] = *block
	}
	return blocks
}
