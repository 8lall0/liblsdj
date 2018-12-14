package lsdj

import "io"

type pulseWave byte

const (
	pulseWave125 pulseWave = iota
	pulseWave25
	pulseWave50
	pulseWave75
)

type instrumentPulse struct {
	pulseWidth pulseWave
	length     byte // 0x40 and above = unlimited
	sweep      byte

	plVibSpeed   plVibSpeed
	vibShape     vibShape
	vibDirection vibDirection

	transpose  byte
	drumMode   byte
	pulse2tune byte
	fineTune   byte
}

func (i *instrumentPulse) clearInstrument() {
	i.pulseWidth = pulseWave125
	i.length = instrumentUnlimitedLength
	i.sweep = 0xFF

	i.plVibSpeed = plVibFast
	i.vibShape = vibTriangle
	i.vibDirection = vipUp

	i.transpose = 1
	i.drumMode = 0
	i.pulse2tune = 0
	i.fineTune = 0
}

func (i *instrumentPulse) read(r io.ReadSeeker) {
	//TODO: read_pulse_instrument
}
