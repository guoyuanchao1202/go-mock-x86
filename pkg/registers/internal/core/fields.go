package core

import "go-mock-x86/pkg/core"

type RegisterField struct {
	name  string
	value uint64
}

func NewRegisterField(name string, startOffset, endOffset byte) core.IField {
	res := &RegisterField{
		name: name,
		value: core.GetValue(startOffset, endOffset),
	}
	return res
}

func (flag *RegisterField) String() string {
	return flag.name
}

func (flag *RegisterField) Value() uint64 {
	return flag.value
}
