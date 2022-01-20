package eflagsflags

import (
	"go-mock-x86/pkg/core"
	core2 "go-mock-x86/pkg/registers/internal/core"
)

const (
	CFFlag   = "CF"
	PFFlag   = "PF"
	AFFlag   = "AF"
	ZFFlag   = "ZF"
	SFFlag   = "SF"
	TFFlag   = "TF"
	IFFlag   = "IF"
	DFFlag   = "DF"
	OFFlag   = "OF"
	IOPLFlag = "IOPL"
	NTFlag   = "NT"
	RFFlag   = "RF"
	VMFlag   = "VM"
	ACFlag   = "AC"
	VIFFlag  = "VIF"
	VIPFlag  = "VIP"
	IDFlag   = "ID"
)

var (
	CF   core.IFlag
	PF   core.IFlag
	AF   core.IFlag
	ZF   core.IFlag
	SF   core.IFlag
	TF   core.IFlag
	IF   core.IFlag
	DF   core.IFlag
	OF   core.IFlag
	IOPL core.IFlag
	NT   core.IFlag
	RF   core.IFlag
	VM   core.IFlag
	AC   core.IFlag
	VIF  core.IFlag
	VIP  core.IFlag
	ID   core.IFlag
)

func init() {
	CF = core2.NewRegisterFlag(CFFlag, 0, 1)
	PF = core2.NewRegisterFlag(PFFlag, 2, 3)
	AF = core2.NewRegisterFlag(AFFlag, 4, 5)
	ZF = core2.NewRegisterFlag(ZFFlag, 6, 7)
	SF = core2.NewRegisterFlag(SFFlag, 7, 8)
	TF = core2.NewRegisterFlag(TFFlag, 8, 9)
	IF = core2.NewRegisterFlag(IFFlag, 9, 10)
	DF = core2.NewRegisterFlag(DFFlag, 10, 11)
	OF = core2.NewRegisterFlag(OFFlag, 11, 12)
	IOPL = core2.NewRegisterFlag(IOPLFlag, 12, 14)
	NT = core2.NewRegisterFlag(NTFlag, 14, 15)
	RF = core2.NewRegisterFlag(RFFlag, 16, 17)
	VM = core2.NewRegisterFlag(VMFlag, 17, 18)
	AC = core2.NewRegisterFlag(ACFlag, 18, 19)
	VIF = core2.NewRegisterFlag(VIFFlag, 19, 20)
	VIP = core2.NewRegisterFlag(VIPFlag, 20, 21)
	ID = core2.NewRegisterFlag(IDFlag, 21, 22)
}
