package cr2fields

import (
	"go-mock-x86/pkg/core"
	core2 "go-mock-x86/pkg/registers/internal/core"
)

const (
	FaultPageAddressField   = "FaultPageAddress"
)

var (
	FaultPageAddress   core.IField
)

func init() {
	FaultPageAddress = core2.NewRegisterField(FaultPageAddressField, 0, 56) // in five-level paging, max support 57-bit liner-address
}
