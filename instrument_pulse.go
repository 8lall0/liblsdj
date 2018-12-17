package liblsdj

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

func (i *instrumentPulse) clear() {
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

func (i *instrumentPulse) read(in *instrument, r io.ReadSeeker) {
	var b byte
	i.pulse2tune, _ = readByte(r)

	b, _ = readByte(r)
	i.length = parseLength(b)

	i.sweep, _ = readByte(r)

	// TODO trova modo di leggere version
	b, _ = readByte(r)
	version := byte(1)
	i.drumMode = parseDrumMode(b, version)
	i.transpose = parseTranspose(b, version)
	i.vibDirection = vibDirection(b & 1)

	// TODO: migliora inserzione in instrument di base???
	in.automate = parseAutomate(b)

	if version < 4 {
		switch (b >> 1) & 3 {
		case 0:
			i.plVibSpeed = plVibFast
			i.vibShape = vibTriangle
		case 1:
			i.plVibSpeed = plVibTick
			i.vibShape = vibSawtooth
		case 2:
			i.plVibSpeed = plVibTick
			i.vibShape = vibTriangle
		case 3:
			i.plVibSpeed = plVibTick
			i.vibShape = vibSquare
		}
	} else {
		if b&0x80 != 0 {
			i.plVibSpeed = plVibStep
		} else if b&0x10 != 0 {
			i.plVibSpeed = plVibTick
		} else {
			i.plVibSpeed = plVibFast
		}

		switch (b >> 1) & 3 {
		case 0:
			i.vibShape = vibTriangle
		case 1:
			i.vibShape = vibSawtooth
		case 2:
			i.vibShape = vibSquare
		}
	}
	b, _ = readByte(r)
	in.table = parseTable(b)

	b, _ = readByte(r)
	in.panning = parsePanning(b)
	i.pulseWidth = parsePulseWidth(b)
	i.fineTune = (b >> 2) & 0xf
	_, _ = r.Seek(8, io.SeekCurrent)
}
