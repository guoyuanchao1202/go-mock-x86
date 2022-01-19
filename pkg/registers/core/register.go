package core

type IRegister interface {
	Set(flag IFlag) error
	Clr(flagName string) error
	Get(flagName string) (uint64, error)
}
