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

func (i *instrumentKit) clearInstrument() {
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

func (i *instrumentKit) read(r io.ReadSeeker) {
	//TODO: read_kit_instrument
}
