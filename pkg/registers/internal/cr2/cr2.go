package cr2

import (
	"fmt"
	"go-mock-x86/pkg/core"
	"go-mock-x86/pkg/registers/flags/cr2flags"
)

type cr2 struct {
	value   uint64
	maskBit uint64
	flagMap map[string]core.IFlag
}

var cr2_ *cr2

func init() {
	flagMap := make(map[string]core.IFlag)
	flagMap[cr2flags.FaultPageAddressFlag] = cr2flags.FaultPageAddress
	cr2_ = &cr2{
		flagMap: flagMap,
	}
}

func CR2() *cr2 {
	return cr2_
}

func (cr2 *cr2) Set(flag core.IFlag) error {
	if _, ok := cr2.flagMap[flag.String()]; !ok {
		return fmt.Errorf("invalid cr0 set operation: unsupported flag of cr2: %s", flag.String())
	}

	cr2.value |= flag.Value()
	return nil
}

func (cr2 *cr2) SetMaskBit(mask uint64) {
	cr2.maskBit = mask
}

func (cr2 *cr2) Clr(flagName string) error {
	flag, ok := cr2.flagMap[flagName]
	if !ok {
		return fmt.Errorf("invalid cr2 clr operation: unsupported flag of cr2: %s", flagName)
	}

	cr2.value ^= flag.Value()

	return nil
}

func (cr2 *cr2) Get(flagName string) (uint64, error) {
	flag, ok := cr2.flagMap[flagName]
	if !ok {
		return 0, fmt.Errorf("invalid cr2 get operation: unsupported flag of cr2: %s", flagName)
	}

	return cr2.value & flag.Value() & cr2.maskBit, nil
}
