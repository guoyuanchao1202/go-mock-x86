package eflags

import (
	"fmt"
	"go-mock-x86/pkg/core"
	"go-mock-x86/pkg/registers/flieds/eflagsfields"
)

type eFlags struct {
	value       uint64
	reservedBit uint64
	maskBit     uint64
	fieldsMap   map[string]core.IField
}

var eflags_ *eFlags

func init() {
	fieldsMap := make(map[string]core.IField)
	fieldsMap[eflagsfields.CFField] = eflagsfields.CF
	fieldsMap[eflagsfields.PFField] = eflagsfields.PF
	fieldsMap[eflagsfields.AFField] = eflagsfields.AF
	fieldsMap[eflagsfields.ZFField] = eflagsfields.ZF
	fieldsMap[eflagsfields.SFField] = eflagsfields.SF
	fieldsMap[eflagsfields.TFField] = eflagsfields.TF
	fieldsMap[eflagsfields.IFField] = eflagsfields.IF
	fieldsMap[eflagsfields.DFField] = eflagsfields.DF
	fieldsMap[eflagsfields.OFField] = eflagsfields.OF
	fieldsMap[eflagsfields.IOPLField] = eflagsfields.IOPL
	fieldsMap[eflagsfields.NTField] = eflagsfields.NT
	fieldsMap[eflagsfields.RFField] = eflagsfields.RF
	fieldsMap[eflagsfields.ACField] = eflagsfields.AC
	fieldsMap[eflagsfields.VIFField] = eflagsfields.VIF
	fieldsMap[eflagsfields.VIPField] = eflagsfields.VIP
	fieldsMap[eflagsfields.IDField] = eflagsfields.ID

	reservedBit := uint64(0)
	for _, Field := range fieldsMap {
		reservedBit += Field.Value()
	}
	reservedBit += core.GetValue(1, 2)
	eflags_ = &eFlags{
		reservedBit: reservedBit,
		fieldsMap:   fieldsMap,
	}
}

func EFLAGS() *eFlags {
	return eflags_
}

func (eflags *eFlags) Set(field core.IField) error {
	if _, ok := eflags.fieldsMap[field.String()]; !ok {
		return fmt.Errorf("invalid eflagsfields set operation: unsupported field of eflagsfields: %s", field.String())
	}

	if (field.Value() | eflags.reservedBit) != eflags.reservedBit {
		return fmt.Errorf("invalid eflagsfields set operation: mustn't set reserved bit")
	}

	eflags.value |= field.Value()
	return nil
}

func (eflags *eFlags) SetMaskBit(mask uint64) {
	eflags.maskBit = mask
}

func (eflags *eFlags) Clr(fieldName string) error {
	field, ok := eflags.fieldsMap[fieldName]
	if !ok {
		return fmt.Errorf("invalid eflagsfields clr operation: unsupported Field of eflagsfields: %s", fieldName)
	}

	eflags.value ^= field.Value()

	return nil
}

func (eflags *eFlags) Get(fieldName string) (uint64, error) {
	field, ok := eflags.fieldsMap[fieldName]
	if !ok {
		return 0, fmt.Errorf("invalid eflagsfields get operation: unsupported Field of eflagsfields: %s", fieldName)
	}

	return eflags.value & field.Value() & eflags.maskBit, nil
}
