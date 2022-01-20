package cr4flags

import (
	"go-mock-x86/pkg/core"
	core2 "go-mock-x86/pkg/registers/internal/core"
)

const (
	VMEFlag  = "VME"
	PVIFlag  = "PVI"
	TSDFlag  = "TSD"
	PSEFlag  = "PSE"
	PAEFlag  = "PAE"
	MCEFlag  = "MCE"
	PGEFlag  = "PGE"
	PCEFlag  = "PCE"
	UMIPFlag = "UMIP"
	VMXEFlag = "VMXE"
	SMXEFlag = "SMXE"
	KLFlag   = "KL"
	SMEPFlag = "SMEP"
	SMAPFlag = "SMAP"
	PKEFlag  = "PKE"
	CETFlag  = "CET"
	PKSFlag  = "PKS"
)

var (
	VME  core.IFlag
	PVI  core.IFlag
	TSD  core.IFlag
	PSE  core.IFlag
	PAE  core.IFlag
	MCE  core.IFlag
	PGE  core.IFlag
	PCE  core.IFlag
	UMIP core.IFlag
	VMXE core.IFlag
	SMXE core.IFlag
	KL   core.IFlag
	SMEP core.IFlag
	SMAP core.IFlag
	PKE  core.IFlag
	CET  core.IFlag
	PKS  core.IFlag
)

func init() {
	VME = core2.NewRegisterFlag(VMEFlag, 0, 1)
	PVI = core2.NewRegisterFlag(PVIFlag, 1, 2)
	TSD = core2.NewRegisterFlag(TSDFlag, 2, 3)
	PSE = core2.NewRegisterFlag(PSEFlag, 3, 4)
	PAE = core2.NewRegisterFlag(PAEFlag, 4, 5)
	MCE = core2.NewRegisterFlag(MCEFlag, 5, 6)
	PGE = core2.NewRegisterFlag(PGEFlag, 7, 8)
	PCE = core2.NewRegisterFlag(PCEFlag, 8, 9)
	UMIP = core2.NewRegisterFlag(UMIPFlag, 11, 12)
	VMXE = core2.NewRegisterFlag(VMXEFlag, 13, 14)
	SMXE = core2.NewRegisterFlag(SMXEFlag, 14, 15)
	KL = core2.NewRegisterFlag(KLFlag, 19, 20)
	SMEP = core2.NewRegisterFlag(SMEPFlag, 20, 21)
	SMAP = core2.NewRegisterFlag(SMAPFlag, 21, 22)
	PKE = core2.NewRegisterFlag(PKEFlag, 22, 23)
	CET = core2.NewRegisterFlag(CETFlag, 23, 24)
	PKS = core2.NewRegisterFlag(PKSFlag, 24, 25)
}
