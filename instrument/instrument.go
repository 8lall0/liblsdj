package instrument

import (
	"github.com/8lall0/liblsdj/panning"
)

const (
	LSDJ_INSTRUMENT_NAME_LENGTH         int = 5
	LSDJ_LSDJ_DEFAULT_INSTRUMENT_LENGTH int = 16

	LSDJ_NO_TABLE                    byte = 0x20
	LSDJ_INSTRUMENT_UNLIMITED_LENGTH byte = 0x40
	LSDJ_KIT_LENGTH_AUTO             byte = 0x0
)

var LSDJ_DEFAULT_INSTRUMENT = [LSDJ_LSDJ_DEFAULT_INSTRUMENT_LENGTH]byte{0, 0xA8, 0, 0, 0xFF, 0, 0, 3, 0, 0, 0xD0, 0, 0, 0, 0xF3, 0}

const (
	LSDJ_INSTR_PULSE = iota + 1
	LSDJ_INSTR_WAVE
	LSDJ_INSTR_KIT
	LSDJ_INSTR_NOISE
)

// Structure representing one instrument
type Instrument struct {
	name/*[LSDJ_INSTRUMENT_NAME_LENGTH]*/ string
	instrument_type int

	panning panning.Panning

	envVol struct {
		envelope byte
		volume   byte
	}

	table    byte // 0x20 or higher = LSDJ_NO_TABLE
	automate byte

	instrument struct {
		pulse pulseT
		wave  waveT
		kit   kitT
		noise noiseT
	}
}

// Copy a instrument
func (dest *Instrument) CopyFrom(src *Instrument) {
	tmp := *src
	dest = &tmp
}

// Clear all instrument data to factory settings
func (ins *Instrument) ClearInstrument() {
	ins.name = ""
	ins.ClearPulse()
}

func (ins *Instrument) ClearPulse() {
	ins.instrument_type = LSDJ_INSTR_PULSE
	ins.envVol.envelope = 0xA8
	ins.panning = panning.LSDJ_PAN_LEFT_RIGHT
	ins.table = LSDJ_NO_TABLE
	ins.automate = 0

	ins.instrument.pulse.pulseWidth = LSDJ_PULSE_WAVE_PW_125
	ins.instrument.pulse.length = LSDJ_INSTRUMENT_UNLIMITED_LENGTH
	ins.instrument.pulse.sweep = 0xFF
	ins.instrument.pulse.plvibSpeed = LSDJ_PLVIB_FAST
	ins.instrument.pulse.vibShape = LSDJ_VIB_TRIANGLE
	ins.instrument.pulse.vibratoDirection = LSDJ_VIB_UP
	ins.instrument.pulse.transpose = 1
	ins.instrument.pulse.drumMode = 0
	ins.instrument.pulse.pulse2tune = 0
	ins.instrument.pulse.fineTune = 0
}

func (ins *Instrument) ClearWave() {
	ins.instrument_type = LSDJ_INSTR_WAVE
	ins.envVol.volume = 3
	ins.panning = panning.LSDJ_PAN_LEFT_RIGHT
	ins.table = LSDJ_NO_TABLE
	ins.automate = 0

	ins.instrument.wave.plvibSpeed = LSDJ_PLVIB_FAST
	ins.instrument.wave.vibShape = LSDJ_VIB_TRIANGLE
	ins.instrument.wave.vibratoDirection = LSDJ_VIB_UP
	ins.instrument.wave.transpose = 1
	ins.instrument.wave.drumMode = 0
	ins.instrument.wave.synth = 0
	ins.instrument.wave.playback = LSDJ_PLAY_ONCE
	ins.instrument.wave.length = 0x0F
	ins.instrument.wave.repeat = 0
	ins.instrument.wave.speed = 4
}

func (ins *Instrument) ClearKit() {
	ins.instrument_type = LSDJ_INSTR_KIT
	ins.envVol.volume = 3
	ins.panning = panning.LSDJ_PAN_LEFT_RIGHT
	ins.table = LSDJ_NO_TABLE
	ins.automate = 0

	ins.instrument.kit.kit1 = 0
	ins.instrument.kit.offset1 = 0
	ins.instrument.kit.length1 = LSDJ_KIT_LENGTH_AUTO
	ins.instrument.kit.loop1 = LSDJ_KIT_LOOP_OFF

	ins.instrument.kit.kit2 = 0
	ins.instrument.kit.offset2 = 0
	ins.instrument.kit.length2 = LSDJ_KIT_LENGTH_AUTO
	ins.instrument.kit.loop1 = LSDJ_KIT_LOOP_OFF

	ins.instrument.kit.pitch = 0
	ins.instrument.kit.halfSpeed = 0
	ins.instrument.kit.distortion = LSDJ_KIT_DIST_CLIP
	ins.instrument.kit.plvibSpeed = LSDJ_PLVIB_FAST
	ins.instrument.kit.vibShape = LSDJ_VIB_TRIANGLE
}

func (ins *Instrument) ClearNoise() {
	ins.instrument_type = LSDJ_INSTR_NOISE
	ins.envVol.envelope = 0xA8
	ins.panning = panning.LSDJ_PAN_LEFT_RIGHT
	ins.table = LSDJ_NO_TABLE
	ins.automate = 0

	ins.instrument.noise.length = LSDJ_INSTRUMENT_UNLIMITED_LENGTH
	ins.instrument.noise.shape = 0xFF
	ins.instrument.noise.sCommand = LSDJ_SCOMMAND_FREE
}

func parseLength(b byte) byte {
	if b&0x40 == 1 {
		return (^b) & 0x3F
	}
	return LSDJ_INSTRUMENT_UNLIMITED_LENGTH
}

func parseTable(b byte) byte {
	if b&0x20 == 1 {
		return b & 0x1F
	}
	return LSDJ_NO_TABLE
}

func parsePanning(b byte) panning.Panning {
	return panning.Panning(b & 3)
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

func (ins *Instrument) SetName(name string) {
	ins.name = name
}

func (ins *Instrument) GetName() string {
	return ins.name
}

func (ins *Instrument) SetPanning(pan panning.Panning) {
	ins.panning = pan
}

func (ins *Instrument) GetPanning() panning.Panning {
	return ins.panning
}

// Instrument I/O
/*void lsdj_instrument_read(lsdj_vio_t* vio, unsigned char version, lsdj_instrument_t* instrument, lsdj_error_t** error);
void lsdj_instrument_write(const lsdj_instrument_t* instrument, unsigned char version, lsdj_vio_t* vio, lsdj_error_t** error);*/
