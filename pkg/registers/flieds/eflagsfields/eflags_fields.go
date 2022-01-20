package eflagsfields

import (
	"go-mock-x86/pkg/core"
	core2 "go-mock-x86/pkg/registers/internal/core"
)

const (
	CFField   = "CF"
	PFField   = "PF"
	AFField   = "AF"
	ZFField   = "ZF"
	SFField   = "SF"
	TFField   = "TF"
	IFField   = "IF"
	DFField   = "DF"
	OFField   = "OF"
	IOPLField = "IOPL"
	NTField   = "NT"
	RFField   = "RF"
	VMField   = "VM"
	ACField   = "AC"
	VIFField  = "VIF"
	VIPField  = "VIP"
	IDField   = "ID"
)

var (
	CF   core.IField
	PF   core.IField
	AF   core.IField
	ZF   core.IField
	SF   core.IField
	TF   core.IField
	IF   core.IField
	DF   core.IField
	OF   core.IField
	IOPL core.IField
	NT   core.IField
	RF   core.IField
	VM   core.IField
	AC   core.IField
	VIF  core.IField
	VIP  core.IField
	ID   core.IField
)

func init() {
	CF = core2.NewRegisterField(CFField, 0, 1)
	PF = core2.NewRegisterField(PFField, 2, 3)
	AF = core2.NewRegisterField(AFField, 4, 5)
	ZF = core2.NewRegisterField(ZFField, 6, 7)
	SF = core2.NewRegisterField(SFField, 7, 8)
	TF = core2.NewRegisterField(TFField, 8, 9)
	IF = core2.NewRegisterField(IFField, 9, 10)
	DF = core2.NewRegisterField(DFField, 10, 11)
	OF = core2.NewRegisterField(OFField, 11, 12)
	IOPL = core2.NewRegisterField(IOPLField, 12, 14)
	NT = core2.NewRegisterField(NTField, 14, 15)
	RF = core2.NewRegisterField(RFField, 16, 17)
	VM = core2.NewRegisterField(VMField, 17, 18)
	AC = core2.NewRegisterField(ACField, 18, 19)
	VIF = core2.NewRegisterField(VIFField, 19, 20)
	VIP = core2.NewRegisterField(VIPField, 20, 21)
	ID = core2.NewRegisterField(IDField, 21, 22)
}
