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

func (i *instrumentKit) read(in *instrument, r io.ReadSeeker, version byte) {
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
	i.vibDirection = vibDirection(b & 1)

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

	// byte 6
	b, _ = readByte(r)
	in.table = parseTable(b)

	// byte 7
	b, _ = readByte(r)
	in.panning = parsePanning(b)

	// byte 8
	i.pitch, _ = readByte(r)

	// byte 9
	b, _ = readByte(r)
	if (b>>7)&1 != 0 {
		i.loop2 = kitLoopAttack
	}
	i.kit2 = b & 0x3f

	// byte 10
	b, _ = readByte(r)
	i.distortion = parseKitDistortion(b)
	// byte 11
	i.length2, _ = readByte(r)
	// byte 12
	i.offset1, _ = readByte(r)
	// byte 13
	i.offset2, _ = readByte(r)

	_, _ = r.Seek(2, io.SeekCurrent) // Bytes 14 and 15 are empty
}

func (i *instrumentKit) write(in *instrument, w io.WriteSeeker, version byte) {
	var b byte

	_ = writeByte(2, w)
	_ = writeByte(createWaveVolumeByte(in.envelopeVolume), w)

	// An advice: ternary operator sucks
	if i.loop1 == kitLoopAttack {
		b = 0x80
	} else {
		b = 0x0
	}
	if i.halfSpeed == 1 {
		b = b | 0x40
	} else {
		b = b | 0x0
	}
	b = b | (i.kit1 & 0x3f)
	_ = writeByte(b, w)

	_ = writeByte(i.length1, w)
	_ = writeByte(0xff, w)

	if i.loop1 == kitLoopOn {
		b = 0x40
	} else {
		b = 0x0
	}

	// Byte 5
	if i.loop2 == kitLoopOn {
		b = b | 0x40
	} else {
		b = b | 0x0
	}
	b = b | createAutomateByte(in.automate)

	if version < 4 {
		b = b | (byte(i.plVibSpeed)&3)<<1
	} else {
		switch i.plVibSpeed {
		// plVibFast does nothing
		case plVibTick:
			b = b | 0x10
		case plVibStep:
			b = b | 0x80
		}
	}
	_ = writeByte(b, w)

	_ = writeByte(createTableByte(in.table), w)
	_ = writeByte(createPanningByte(in.panning), w)
	_ = writeByte(i.pitch, w)

	if i.loop2 == kitLoopAttack {
		b = 0x80
	} else {
		b = 0x0
	}
	b = b | (i.kit2 & 0x3f)
	_ = writeByte(b, w)

	_ = writeByte(createKitDistortionByte(i.distortion), w)
	_ = writeByte(i.length2, w)
	_ = writeByte(i.offset1, w)
	_ = writeByte(i.offset2, w)
	_ = writeByte(0xf3, w) // Byte 14 is empty
	_ = writeByte(0, w)    // Byte 15 is empty
}
