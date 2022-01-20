package cr0

import (
	"fmt"
	"go-mock-x86/pkg/core"
	"go-mock-x86/pkg/registers/flags/cr0flags"
)

// 11100000000001010000000000111111

type cr0 struct {
	value       uint64
	reservedBit uint64
	flagMap     map[string]core.IFlag
}

var cr0_ *cr0

func CR0() *cr0 {
	return cr0_
}

func init() {
	flagMap := make(map[string]core.IFlag)
	flagMap[cr0flags.PEFlag] = cr0flags.PE
	flagMap[cr0flags.MPFlag] = cr0flags.MP
	flagMap[cr0flags.EMFlag] = cr0flags.EM
	flagMap[cr0flags.TSFlag] = cr0flags.TS
	flagMap[cr0flags.ETFlag] = cr0flags.ET
	flagMap[cr0flags.NEFlag] = cr0flags.NE
	flagMap[cr0flags.WPFlag] = cr0flags.WP
	flagMap[cr0flags.AMFlag] = cr0flags.AM
	flagMap[cr0flags.NWFlag] = cr0flags.NW
	flagMap[cr0flags.CDFlag] = cr0flags.CD
	flagMap[cr0flags.PGFlag] = cr0flags.PG

	reservedBit := uint64(0)
	for _, flag := range flagMap {
		reservedBit += flag.Value()
	}

	cr0_ = &cr0{
		reservedBit: reservedBit,
		flagMap:     flagMap,
	}
}

func (cr0 *cr0) Set(flag core.IFlag) error {
	if _, ok := cr0.flagMap[flag.String()]; !ok {
		return fmt.Errorf("invalid cr0 set operation: unsupported flag of cr0: %s", flag.String())
	}

	if (flag.Value() | cr0.reservedBit) != cr0.reservedBit {
		return fmt.Errorf("invalid cr0 set operation: mustn't set reserved bit")
	}

	cr0.value |= flag.Value()
	return nil
}

func (cr0 *cr0) Clr(flagName string) error {
	flag, ok := cr0.flagMap[flagName]
	if !ok {
		return fmt.Errorf("invalid cr0 clr operation: unsupported flag of cr0: %s", flagName)
	}

	cr0.value ^= flag.Value()

	return nil
}

func (cr0 *cr0) Get(flagName string) (uint64, error) {
	flag, ok := cr0.flagMap[flagName]
	if !ok {
		return 0, fmt.Errorf("invalid cr0 get operation: unsupported flag of cr0: %s", flagName)
	}

	return cr0.value & flag.Value(), nil
}
