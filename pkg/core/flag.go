package core

type IFlag interface {
	String() string
	Value() uint64
}

const (
	Bit0 = uint64(1) << iota
	Bit1
	Bit2
	Bit3
	Bit4
	Bit5
	Bit6
	Bit7
	Bit8
	Bit9
	Bit10
	Bit11
	Bit12
	Bit13
	Bit14
	Bit15
	Bit16
	Bot17
	Bit18
	Bit19
	Bit20
	Bit21
	Bit22
	Bit23
	Bit24
	Bit25
	Bit26
	Bit27
	Bit28
	Bit29
	Bit30
	Bit31
	Bit32
	Bit33
	Bit34
	Bit35
	Bit36
	Bit37
	Bit38
	Bit39
	Bit40
	Bot41
	Bit42
	Bit43
	Bit44
	Bit45
	Bit46
	Bit47
	Bit48
	Bit49
	Bit50
	Bit51
	Bit52
	Bit53
	Bit54
	Bit55
	Bit56
	Bit57
	Bit58
	Bit59
	Bit60
	Bit61
	Bit62
	Bit63
)

var offSetBitMap = map[byte]uint64{
	0:  Bit0,
	1:  Bit1,
	2:  Bit2,
	3:  Bit3,
	4:  Bit4,
	5:  Bit5,
	6:  Bit6,
	7:  Bit7,
	8:  Bit8,
	9:  Bit9,
	10: Bit10,
	11: Bit11,
	12: Bit12,
	13: Bit13,
	14: Bit14,
	15: Bit15,
	16: Bit16,
	17: Bot17,
	18: Bit18,
	19: Bit19,
	20: Bit20,
	21: Bit21,
	22: Bit22,
	23: Bit23,
	24: Bit24,
	25: Bit25,
	26: Bit26,
	27: Bit27,
	28: Bit28,
	29: Bit29,
	30: Bit30,
	31: Bit31,
	32: Bit32,
	33: Bit33,
	34: Bit34,
	35: Bit35,
	36: Bit36,
	37: Bit37,
	38: Bit38,
	39: Bit39,
	40: Bit40,
	41: Bot41,
	42: Bit42,
	43: Bit43,
	44: Bit44,
	45: Bit45,
	46: Bit46,
	47: Bit47,
	48: Bit48,
	49: Bit49,
	50: Bit50,
	51: Bit51,
	52: Bit52,
	53: Bit53,
	54: Bit54,
	55: Bit55,
	56: Bit56,
	57: Bit57,
	58: Bit58,
	59: Bit59,
	60: Bit60,
	61: Bit61,
	62: Bit62,
	63: Bit63,
}

func GetValue(startOffset, endOffset byte) uint64 {
	res := uint64(0)
	for i := startOffset; i < endOffset; i++ {
		res += offSetBitMap[i]
	}
	return res
}
