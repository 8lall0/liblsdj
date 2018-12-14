package lsdj

import "io"

type sCommand byte

const (
	sCommandFree sCommand = iota
	sCommandStable
)

type instrumentNoise struct {
	length   byte
	shape    byte
	sCommand sCommand
}

func (i *instrumentNoise) clearInstrument() {
	i.length = instrumentUnlimitedLength
	i.shape = 0xFF
	i.sCommand = sCommandFree
}

func (i *instrumentNoise) read(r io.ReadSeeker) {
	//TODO: read_noise_instrument
}
