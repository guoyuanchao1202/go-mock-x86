package core

import "go-mock-x86/pkg/core"

type RegisterFlag struct {
	name  string
	value uint64
}

func NewRegisterFlag(name string, startOffset, endOffset byte) core.IFlag {
	res := &RegisterFlag{
		name:  name,
		value: core.GetValue(startOffset, endOffset),
	}
	return res
}

func (flag *RegisterFlag) String() string {
	return flag.name
}

func (flag *RegisterFlag) Value() uint64 {
	return flag.value
}
