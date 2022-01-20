package core

type IRegister interface {
	Set(flag IField) error
	Clr(flagName string) error
	Get(flagName string) (uint64, error)
}