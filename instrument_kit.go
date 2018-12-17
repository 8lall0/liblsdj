package liblsdj

import "io"

type kitLoopMode byte
type kitDistortion byte
type kitPspeed byte

const (
	kitDistClip   kitDistortion = 0xD0
	kitDistShape  kitDistortion = 0xD1
	kitDistShape2 kitDistortion = 0xD2
	kitDistWrap   kitDistortion = 0xD3

	kitLoopOff kitLoopMode = iota
	kitLoopOn
	kitLoopAttack

	kitPspeedFast kitPspeed = iota
	kitPspeedSlow
	kitPspeedStep
)

type instrumentKit struct {
	kit1    byte
	offset1 byte
	length1 byte
	loop1   kitLoopMode

	kit2    byte
	offset2 byte
	length2 byte
	loop2   kitLoopMode

	pitch        byte
	halfSpeed    byte
	distortion   kitDistortion
	plVibSpeed   plVibSpeed
	vibShape     vibShape
	vibDirection vibDirection
}

func (i *instrumentKit) clear() {
	i.kit1 = 0
	i.offset1 = 0
	i.length1 = kitLengthAuto
	i.loop1 = kitLoopOff

	i.kit2 = 0
	i.offset2 = 0
	i.length2 = kitLengthAuto
	i.loop2 = kitLoopOff

	i.pitch = 0
	i.halfSpeed = 0
	i.distortion = kitDistClip
	i.plVibSpeed = plVibFast
	i.vibShape = vibTriangle
}

func (i *instrumentKit) read(in *instrument, r io.ReadSeeker) {
	var b byte

	i.loop1 = kitLoopOff
	i.loop2 = kitLoopOff

	b, _ = readByte(r)
	// TODO controlla condizioni
	if (b>>7)&1 != 0 {
		i.loop1 = kitLoopAttack
	}
	i.halfSpeed = (b >> 6) & 1
	i.kit1 = b & 0x3f
	i.length1, _ = readByte(r)

	// Byte 4 vuoto
	_, _ = r.Seek(1, io.SeekCurrent)
	b, _ = readByte(r)

	if i.loop1 != kitLoopAttack {
		if b&0x40 != 0 {
			i.loop1 = kitLoopOn
		} else {
			i.loop1 = kitLoopOff
		}
	}

	if b&0x20 != 0 {
		i.loop2 = kitLoopOn
	} else {
		i.loop2 = kitLoopOff
	}

	in.automate = parseAutomate(b)

	// TODO: continua da qui

	version := byte(1)

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
