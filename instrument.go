package liblsdj

type lsdj_plvib_speed byte
type lsdj_vib_shape byte
type lsdj_vib_direction byte

const (
	lsdj_INSTRUMENT_NAME_LENGTH         int  = 5
	lsdj_LSDJ_DEFAULT_INSTRUMENT_LENGTH int  = 16
	lsdj_NO_TABLE                       byte = 0x20
	lsdj_INSTRUMENT_UNLIMITED_LENGTH    byte = 0x40
	lsdj_KIT_LENGTH_AUTO                byte = 0x0

	lsdj_PLVIB_FAST   lsdj_plvib_speed   = 0
	lsdj_PLVIB_TICK   lsdj_plvib_speed   = 1
	lsdj_PLVIB_STEP   lsdj_plvib_speed   = 2
	lsdj_VIB_TRIANGLE lsdj_vib_shape     = 0
	lsdj_VIB_SAWTOOTH lsdj_vib_shape     = 1
	lsdj_VIB_SQUARE   lsdj_vib_shape     = 2
	lsdj_VIB_UP       lsdj_vib_direction = 0
	lsdj_VIB_DOWN     lsdj_vib_direction = 1

	lsdj_INSTR_PULSE = iota + 1
	lsdj_INSTR_WAVE
	lsdj_INSTR_KIT
	lsdj_INSTR_NOISE
)

var lsdj_DEFAULT_INSTRUMENT = [lsdj_LSDJ_DEFAULT_INSTRUMENT_LENGTH]byte{0, 0xA8, 0, 0, 0xFF, 0, 0, 3, 0, 0, 0xD0, 0, 0, 0, 0xF3, 0}

type instrument interface {
	read(r *vio, ver byte)
	clear()
}

type instrumentContainerA []*instrumentContainer

type instrumentContainer struct {
	name           []byte //lsdj_INSTRUMENT_NAME_LENGTH]
	instrumentType int
	instrument     instrument
}

func (ic *instrumentContainer) getInstrument(r *vio, version byte) {
	var iType byte

	iType = r.readByte()
	switch int(iType) {
	case 0:
		ic.instrumentType = lsdj_INSTR_PULSE
		ic.instrument = new(pulseT)
	case 1:
		ic.instrumentType = lsdj_INSTR_WAVE
		ic.instrument = new(waveT)
	case 2:
		ic.instrumentType = lsdj_INSTR_KIT
		ic.instrument = new(kitT)
	case 3:
		ic.instrumentType = lsdj_INSTR_NOISE
		ic.instrument = new(noiseT)
	}
	ic.instrument.read(r, version)
}

func (iC *instrumentContainerA) initialize(allocTable []byte) {
	*iC = make([]*instrumentContainer, lsdj_INSTRUMENT_COUNT)
	for i := 0; i < lsdj_INSTRUMENT_COUNT; i++ {
		if allocTable[i] != 0 {
			(*iC)[i] = new(instrumentContainer)
		} else {
			(*iC)[i] = nil
		}
	}
}

func (iC instrumentContainerA) writeName(r *vio) {
	for i := 0; i < lsdj_INSTRUMENT_COUNT; i++ {
		if iC[i] != nil {
			iC[i].name = r.read(lsdj_INSTRUMENT_NAME_LENGTH)
		} else {
			r.seekCur(lsdj_INSTRUMENT_NAME_LENGTH)
		}
	}
}

func (iC instrumentContainerA) writeInstrument(r *vio, version byte) {
	for i := 0; i < lsdj_INSTRUMENT_COUNT; i++ {
		if iC[i] != nil {
			iC[i].getInstrument(r, version)
		} else {
			r.seekCur(lsdj_INSTRUMENT_COUNT)
		}
	}
}

func parseLength(b byte) byte {
	if b&0x40 == 1 {
		return (^b) & 0x3F
	}
	return lsdj_INSTRUMENT_UNLIMITED_LENGTH
}

func parseTable(b byte) byte {
	if b&0x20 == 1 {
		return b & 0x1F
	}
	return lsdj_NO_TABLE
}

func parsePanning(b byte) panning {
	return panning(b & 3)
}

func parseDrumMode(b byte, vers byte) byte {
	if vers < 3 {
		return 0
	}
	if b&0x40 == 1 {
		return 1
	}
	return 0
}

func parseTranspose(b byte, vers byte) byte {
	if vers < 3 {
		return 0
	}
	if b&0x20 == 1 {
		return 1
	}
	return 0
}

func parseAutomate(b byte) byte {
	return (b >> 3) & 0x1
}

func parsePulseWidth(b byte) lsdj_pulse_wave {
	return lsdj_pulse_wave((b >> 6) & 0x3)
}

func parsePlaybackMode(b byte) lsdj_playback_mode {
	return lsdj_playback_mode(b & 0x3)
}

func parseKitDistortion(b byte) lsdj_kit_distortion {
	return lsdj_kit_distortion(b)
}

func parseScommand(b byte) lsdj_scommand_type {
	return lsdj_scommand_type(b & 1)
}
