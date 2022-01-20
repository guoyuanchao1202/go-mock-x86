package eflags

import (
	"fmt"
	"go-mock-x86/pkg/core"
	"go-mock-x86/pkg/registers/flags/eflagsflags"
)

type eflags struct {
	value       uint64
	reservedBit uint64
	maskBit     uint64
	flagMap     map[string]core.IFlag
}

var eflags_ *eflags


func init() {
	flagMap := make(map[string]core.IFlag)
	flagMap[eflagsflags.CFFlag] = eflagsflags.CF
	flagMap[eflagsflags.PFFlag] = eflagsflags.PF
	flagMap[eflagsflags.AFFlag] = eflagsflags.AF
	flagMap[eflagsflags.ZFFlag] = eflagsflags.ZF
	flagMap[eflagsflags.SFFlag] = eflagsflags.SF
	flagMap[eflagsflags.TFFlag] = eflagsflags.TF
	flagMap[eflagsflags.IFFlag] = eflagsflags.IF
	flagMap[eflagsflags.DFFlag] = eflagsflags.DF
	flagMap[eflagsflags.OFFlag]  =eflagsflags.OF
	flagMap[eflagsflags.IOPLFlag] = eflagsflags.IOPL
	flagMap[eflagsflags.NTFlag] = eflagsflags.NT
	flagMap[eflagsflags.RFFlag] = eflagsflags.RF
	flagMap[eflagsflags.ACFlag] = eflagsflags.AC
	flagMap[eflagsflags.VIFFlag] = eflagsflags.VIF
	flagMap[eflagsflags.VIPFlag] = eflagsflags.VIP
	flagMap[eflagsflags.IDFlag] = eflagsflags.ID

	reservedBit := uint64(0)
	for _, flag := range flagMap {
		reservedBit += flag.Value()
	}
	reservedBit += core.GetValue(1, 2)
	eflags_ = &eflags{
		reservedBit: reservedBit,
		flagMap:     flagMap,
	}
}

func EFLAGS() *eflags {
	return eflags_
}

func (eflags *eflags) Set(flag core.IFlag) error {
	if _, ok := eflags.flagMap[flag.String()]; !ok {
		return fmt.Errorf("invalid eflagsflags set operation: unsupported flag of eflagsflags: %s", flag.String())
	}

	if (flag.Value() | eflags.reservedBit) != eflags.reservedBit {
		return fmt.Errorf("invalid eflagsflags set operation: mustn't set reserved bit")
	}

	eflags.value |= flag.Value()
	return nil
}

func (eflags *eflags) SetMaskBit(mask uint64) {
	eflags.maskBit = mask
}

func (eflags *eflags) Clr(flagName string) error {
	flag, ok := eflags.flagMap[flagName]
	if !ok {
		return fmt.Errorf("invalid eflagsflags clr operation: unsupported flag of eflagsflags: %s", flagName)
	}

	eflags.value ^= flag.Value()

	return nil
}

func (eflags *eflags) Get(flagName string) (uint64, error) {
	flag, ok := eflags.flagMap[flagName]
	if !ok {
		return 0, fmt.Errorf("invalid eflagsflags get operation: unsupported flag of eflagsflags: %s", flagName)
	}

	return eflags.value & flag.Value() & eflags.maskBit, nil
}
