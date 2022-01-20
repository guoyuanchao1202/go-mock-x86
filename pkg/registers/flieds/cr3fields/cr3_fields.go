package cr3fields

import (
	"go-mock-x86/pkg/core"
	core2 "go-mock-x86/pkg/registers/internal/core"
)

const (
	PWTField     = "PWT"
	PCDField     = "PCD"
	PDEPAField   = "PDE table phyAddr"
	PDPTEPAField = "PDPTE table phyAddr"
	PML4EPAField = "PML4 table phyAddr"
	PML5EPAField = "PML5 table phyAddr"
	PCIDField    = "PCID"
)

var (
	PWT     core.IField
	PCD     core.IField
	PDEPA   core.IField
	PDPTEPA core.IField
	PML4EPA core.IField
	PML5EPA core.IField
	PCID    core.IField
)

func init() {
	PWT = core2.NewRegisterField(PWTField, 3, 4)
	PCD = core2.NewRegisterField(PCDField, 4, 5)
	PDEPA = core2.NewRegisterField(PDEPAField, 12, 32)
	PDPTEPA = core2.NewRegisterField(PDPTEPAField, 5, 32)
	PML4EPA = core2.NewRegisterField(PML4EPAField, 12, 47)
	PML5EPA = core2.NewRegisterField(PML5EPAField, 12, 56)
	PCID = core2.NewRegisterField(PCIDField, 0, 12)
}
