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

// Structure representing one instrument
type instrument struct {
	name            []byte /*[lsdj_INSTRUMENT_NAME_LENGTH]*/
	instrument_type int

	panning panning

	envVol struct {
		envelope byte
		volume   byte
	}

	table    byte // 0x20 or higher = lsdj_NO_TABLE
	automate byte

	instrument struct {
		pulse pulseT
		wave  waveT
		kit   kitT
		noise noiseT
	}
}

// copy a instrument
func (i *instrument) Copy() *instrument {
	return &(*i)
}

// clear all instrument groove to factory settings
func (ins *instrument) clear() {
	ins.name = ""
	ins.clearPulse()
}

func (ins *instrument) clearPulse() {
	ins.instrument_type = lsdj_INSTR_PULSE
	ins.envVol.envelope = 0xA8
	ins.panning = lsdj_PAN_LEFT_RIGHT
	ins.table = lsdj_NO_TABLE
	ins.automate = 0

	ins.instrument.pulse.pulseWidth = lsdj_PULSE_WAVE_PW_125
	ins.instrument.pulse.length = lsdj_INSTRUMENT_UNLIMITED_LENGTH
	ins.instrument.pulse.sweep = 0xFF
	ins.instrument.pulse.plvibSpeed = lsdj_PLVIB_FAST
	ins.instrument.pulse.vibShape = lsdj_VIB_TRIANGLE
	ins.instrument.pulse.vibratoDirection = lsdj_VIB_UP
	ins.instrument.pulse.transpose = 1
	ins.instrument.pulse.drumMode = 0
	ins.instrument.pulse.pulse2tune = 0
	ins.instrument.pulse.fineTune = 0
}

func (ins *instrument) clearWave() {
	ins.instrument_type = lsdj_INSTR_WAVE
	ins.envVol.volume = 3
	ins.panning = lsdj_PAN_LEFT_RIGHT
	ins.table = lsdj_NO_TABLE
	ins.automate = 0

	ins.instrument.wave.plvibSpeed = lsdj_PLVIB_FAST
	ins.instrument.wave.vibShape = lsdj_VIB_TRIANGLE
	ins.instrument.wave.vibratoDirection = lsdj_VIB_UP
	ins.instrument.wave.transpose = 1
	ins.instrument.wave.drumMode = 0
	ins.instrument.wave.synth = 0
	ins.instrument.wave.playback = lsdj_PLAY_ONCE
	ins.instrument.wave.length = 0x0F
	ins.instrument.wave.repeat = 0
	ins.instrument.wave.speed = 4
}

func (ins *instrument) clearKit() {
	ins.instrument_type = lsdj_INSTR_KIT
	ins.envVol.volume = 3
	ins.panning = lsdj_PAN_LEFT_RIGHT
	ins.table = lsdj_NO_TABLE
	ins.automate = 0

	ins.instrument.kit.kit1 = 0
	ins.instrument.kit.offset1 = 0
	ins.instrument.kit.length1 = lsdj_KIT_LENGTH_AUTO
	ins.instrument.kit.loop1 = lsdj_KIT_LOOP_OFF

	ins.instrument.kit.kit2 = 0
	ins.instrument.kit.offset2 = 0
	ins.instrument.kit.length2 = lsdj_KIT_LENGTH_AUTO
	ins.instrument.kit.loop1 = lsdj_KIT_LOOP_OFF

	ins.instrument.kit.pitch = 0
	ins.instrument.kit.halfSpeed = 0
	ins.instrument.kit.distortion = lsdj_KIT_DIST_CLIP
	ins.instrument.kit.plvibSpeed = lsdj_PLVIB_FAST
	ins.instrument.kit.vibShape = lsdj_VIB_TRIANGLE
}

func (ins *instrument) clearNoise() {
	ins.instrument_type = lsdj_INSTR_NOISE
	ins.envVol.envelope = 0xA8
	ins.panning = lsdj_PAN_LEFT_RIGHT
	ins.table = lsdj_NO_TABLE
	ins.automate = 0

	ins.instrument.noise.length = lsdj_INSTRUMENT_UNLIMITED_LENGTH
	ins.instrument.noise.shape = 0xFF
	ins.instrument.noise.sCommand = lsdj_SCOMMAND_FREE
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

func (ins *instrument) SetName(name string) {
	ins.name = name
}

func (ins *instrument) GetName() string {
	return ins.name
}

func (ins *instrument) SetPanning(pan panning) {
	ins.panning = pan
}

func (ins *instrument) GetPanning() panning {
	return ins.panning
}

// instrument I/O
/*void lsdj_instrument_read(lsdj_vio_t* vio, unsigned char version, lsdj_instrument_t* instrument, lsdj_error_t** error);
void lsdj_instrument_write(const lsdj_instrument_t* instrument, unsigned char version, lsdj_vio_t* vio, lsdj_error_t** error);*/
