package core

import (
	"math"
)

type flagValue struct {
	startOffset byte
	endOffset   byte
}

func (flagVal *flagValue) parse() uint64 {

	res := uint64(0)
	for i := flagVal.startOffset; i < flagVal.endOffset; i++ {
		res += uint64(math.Pow10(int(i)))
	}

	return res
}
