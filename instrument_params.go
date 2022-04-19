package liblsdj

type lengthParam struct {
	lengthUnlimited bool
	length          byte
}

func (l *lengthParam) set(b byte) {
	l.lengthUnlimited = getBit(b, 6) == 0
	l.length = clearBit(b, 6)
}

const (
	b5VibratoDirection uint8 = 1 << iota
	b5VibratoShape2
	b5VibratoShape1
	b5tableSpeed
	b5PitchSpeed
	b5Transpose
	b5PitchDrum
	b5PitchStep
)

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
	l.vibratoShape = getBit(b, 2)*2 + getBit(b, 1)
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

type instrWaveSynth struct {
	wave  byte
	synth byte
}

func (t *instrWaveSynth) set(b byte) {
	t.wave = b
	for i := 7; i >= 4; i-- {
		t.synth = getBit(b, uint(i))
	}
}
