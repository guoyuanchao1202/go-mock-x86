package cr1

import (
	"fmt"
	"go-mock-x86/pkg/core"
)

type cr1 struct {
	value       uint64
	reservedBit uint64
	fieldMap     map[string]core.IField
}

var cr1_ *cr1

// CR1 this register not used, so all bit are reserved
func CR1() *cr1 {
	return cr1_
}

func init() {
	cr1_ = &cr1{}
}

func (cr1 *cr1) Set(Field core.IField) error {
	return fmt.Errorf("reserved register, not support set operation")
}

func (cr1 *cr1) Clr(FieldName string) error {
	return fmt.Errorf("reserved register, not support clr operation")
}

func (cr1 *cr1) Get(FieldName string) (uint64, error) {
	return 0, fmt.Errorf("reserved register, not support get operation")
}
