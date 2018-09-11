package rs_coding

import (
	"fmt"
	"github.com/nansi8/math"
	"sort"
)

type ErrNotSufficientBlocks struct {
	RequiredBlocks int
	ActualBlocks   int
}

func (e ErrNotSufficientBlocks) Error() string {
	return fmt.Sprintf("Can not decode as %d blocks required, but only %d are present", e.RequiredBlocks, e.ActualBlocks)
}

type Decoder struct {
	dataBlocks     int
	checksumBlocks int
	degree         byte
}

func NewDecoder(dataBlocks, checksumBlocks int, degree byte) *Decoder {
	decoder := new(Decoder)
	decoder.dataBlocks = dataBlocks
	decoder.checksumBlocks = checksumBlocks
	decoder.degree = degree
	return decoder
}

func (d *Decoder) Decode(blocks []Block) ([]byte, error) {
	galoisAlgebra := math.New(d.degree)
	blocksNumber := len(blocks)
	if blocksNumber < d.dataBlocks {
		return nil, ErrNotSufficientBlocks{d.dataBlocks, blocksNumber}
	}

	decodedResult := make([]byte, 0)
	sort.Slice(blocks, func(i, j int) bool {
		return blocks[i].index < blocks[j].index
	})
	blocks = blocks[:d.dataBlocks]
	indexes := getIndexes(blocks)
	checkMatrix := d.getCheckMatrix(vandermore(d.checksumBlocks, d.dataBlocks, galoisAlgebra))
	correctionMatrix := make([][]byte, 0)
	for i := 0; i < len(checkMatrix); i++ {
		if containsIndex(indexes, i) {
			correctionMatrix = append(correctionMatrix, checkMatrix[i])
		}
	}
	for i := 0; i < len(blocks[0].w); i++ {
		encoded := make([]byte, d.dataBlocks)
		for j := 0; j < d.dataBlocks; j++ {
			encoded[j] = blocks[j].w[i]
		}
		encodedMatrix := getDataBlockMatrix(encoded)
		decodedMatrix := math.Mul(math.Reverse(correctionMatrix, galoisAlgebra), encodedMatrix, galoisAlgebra)
		decoded := getCheckBlock(decodedMatrix)
		decodedResult = append(decodedResult, decoded...)
	}
	return decodedResult, nil
}

func (d *Decoder) getCheckMatrix(vandermore [][]byte) [][]byte {
	checkMatrix := make([][]byte, d.dataBlocks+d.checksumBlocks)
	checkMatrix = identity(d.dataBlocks)
	for _, r := range vandermore {
		checkMatrix = append(checkMatrix, r)
	}
	return checkMatrix
}

func getIndexes(blocks []Block) []int {
	indexes := make([]int, len(blocks))
	for i, v := range blocks {
		indexes[i] = v.index
	}
	return indexes
}

func containsIndex(indexes []int, index int) bool {
	for _, v := range indexes {
		if v == index {
			return true
		}
	}
	return false
}
