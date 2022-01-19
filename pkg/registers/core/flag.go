package core

type IFlag interface {
	String() string
	Value() uint64
}

type RegisterFlag struct {
	name  string
	value *flagValue
}

func NewRegisterFlag(name string, startOffset, endOffset byte) IFlag {
	res := &RegisterFlag{
		name: name,
		value: &flagValue{
			startOffset: startOffset,
			endOffset:   endOffset,
		},
	}
	return res
}

func (flag *RegisterFlag) String() string {
	return flag.name
}

func (flag *RegisterFlag) Value() uint64 {
	return flag.value.parse()
}
