package consts

type ProcessorMode string

const (
	RealMode             ProcessorMode = "real mode"
	ProtectMode          ProcessorMode = "protect"
	X86CompatibilityMode ProcessorMode = "x86-compatibility mode"
	X8664Mode            ProcessorMode = "x86-64 mode"
)
