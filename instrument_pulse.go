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

func (i *instrumentPulse) write(in *instrument, w io.WriteSeeker, version byte) {
	var b byte

	b = 0

	_ = writeByte(b, w)
	_ = writeByte(in.envelopeVolume, w)
	_ = writeByte(i.pulse2tune, w)

	_ = writeByte(createLengthByte(i.length), w)
	_ = writeByte(i.sweep, w)

	b = createDrumModeByte(i.drumMode, version)
	b |= createTransposeByte(i.transpose, version)
	b |= createAutomateByte(in.automate)
	b |= createVibratoDirectionByte(i.vibDirection)

	if version < 4 {
		switch i.vibShape {
		case vibSawtooth:
			b |= 2
		case vibSquare:
			b |= 6
		case vibTriangle:
			if i.plVibSpeed != plVibFast {
				b |= 4
			}
		}
	} else {
		b |= (byte(i.vibShape) & 3) << 1
		if i.plVibSpeed == plVibTick {
			b |= 0x10
		} else if i.plVibSpeed == plVibStep {
			b |= 0x80
		}
	}

	_ = writeByte(b, w)
	_ = writeByte(createTableByte(in.table), w)

	b = createPulseWidthByte(i.pulseWidth) | ((byte(i.fineTune) & 0xf) << 2) | createPanningByte(in.panning)
	_ = writeByte(b, w)

	empty := []byte{0, 0, 0xD0, 0, 0, 0, 0xF3, 0}
	_, _ = w.Write(empty)
}
