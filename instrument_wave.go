package liblsdj

import "io"

type playbackMode byte

const (
	playOnce playbackMode = iota
	playLoop
	playPingPong
	playManual
)

type instrumentWave struct {
	plVibSpeed   plVibSpeed
	vibShape     vibShape
	vibDirection vibDirection

	transpose byte
	drumMode  byte
	synth     byte
	playback  playbackMode
	length    byte
	repeat    byte
	speed     byte
}

func (i *instrumentWave) clear() {
	i.plVibSpeed = plVibFast
	i.vibShape = vibTriangle
	i.vibDirection = vipUp

	i.transpose = 1
	i.drumMode = 0
	i.synth = 0
	i.playback = playOnce
	i.length = 0x0F
	i.repeat = 0
	i.speed = 4
}

func (i *instrumentWave) read(in *instrument, r io.ReadSeeker, version byte) {
	var b byte

	b, _ = readByte(r)
	i.synth = (b >> 4) & 0xf
	i.repeat = b & 0xf

	// Byte 3 e 4 sono vuoti
	_, _ = r.Seek(2, io.SeekCurrent)

	b, _ = readByte(r)
	i.drumMode = parseDrumMode(b, version)
	i.transpose = parseTranspose(b, version)

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

	// byte 8 è vuoto
	_, _ = r.Seek(1, io.SeekCurrent)

	b, _ = readByte(r)
	i.playback = parsePlaybackMode(b)

	// WAVE length and speed changed in version 6
	if version >= 7 {
		b, _ = readByte(r) // Byte 10
		i.length = 0xf - (b & 0xf)

		b, _ = readByte(r) // Byte 11
		i.speed = b + 4
	} else if version == 6 {
		b, _ = readByte(r) // Byte 10
		i.length = b & 0xf

		b, _ = readByte(r) // Byte 11
		i.speed = b + 1
	} else {
		_, _ = r.Seek(2, io.SeekCurrent) // Bytes 12-13 are empty
	}

	_, _ = r.Seek(2, io.SeekCurrent) // Bytes 10-13 are empty
	b, _ = readByte(r)
	if version < 6 {
		i.length = (b >> 4) & 0xf
		i.speed = (b & 0xf) + 1
	}

	_, _ = r.Seek(1, io.SeekCurrent) // Byte 15 is empty
}

func (i *instrumentWave) write(in *instrument, w io.WriteSeeker, version byte) {
	var b byte

	// byte 0
	_ = writeByte(1, w)
	// byte 1
	_ = writeByte(createWaveVolumeByte(in.envelopeVolume), w)
	// byte 2
	b = ((i.synth & 0xf) << 4) | (i.repeat & 0xf)
	_ = writeByte(b, w)
	// byte 3 empty
	_ = writeByte(0, w)
	// byte 4 empty
	_ = writeByte(0xff, w)

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
	_ = writeByte(createPanningByte(in.panning), w)
	_ = writeByte(0, w)

	_ = writeByte(createPlaybackModeByte(i.playback), w)

	if version >= 7 {
		b = 0xf - (i.length & 0xF)
		_ = writeByte(b, w)

		b = i.speed - 4
		_ = writeByte(b, w)
	} else if version == 6 {
		b = i.length & 0xF
		_ = writeByte(b, w)

		b = i.speed - 1
		_ = writeByte(b, w)
	} else {
		_ = writeByte(0, w) // byte 10 is empty
		_ = writeByte(0, w) // byte 11 is empty
	}

	_ = writeByte(0, w) // byte 12 is empty
	_ = writeByte(0, w) // byte 13 is empty

	if version < 6 {
		b = ((i.length & 0xf) << 4) | ((i.speed - 1) & 0xf)
		_ = writeByte(b, w)
	} else {
		_ = writeByte(0, w) // byte 14 is empty
	}

	_ = writeByte(0, w) // byte 15 is empty
}
