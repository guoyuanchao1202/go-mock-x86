package cr2flags

import (
	"go-mock-x86/pkg/core"
	core2 "go-mock-x86/pkg/registers/internal/core"
)

const (
	FaultPageAddressFlag   = "FaultPageAddress"
)

var (
	FaultPageAddress   core.IFlag
)

func init() {
	FaultPageAddress = core2.NewRegisterFlag(FaultPageAddressFlag, 0, 56) // in five-level paging, max support 57-bit liner-address
}
