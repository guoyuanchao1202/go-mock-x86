package cr0

import (
	"fmt"
	"go-mock-x86/pkg/registers/core"
)

const (
	bitMask = 11100000000001010000000000111111
)

type CR0 struct {
	value       uint64
	reservedBit uint64
	flagMap     map[string]core.IFlag
}

func NewCR0() core.IRegister {
	return &CR0{
		reservedBit: bitMask,
		flagMap:     flagMap,
	}
}

func (cr0 *CR0) Set(flag core.IFlag) error {
	if _, ok := cr0.flagMap[flag.String()]; !ok {
		return fmt.Errorf("invalid cr0 set operation: unsupported flag of cr0: %s", flag.String())
	}

	if (flag.Value() | cr0.reservedBit) != cr0.reservedBit {
		return fmt.Errorf("invalid cr0 set operation: mustn't set reserved bit")
	}

	cr0.value |= flag.Value()
	return nil
}

func (cr0 *CR0) Clr(flagName string) error {
	flag, ok := cr0.flagMap[flagName]
	if !ok {
		return fmt.Errorf("invalid cr0 clr operation: unsupported flag of cr0: %s", flagName)
	}

	cr0.value ^= flag.Value()

	return nil
}

func (cr0 *CR0) Get(flagName string) (uint64, error) {
	flag, ok := cr0.flagMap[flagName]
	if !ok {
		return 0, fmt.Errorf("invalid cr0 get operation: unsupported flag of cr0: %s", flagName)
	}

	return flag.Value(), nil
}
