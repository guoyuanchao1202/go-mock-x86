package cr0fields

import (
	"go-mock-x86/pkg/core"
	core2 "go-mock-x86/pkg/registers/internal/core"
)

const (
	PEField = "PE"
	MPField = "MP"
	EMField = "EM"
	TSField = "TS"
	ETField = "ET"
	NEField = "NE"
	WPField = "WP"
	AMField = "AM"
	NWField = "NW"
	CDField = "CD"
	PGField = "PG"
)

var (
	PE core.IField
	MP core.IField
	EM core.IField
	TS core.IField
	ET core.IField
	NE core.IField
	WP core.IField
	AM core.IField
	NW core.IField
	CD core.IField
	PG core.IField
)

func init() {
	PE = core2.NewRegisterField(PEField, 0, 1)
	MP = core2.NewRegisterField(MPField, 1, 2)
	EM = core2.NewRegisterField(EMField, 2, 3)
	TS = core2.NewRegisterField(TSField, 3, 4)
	ET = core2.NewRegisterField(ETField, 4, 5)
	NE = core2.NewRegisterField(NEField, 5, 6)
	WP = core2.NewRegisterField(WPField, 16, 17)
	AM = core2.NewRegisterField(AMField, 18, 19)
	NW = core2.NewRegisterField(NWField, 29, 30)
	CD = core2.NewRegisterField(CDField, 30, 31)
	PG = core2.NewRegisterField(PGField, 31, 32)
}
