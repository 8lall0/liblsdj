package lsdj

type channel byte

const (
	pulse1T channel = iota
	pulse2T
	waveT
	noiseT
)
