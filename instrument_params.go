package liblsdj

type lengthParam struct {
	lengthUnlimited bool
	length          byte
}

func (l *lengthParam) set(b byte) {
	l.lengthUnlimited = getBit(b, 6) == 0
	l.length = clearBit(b, 6)
}

type byte5Param struct {
	pitchStep        bool
	pitchDrum        bool
	transpose        bool
	pitchSpeed       bool
	tableSpeed       bool
	vibratoShape     byte
	vibratoDirection bool
}

func (l *byte5Param) set(b byte) {
	l.pitchStep = getBit(b, 7) == 1
	l.pitchDrum = getBit(b, 6) == 1
	l.transpose = getBit(b, 5) == 0
	l.pitchSpeed = getBit(b, 4) == 0
	l.tableSpeed = getBit(b, 3) == 0
	//l.vibratoShape = get a range of bit from 1 to 2
	l.vibratoDirection = getBit(b, 0) == 0
}

type instrTableParam struct {
	tableOn bool
	table   byte
}

func (t *instrTableParam) set(b byte) {
	t.tableOn = getBit(b, 5) == 1
	// Table is [0:4]
	t.table = clearBit(b, 5)
}
