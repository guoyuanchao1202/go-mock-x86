package cr4

import (
	"fmt"
	"go-mock-x86/pkg/core"
)


type cr4 struct {
	value       uint64
	reservedBit uint64
	maskBit     uint64
	flagMap     map[string]core.IFlag
}

var cr4_ *cr4

func init() {

}

func CR4() *cr4 {
	return cr4_
}

func (cr4 *cr4) Set(flag core.IFlag) error {
	if _, ok := cr4.flagMap[flag.String()]; !ok {
		return fmt.Errorf("invalid cr0 set operation: unsupported flag of cr4: %s", flag.String())
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
