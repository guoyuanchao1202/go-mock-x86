package cr4

import (
	"fmt"
	"go-mock-x86/pkg/core"
	"go-mock-x86/pkg/registers/flieds/cr4fields"
)

type cr4 struct {
	value       uint64
	reservedBit uint64
	maskBit uint64
	fields  map[string]core.IField
}

var cr4_ *cr4

func init() {
	fieldsMap := make(map[string]core.IField)
	fieldsMap[cr4fields.VMEField] = cr4fields.VME
	fieldsMap[cr4fields.PVIField] = cr4fields.PVI
	fieldsMap[cr4fields.TSDField] = cr4fields.TSD
	fieldsMap[cr4fields.PSEField] = cr4fields.PSE
	fieldsMap[cr4fields.PAEField] = cr4fields.PAE
	fieldsMap[cr4fields.MCEField] = cr4fields.MCE
	fieldsMap[cr4fields.PGEField] = cr4fields.PGE
	fieldsMap[cr4fields.PCEField] = cr4fields.PCE
	fieldsMap[cr4fields.UMIPField] = cr4fields.UMIP
	fieldsMap[cr4fields.VMXEField] = cr4fields.VMXE
	fieldsMap[cr4fields.SMXEField] = cr4fields.SMXE
	fieldsMap[cr4fields.KLField] = cr4fields.KL
	fieldsMap[cr4fields.SMEPField] = cr4fields.SMEP
	fieldsMap[cr4fields.SMAPField] = cr4fields.SMAP
	fieldsMap[cr4fields.PKEField] = cr4fields.PKE
	fieldsMap[cr4fields.CETField] = cr4fields.CET
	fieldsMap[cr4fields.PKSField] = cr4fields.PKS

	reservedBit := uint64(0)
	for _, Field := range fieldsMap {
		reservedBit += Field.Value()
	}
	cr4_ = &cr4{
		reservedBit: reservedBit,
		fields:      fieldsMap,
	}
}

func CR4() *cr4 {
	return cr4_
}

func (cr4 *cr4) Set(fields core.IField) error {
	if _, ok := cr4.fields[fields.String()]; !ok {
		return fmt.Errorf("invalid cr4 set operation: unsupported fields of cr4: %s", fields.String())
	}

	if (fields.Value() | cr4.reservedBit) != cr4.reservedBit {
		return fmt.Errorf("invalid cr4 set operation: mustn't set reserved bit")
	}

	cr4.value |= fields.Value()
	return nil
}

func (cr4 *cr4) SetMaskBit(mask uint64) {
	cr4.maskBit = mask
}

func (cr4 *cr4) Clr(fieldsName string) error {
	fields, ok := cr4.fields[fieldsName]
	if !ok {
		return fmt.Errorf("invalid cr4 clr operation: unsupported fields of cr4: %s", fieldsName)
	}

	cr4.value ^= fields.Value()

	return nil
}

func (cr4 *cr4) Get(fieldName string) (uint64, error) {
	field, ok := cr4.fields[fieldName]
	if !ok {
		return 0, fmt.Errorf("invalid cr4 get operation: unsupported Field of cr4: %s", fieldName)
	}

	return cr4.value & field.Value() & cr4.maskBit, nil
}
