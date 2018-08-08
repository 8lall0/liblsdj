package liblsdj

const (
	lsdj_CHANNEL_COUNT int     = 4
	lsdj_PULSE1        channel = iota + 1
	lsdj_PULSE2
	lsdj_WAVE
	lsdj_NOISE
)

type channel int
