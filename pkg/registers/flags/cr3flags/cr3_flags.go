package cr3flags

import (
	"go-mock-x86/pkg/core"
	core2 "go-mock-x86/pkg/registers/internal/core"
)

const (
	PWTFlag     = "PWT"
	PCDFlag     = "PCD"
	PDEPAFlag   = "PDE table phyAddr"
	PDPTEPAFlag = "PDPTE table phyAddr"
	PML4EPAFlag = "PML4 table phyAddr"
	PML5EPAFlag = "PML5 table phyAddr"
	PCIDFlag    = "PCID"
)

var (
	PWT     core.IFlag
	PCD     core.IFlag
	PDEPA   core.IFlag
	PDPTEPA core.IFlag
	PML4EPA core.IFlag
	PML5EPA core.IFlag
	PCID    core.IFlag
)

func init() {
	PWT = core2.NewRegisterFlag(PWTFlag, 3, 4)
	PCD = core2.NewRegisterFlag(PCDFlag, 4, 5)
	PDEPA = core2.NewRegisterFlag(PDEPAFlag, 12, 32)
	PDPTEPA = core2.NewRegisterFlag(PDPTEPAFlag, 5, 32)
	PML4EPA = core2.NewRegisterFlag(PML4EPAFlag, 12, 47)
	PML5EPA = core2.NewRegisterFlag(PML5EPAFlag, 12, 56)
	PCID = core2.NewRegisterFlag(PCIDFlag, 0, 12)
}
