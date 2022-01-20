package cr4

import (
	"fmt"
	"go-mock-x86/pkg/core"
	"go-mock-x86/pkg/registers/flags/cr4flags"
)

type cr4 struct {
	value       uint64
	reservedBit uint64
	maskBit     uint64
	flagMap     map[string]core.IFlag
}

var cr4_ *cr4

func init() {
	flagMap := make(map[string]core.IFlag)
	flagMap[cr4flags.VMEFlag] = cr4flags.VME
	flagMap[cr4flags.PVIFlag] = cr4flags.PVI
	flagMap[cr4flags.TSDFlag] = cr4flags.TSD
	flagMap[cr4flags.PSEFlag] = cr4flags.PSE
	flagMap[cr4flags.PAEFlag] = cr4flags.PAE
	flagMap[cr4flags.MCEFlag] = cr4flags.MCE
	flagMap[cr4flags.PGEFlag] = cr4flags.PGE
	flagMap[cr4flags.PCEFlag] = cr4flags.PCE
	flagMap[cr4flags.UMIPFlag] = cr4flags.UMIP
	flagMap[cr4flags.VMXEFlag] = cr4flags.VMXE
	flagMap[cr4flags.SMXEFlag] = cr4flags.SMXE
	flagMap[cr4flags.KLFlag] = cr4flags.KL
	flagMap[cr4flags.SMEPFlag] = cr4flags.SMEP
	flagMap[cr4flags.SMAPFlag] = cr4flags.SMAP
	flagMap[cr4flags.PKEFlag] = cr4flags.PKE
	flagMap[cr4flags.CETFlag] = cr4flags.CET
	flagMap[cr4flags.PKSFlag] = cr4flags.PKS

	reservedBit := uint64(0)
	for _, flag := range flagMap {
		reservedBit += flag.Value()
	}
	cr4_ = &cr4{
		reservedBit: reservedBit,
		flagMap:     flagMap,
	}
}

func CR4() *cr4 {
	return cr4_
}

func (cr4 *cr4) Set(flag core.IFlag) error {
	if _, ok := cr4.flagMap[flag.String()]; !ok {
		return fmt.Errorf("invalid cr4 set operation: unsupported flag of cr4: %s", flag.String())
	}

	if (flag.Value() | cr4.reservedBit) != cr4.reservedBit {
		return fmt.Errorf("invalid cr4 set operation: mustn't set reserved bit")
	}

	cr4.value |= flag.Value()
	return nil
}

func (cr4 *cr4) SetMaskBit(mask uint64) {
	cr4.maskBit = mask
}

func (cr4 *cr4) Clr(flagName string) error {
	flag, ok := cr4.flagMap[flagName]
	if !ok {
		return fmt.Errorf("invalid cr4 clr operation: unsupported flag of cr4: %s", flagName)
	}

	cr4.value ^= flag.Value()

	return nil
}

func (cr4 *cr4) Get(flagName string) (uint64, error) {
	flag, ok := cr4.flagMap[flagName]
	if !ok {
		return 0, fmt.Errorf("invalid cr4 get operation: unsupported flag of cr4: %s", flagName)
	}

	return cr4.value & flag.Value() & cr4.maskBit, nil
}
