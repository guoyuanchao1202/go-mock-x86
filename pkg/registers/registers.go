package registers

import (
	"fmt"
	"go-mock-x86/pkg/core"
	"go-mock-x86/pkg/registers/consts"
	"go-mock-x86/pkg/registers/internal/cr0"
	"go-mock-x86/pkg/registers/internal/cr1"
	"go-mock-x86/pkg/registers/internal/cr2"
	"go-mock-x86/pkg/registers/internal/cr3"
	"go-mock-x86/pkg/registers/internal/cr4"
	"go-mock-x86/pkg/registers/internal/eflags"
)

type registerProcessFunc func(register core.IRegister, flag core.IFlag) error

type RegisterManager struct {
	registers                 map[consts.RegisterName]core.IRegister
	registerSetProcessFuncMap map[consts.RegisterName]registerProcessFunc
	registerClrProcessFuncMap map[consts.RegisterName]registerProcessFunc
}

var registerManager *RegisterManager

func Registers() *RegisterManager {
	return registerManager
}

func init() {
	registerManager = &RegisterManager{
		registers: map[consts.RegisterName]core.IRegister{
			consts.CR0:    cr0.CR0(),
			consts.CR1:    cr1.CR1(),
			consts.CR2:    cr2.CR2(),
			consts.CR3:    cr3.CR3(),
			consts.CR4:    cr4.CR4(),
			consts.EFLAGS: eflags.EFLAGS(),
		},
	}

	registerManager.registerSetProcessFuncMap = map[consts.RegisterName]registerProcessFunc{
		consts.CR0:    registerManager.processCR0SetOperation,
		consts.CR1:    registerManager.processCR1SetOperation,
		consts.CR2:    registerManager.processCR2SetOperation,
		consts.CR3:    registerManager.processCR3SetOperation,
		consts.CR4:    registerManager.processCR4ClrOperation,
		consts.EFLAGS: registerManager.processEFLAGSSetOperation,
	}

	registerManager.registerClrProcessFuncMap = map[consts.RegisterName]registerProcessFunc{
		consts.CR0:    registerManager.processCR0ClrOperation,
		consts.CR1:    registerManager.processCR1ClrOperation,
		consts.CR2:    registerManager.processCR2ClrOperation,
		consts.CR3:    registerManager.processCR3ClrOperation,
		consts.CR4:    registerManager.processCR4ClrOperation,
		consts.EFLAGS: registerManager.processEFLAGSClrOperation,
	}
}

func (rm *RegisterManager) Set(registerName consts.RegisterName, flag core.IFlag) error {
	return rm.do(registerName, flag, rm.registerSetProcessFuncMap)
}

func (rm *RegisterManager) Clr(registerName consts.RegisterName, flag core.IFlag) error {
	return rm.do(registerName, flag, rm.registerClrProcessFuncMap)
}

func (rm *RegisterManager) do(registerName consts.RegisterName, flag core.IFlag, processFuncMap map[consts.RegisterName]registerProcessFunc) error {
	register, ok := rm.registers[registerName]
	if !ok {
		return fmt.Errorf("invalid opeation: unsupport registerManager: %s", registerName)
	}

	return processFuncMap[registerName](register, flag)
}
