package liblsdj

type lsdj_kit_loop_mode byte
type lsdj_kit_distortion byte
type lsdj_kit_pspeed byte

const (
	lsdj_KIT_LOOP_OFF    lsdj_kit_loop_mode = 0
	lsdj_KIT_LOOP_ON     lsdj_kit_loop_mode = 1
	lsdj_KIT_LOOP_ATTACK lsdj_kit_loop_mode = 2

	lsdj_KIT_DIST_CLIP   lsdj_kit_distortion = 0xD0
	lsdj_KIT_DIST_SHAPE  lsdj_kit_distortion = 0xD1
	lsdj_KIT_DIST_SHAPE2 lsdj_kit_distortion = 0xD2
	lsdj_KIT_DIST_WRAP   lsdj_kit_distortion = 0xD3

	lsdj_KIT_PSPEED_FAST lsdj_kit_pspeed = 0
	lsdj_KIT_PSPEED_SLOW lsdj_kit_pspeed = 1
	lsdj_KIT_PSPEED_STEP lsdj_kit_pspeed = 2
)

type kitT struct {
	name     []byte //lsdj_INSTRUMENT_NAME_LENGTH]
	insType  int
	panning  panning
	volume   byte
	table    byte // 0x20 or higher = lsdj_NO_TABLE
	automate byte
	kit      struct {
		kit1    byte
		offset1 byte
		length1 byte
		loop1   lsdj_kit_loop_mode

		kit2    byte
		offset2 byte
		length2 byte
		loop2   lsdj_kit_loop_mode

		pitch            byte
		halfSpeed        byte
		distortion       lsdj_kit_distortion
		plvibSpeed       lsdj_plvib_speed
		vibShape         lsdj_vib_shape
		vibratoDirection lsdj_vib_direction
	}
}

func (i *kitT) read(r *vio, ver byte) {
	var b byte

	i.insType = lsdj_INSTR_KIT
	i.volume = r.readSingle()

	i.kit.loop1 = lsdj_KIT_LOOP_OFF
	i.kit.loop2 = lsdj_KIT_LOOP_OFF

	b = r.readSingle()
	if (b>>7)&1 == 1 {
		i.kit.loop1 = lsdj_KIT_LOOP_ATTACK
	}
	i.kit.halfSpeed = (b >> 6) & 1
	i.kit.kit1 = b & 0x3F
	i.kit.length1 = r.readSingle()

	// Byte 4 is empty
	r.seek(r.getCur() + 1)

	b = r.readSingle()
	if i.kit.loop1 != lsdj_KIT_LOOP_ATTACK {
		if (b & 0x40) == 1 {
			i.kit.loop1 = lsdj_KIT_LOOP_ON
		} else {
			i.kit.loop1 = lsdj_KIT_LOOP_OFF
		}
	}
	if (b & 0x40) == 1 {
		i.kit.loop2 = lsdj_KIT_LOOP_ON
	} else {
		i.kit.loop2 = lsdj_KIT_LOOP_OFF
	}
	i.automate = parseAutomate(b)
	i.kit.vibratoDirection = lsdj_vib_direction(b & 1)

	if int(ver) < 4 {
		switch int(b>>1) & 3 {
		case 0:
			i.kit.plvibSpeed = lsdj_PLVIB_FAST
			i.kit.vibShape = lsdj_VIB_TRIANGLE
		case 1:
			i.kit.plvibSpeed = lsdj_PLVIB_TICK
			i.kit.vibShape = lsdj_VIB_TRIANGLE
		case 2:
			i.kit.plvibSpeed = lsdj_PLVIB_STEP
			i.kit.vibShape = lsdj_VIB_TRIANGLE
		}
	} else {
		if b&0x80 == 1 {
			i.kit.plvibSpeed = lsdj_PLVIB_STEP
		} else if b&0x10 == 1 {
			i.kit.plvibSpeed = lsdj_PLVIB_TICK
		} else {
			i.kit.plvibSpeed = lsdj_PLVIB_FAST
		}

		switch int((b >> 1) & 3) {
		case 0:
			i.kit.vibShape = lsdj_VIB_TRIANGLE
		case 1:
			i.kit.vibShape = lsdj_VIB_SAWTOOTH
		case 2:
			i.kit.vibShape = lsdj_VIB_SQUARE
		}
	}

	i.table = parseTable(r.readSingle())
	i.panning = parsePanning(r.readSingle())
	i.kit.pitch = r.readSingle()

	b = r.readSingle()
	if (b>>7)&1 == 1 {
		i.kit.loop2 = lsdj_KIT_LOOP_ATTACK
	}
	i.kit.kit2 = b & 0x3F

	i.kit.distortion = parseKitDistortion(r.readSingle())
	i.kit.length2 = r.readSingle()
	i.kit.offset1 = r.readSingle()
	i.kit.offset2 = r.readSingle()

	r.seek(r.getCur() + 2)
}

func (i *kitT) clear() {
	i.insType = lsdj_INSTR_KIT
	i.volume = 3
	i.panning = lsdj_PAN_LEFT_RIGHT
	i.table = lsdj_NO_TABLE
	i.automate = 0

	i.kit.kit1 = 0
	i.kit.offset1 = 0
	i.kit.length1 = lsdj_KIT_LENGTH_AUTO
	i.kit.loop1 = lsdj_KIT_LOOP_OFF

	i.kit.kit2 = 0
	i.kit.offset2 = 0
	i.kit.length2 = lsdj_KIT_LENGTH_AUTO
	i.kit.loop1 = lsdj_KIT_LOOP_OFF

	i.kit.pitch = 0
	i.kit.halfSpeed = 0
	i.kit.distortion = lsdj_KIT_DIST_CLIP
	i.kit.plvibSpeed = lsdj_PLVIB_FAST
	i.kit.vibShape = lsdj_VIB_TRIANGLE
}

func (i *kitT) setName(name []byte) {
	i.name = name
}
