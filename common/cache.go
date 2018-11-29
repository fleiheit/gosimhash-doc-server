package common

import (
	"github.com/fleiheit/gosimhash"
	"github.com/fleiheit/gosimhash-doc-server/model"
)

type SimhashLimit int

const (
	Limit1  = SimhashLimit(1)
	Limit3  = SimhashLimit(3)
	Limit7  = SimhashLimit(7)
	Limit15 = SimhashLimit(15)
)

type SimhashCache interface {
	Init(docIds []string, simhashList []uint64, timeouts []int64) int

	InsertIfNotDuplicated(docId string, simhash uint64, age int64) (bool, *model.Document, error)
}

const (
	Mask2  uint64 = 0xffffffff
	Mask4  uint64 = 0xffff
	Mask8  uint64 = 0xff
	Mask16 uint64 = 0xf
)

var MASKS = map[int]uint64{2: Mask2, 4: Mask4, 8: Mask8, 16: Mask16}

type SimhashOperator struct {
	partNumber int
	mask       uint64
}

func NewSimhashOperator(number int) *SimhashOperator {
	op := &SimhashOperator{}
	op.partNumber = number
	op.mask = MASKS[number]
	return op
}

func (op *SimhashOperator) Check() bool {
	return Check2Power(op.partNumber)
}

func (op *SimhashOperator) Cut(simhash uint64) []uint64 {
	var ret = make([]uint64, op.partNumber)
	var move = uint(gosimhash.BitsLength / op.partNumber)
	for i := 0; i < op.partNumber; i++ {
		ret[i] = op.mask & simhash
		simhash >>= move
	}
	return ret
}
