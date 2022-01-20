package cr3

import (
	"fmt"
	"go-mock-x86/pkg/core"
	"go-mock-x86/pkg/registers/flieds/cr3fields"
)

type cr3 struct {
	value       uint64
	reservedBit uint64
	maskBit     uint64
	fieldsMap     map[string]core.IField
}

var cr3_ *cr3

func init() {
	fieldsMap := make(map[string]core.IField)
	fieldsMap[cr3fields.PCDField] = cr3fields.PCD
	fieldsMap[cr3fields.PWTField] = cr3fields.PWT
	fieldsMap[cr3fields.PDEPAField] = cr3fields.PDEPA
	fieldsMap[cr3fields.PDPTEPAField] = cr3fields.PDPTEPA
	fieldsMap[cr3fields.PML4EPAField] = cr3fields.PML4EPA
	fieldsMap[cr3fields.PML5EPAField] = cr3fields.PML5EPA
	fieldsMap[cr3fields.PCIDField] = cr3fields.PCID
	cr3_ = &cr3{
		fieldsMap: fieldsMap,
	}
}

func CR3() *cr3 {
	return cr3_
}

func (cr3 *cr3) Set(field core.IField) error {
	if _, ok := cr3.fieldsMap[field.String()]; !ok {
		return fmt.Errorf("invalid cr3 set operation: unsupported field of cr3: %s", field.String())
	}

	if (field.Value() | cr3.reservedBit) != cr3.reservedBit {
		return fmt.Errorf("invalid cr3 set operation: mustn't set reserved bit")
	}

	cr3.value |= field.Value()
	return nil
}

func (cr3 *cr3) SetMaskBit(mask uint64) {
	cr3.maskBit = mask
}

func (cr3 *cr3) SetReserveBit(reservedBit uint64) {
	cr3.reservedBit = reservedBit
}

func (cr3 *cr3) Clr(fieldName string) error {
	field, ok := cr3.fieldsMap[fieldName]
	if !ok {
		return fmt.Errorf("invalid cr3 clr operation: unsupported Field of cr3: %s", fieldName)
	}

	cr3.value ^= field.Value()

	return nil
}

func (cr3 *cr3) Get(fieldName string) (uint64, error) {
	field, ok := cr3.fieldsMap[fieldName]
	if !ok {
		return 0, fmt.Errorf("invalid cr3 get operation: unsupported Field of cr3: %s", fieldName)
	}

	return cr3.value & field.Value() & cr3.maskBit, nil
}
