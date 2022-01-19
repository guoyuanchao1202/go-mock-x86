package cr0

import (
	"go-mock-x86/pkg/core"
	core2 "go-mock-x86/pkg/registers/internal/core"
)

const (
	PEFlag = "PE"
	MPFlag = "MP"
	EMFlag = "EM"
	TSFlag = "TS"
	ETFlag = "ET"
	NEFlag = "NE"
	WPFlag = "WP"
	AMFlag = "AM"
	NWFlag = "NW"
	CDFlag = "CD"
	PGFlag = "PG"
)

var (
	PE core.IFlag
	MP core.IFlag
	EM core.IFlag
	TS core.IFlag
	ET core.IFlag
	NE core.IFlag
	WP core.IFlag
	AM core.IFlag
	NW core.IFlag
	CD core.IFlag
	PG core.IFlag
)

var flagMap map[string]core.IFlag

func initFlag() {
	PE = core2.NewRegisterFlag(PEFlag, 0, 1)
	MP = core2.NewRegisterFlag(MPFlag, 1, 2)
	EM = core2.NewRegisterFlag(EMFlag, 2, 3)
	TS = core2.NewRegisterFlag(TSFlag, 3, 4)
	ET = core2.NewRegisterFlag(ETFlag, 4, 5)
	NE = core2.NewRegisterFlag(NEFlag, 5, 6)
	WP = core2.NewRegisterFlag(WPFlag, 16, 17)
	AM = core2.NewRegisterFlag(AMFlag, 18, 19)
	NW = core2.NewRegisterFlag(NWFlag, 29, 30)
	CD = core2.NewRegisterFlag(CDFlag, 30, 31)
	PG = core2.NewRegisterFlag(PGFlag, 31, 32)

	flagMap = make(map[string]core.IFlag)
	flagMap[PEFlag] = PE
	flagMap[MPFlag] = MP
	flagMap[EMFlag] = EM
	flagMap[TSFlag] = TS
	flagMap[ETFlag] = ET
	flagMap[NEFlag] = NE
	flagMap[WPFlag] = WP
	flagMap[AMFlag] = AM
	flagMap[NWFlag] = NW
	flagMap[CDFlag] = CD
	flagMap[PGFlag] = PG
}
