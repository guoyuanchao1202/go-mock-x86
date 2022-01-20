package cr2

import (
	"fmt"
	"go-mock-x86/pkg/core"
	"go-mock-x86/pkg/registers/flieds/cr2fields"
)

type cr2 struct {
	value   uint64
	maskBit   uint64
	fieldsMap map[string]core.IField
}

var cr2_ *cr2

func init() {
	fieldsMap := make(map[string]core.IField)
	fieldsMap[cr2fields.FaultPageAddressField] = cr2fields.FaultPageAddress
	cr2_ = &cr2{
		fieldsMap: fieldsMap,
	}
}

func CR2() *cr2 {
	return cr2_
}

func (cr2 *cr2) Set(field core.IField) error {
	if _, ok := cr2.fieldsMap[field.String()]; !ok {
		return fmt.Errorf("invalid cr2 set operation: unsupported field of cr2: %s", field.String())
	}

	cr2.value |= field.Value()
	return nil
}

func (cr2 *cr2) SetMaskBit(mask uint64) {
	cr2.maskBit = mask
}

func (cr2 *cr2) Clr(fieldName string) error {
	field, ok := cr2.fieldsMap[fieldName]
	if !ok {
		return fmt.Errorf("invalid cr2 clr operation: unsupported Field of cr2: %s", fieldName)
	}

	cr2.value ^= field.Value()

	return nil
}

func (cr2 *cr2) Get(fieldName string) (uint64, error) {
	field, ok := cr2.fieldsMap[fieldName]
	if !ok {
		return 0, fmt.Errorf("invalid cr2 get operation: unsupported Field of cr2: %s", fieldName)
	}

	return cr2.value & field.Value() & cr2.maskBit, nil
}
