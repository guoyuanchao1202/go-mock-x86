package cr4fields

import (
	"go-mock-x86/pkg/core"
	core2 "go-mock-x86/pkg/registers/internal/core"
)

const (
	VMEField  = "VME"
	PVIField  = "PVI"
	TSDField  = "TSD"
	PSEField  = "PSE"
	PAEField  = "PAE"
	MCEField  = "MCE"
	PGEField  = "PGE"
	PCEField  = "PCE"
	UMIPField = "UMIP"
	VMXEField = "VMXE"
	SMXEField = "SMXE"
	KLField   = "KL"
	SMEPField = "SMEP"
	SMAPField = "SMAP"
	PKEField  = "PKE"
	CETField  = "CET"
	PKSField  = "PKS"
)

var (
	VME  core.IField
	PVI  core.IField
	TSD  core.IField
	PSE  core.IField
	PAE  core.IField
	MCE  core.IField
	PGE  core.IField
	PCE  core.IField
	UMIP core.IField
	VMXE core.IField
	SMXE core.IField
	KL   core.IField
	SMEP core.IField
	SMAP core.IField
	PKE  core.IField
	CET  core.IField
	PKS  core.IField
)

func init() {
	VME = core2.NewRegisterField(VMEField, 0, 1)
	PVI = core2.NewRegisterField(PVIField, 1, 2)
	TSD = core2.NewRegisterField(TSDField, 2, 3)
	PSE = core2.NewRegisterField(PSEField, 3, 4)
	PAE = core2.NewRegisterField(PAEField, 4, 5)
	MCE = core2.NewRegisterField(MCEField, 5, 6)
	PGE = core2.NewRegisterField(PGEField, 7, 8)
	PCE = core2.NewRegisterField(PCEField, 8, 9)
	UMIP = core2.NewRegisterField(UMIPField, 11, 12)
	VMXE = core2.NewRegisterField(VMXEField, 13, 14)
	SMXE = core2.NewRegisterField(SMXEField, 14, 15)
	KL = core2.NewRegisterField(KLField, 19, 20)
	SMEP = core2.NewRegisterField(SMEPField, 20, 21)
	SMAP = core2.NewRegisterField(SMAPField, 21, 22)
	PKE = core2.NewRegisterField(PKEField, 22, 23)
	CET = core2.NewRegisterField(CETField, 23, 24)
	PKS = core2.NewRegisterField(PKSField, 24, 25)
}
