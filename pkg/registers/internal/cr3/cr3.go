package cr3

import (
	"fmt"
	"go-mock-x86/pkg/core"
	"go-mock-x86/pkg/registers/flags/cr3flags"
)

type cr3 struct {
	value       uint64
	reservedBit uint64
	maskBit     uint64
	flagMap     map[string]core.IFlag
}

var cr3_ *cr3

func init() {
	flagMap := make(map[string]core.IFlag)
	flagMap[cr3flags.PCDFlag] = cr3flags.PCD
	flagMap[cr3flags.PWTFlag] = cr3flags.PWT
	flagMap[cr3flags.PDEPAFlag] = cr3flags.PDEPA
	flagMap[cr3flags.PDPTEPAFlag] = cr3flags.PDPTEPA
	flagMap[cr3flags.PML4EPAFlag] = cr3flags.PML4EPA
	flagMap[cr3flags.PML5EPAFlag] = cr3flags.PML5EPA
	flagMap[cr3flags.PCIDFlag] = cr3flags.PCID
	cr3_ = &cr3{
		flagMap: flagMap,
	}
}

func CR3() *cr3 {
	return cr3_
}

func (cr3 *cr3) Set(flag core.IFlag) error {
	if _, ok := cr3.flagMap[flag.String()]; !ok {
		return fmt.Errorf("invalid cr3 set operation: unsupported flag of cr3: %s", flag.String())
	}

	if (flag.Value() | cr3.reservedBit) != cr3.reservedBit {
		return fmt.Errorf("invalid cr3 set operation: mustn't set reserved bit")
	}

	cr3.value |= flag.Value()
	return nil
}

func (cr3 *cr3) SetMaskBit(mask uint64) {
	cr3.maskBit = mask
}

func (cr3 *cr3) SetReserveBit(reservedBit uint64) {
	cr3.reservedBit = reservedBit
}

func (cr3 *cr3) Clr(flagName string) error {
	flag, ok := cr3.flagMap[flagName]
	if !ok {
		return fmt.Errorf("invalid cr3 clr operation: unsupported flag of cr3: %s", flagName)
	}

	cr3.value ^= flag.Value()

	return nil
}

func (cr3 *cr3) Get(flagName string) (uint64, error) {
	flag, ok := cr3.flagMap[flagName]
	if !ok {
		return 0, fmt.Errorf("invalid cr3 get operation: unsupported flag of cr3: %s", flagName)
	}

	return cr3.value & flag.Value() & cr3.maskBit, nil
}
