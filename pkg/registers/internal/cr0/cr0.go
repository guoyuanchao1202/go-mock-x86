package cr0

import (
	"fmt"
	"go-mock-x86/pkg/core"
	"go-mock-x86/pkg/registers/flieds/cr0fields"
)

// 11100000000001010000000000111111

type cr0 struct {
	value       uint64
	reservedBit uint64
	fieldsMap   map[string]core.IField
}

var cr0_ *cr0

func CR0() *cr0 {
	return cr0_
}

func init() {
	fieldsMap := make(map[string]core.IField)
	fieldsMap[cr0fields.PEField] = cr0fields.PE
	fieldsMap[cr0fields.MPField] = cr0fields.MP
	fieldsMap[cr0fields.EMField] = cr0fields.EM
	fieldsMap[cr0fields.TSField] = cr0fields.TS
	fieldsMap[cr0fields.ETField] = cr0fields.ET
	fieldsMap[cr0fields.NEField] = cr0fields.NE
	fieldsMap[cr0fields.WPField] = cr0fields.WP
	fieldsMap[cr0fields.AMField] = cr0fields.AM
	fieldsMap[cr0fields.NWField] = cr0fields.NW
	fieldsMap[cr0fields.CDField] = cr0fields.CD
	fieldsMap[cr0fields.PGField] = cr0fields.PG

	reservedBit := uint64(0)
	for _, Field := range fieldsMap {
		reservedBit += Field.Value()
	}

	cr0_ = &cr0{
		reservedBit: reservedBit,
		fieldsMap:   fieldsMap,
	}
}

func (cr0 *cr0) Set(field core.IField) error {
	if _, ok := cr0.fieldsMap[field.String()]; !ok {
		return fmt.Errorf("invalid cr0 set operation: unsupported field of cr0: %s", field.String())
	}

	if (field.Value() | cr0.reservedBit) != cr0.reservedBit {
		return fmt.Errorf("invalid cr0 set operation: mustn't set reserved bit")
	}

	cr0.value |= field.Value()
	return nil
}

func (cr0 *cr0) Clr(fieldName string) error {
	field, ok := cr0.fieldsMap[fieldName]
	if !ok {
		return fmt.Errorf("invalid cr0 clr operation: unsupported Field of cr0: %s", fieldName)
	}

	cr0.value ^= field.Value()

	return nil
}

func (cr0 *cr0) Get(fieldName string) (uint64, error) {
	field, ok := cr0.fieldsMap[fieldName]
	if !ok {
		return 0, fmt.Errorf("invalid cr0 get operation: unsupported Field of cr0: %s", fieldName)
	}

	return cr0.value & field.Value(), nil
}
