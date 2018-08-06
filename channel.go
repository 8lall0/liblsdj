package liblsdj

type Lsdj_channel_t int

const (
	LSDJ_CHANNEL_COUNT int            = 4
	LSDJ_PULSE1        Lsdj_channel_t = iota + 1
	LSDJ_PULSE2
	LSDJ_WAVE
	LSDJ_NOISE
)
