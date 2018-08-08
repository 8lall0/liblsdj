package liblsdj

type lsdj_pulse_wave byte

const (
	lsdj_PULSE_WAVE_PW_125 lsdj_pulse_wave = 0
	lsdj_PULSE_WAVE_PW_25  lsdj_pulse_wave = 1
	lsdj_PULSE_WAVE_PW_50  lsdj_pulse_wave = 2
	lsdj_PULSE_WAVE_PW_75  lsdj_pulse_wave = 3
)

type pulseT struct {
	name     []byte /*[lsdj_INSTRUMENT_NAME_LENGTH]*/
	insType  int
	panning  panning
	envelope byte // envelope or byte
	table    byte // 0x20 or higher = lsdj_NO_TABLE
	automate byte
	pulse    struct {
		pulseWidth       lsdj_pulse_wave
		length           byte // 0x40 and above = unlimited
		sweep            byte
		plvibSpeed       lsdj_plvib_speed
		vibShape         lsdj_vib_shape
		vibratoDirection lsdj_vib_direction
		transpose        byte
		drumMode         byte
		pulse2tune       byte
		fineTune         byte
	}
}

func (i *pulseT) read(r *vio, ver byte) {
	var b byte

	i.insType = lsdj_INSTR_PULSE
	i.envelope = r.readSingle()
	i.pulse.pulse2tune = r.readSingle()
	i.pulse.length = parseLength(r.readSingle())
	i.pulse.sweep = r.readSingle()

	b = r.readSingle()
	i.pulse.drumMode = parseDrumMode(b, ver)
	i.pulse.transpose = parseTranspose(b, ver)
	i.automate = parseAutomate(b)
	i.pulse.vibratoDirection = lsdj_vib_direction(b & 1)

	if int(ver) < 4 {
		switch int(b>>1) & 3 {
		case 0:
			i.pulse.plvibSpeed = lsdj_PLVIB_FAST
			i.pulse.vibShape = lsdj_VIB_TRIANGLE
		case 1:
			i.pulse.plvibSpeed = lsdj_PLVIB_TICK
			i.pulse.vibShape = lsdj_VIB_SAWTOOTH
		case 2:
			i.pulse.plvibSpeed = lsdj_PLVIB_TICK
			i.pulse.vibShape = lsdj_VIB_TRIANGLE
		case 3:
			i.pulse.plvibSpeed = lsdj_PLVIB_FAST
			i.pulse.vibShape = lsdj_VIB_SQUARE
		}
	} else {
		if b&0x80 == 1 {
			i.pulse.plvibSpeed = lsdj_PLVIB_STEP
		} else if b&0x10 == 1 {
			i.pulse.plvibSpeed = lsdj_PLVIB_TICK
		} else {
			i.pulse.plvibSpeed = lsdj_PLVIB_FAST
		}

		switch int((b >> 1) & 3) {
		case 0:
			i.pulse.vibShape = lsdj_VIB_TRIANGLE
		case 1:
			i.pulse.vibShape = lsdj_VIB_SAWTOOTH
		case 2:
			i.pulse.vibShape = lsdj_VIB_SQUARE
		}
	}

	b = r.readSingle()
	i.table = parseTable(b)
	b = r.readSingle()
	i.pulse.pulseWidth = parsePulseWidth(b)
	i.pulse.fineTune = (b >> 2) & 0xF

	// Bytes 8-15 are empty
	r.seek(r.getCur() + 8)

	// TODO: capire come usare sto assert
	//assert(vio->tell(vio->user_data) - pos == 15);
}

func (i *pulseT) clear() {
	i.insType = lsdj_INSTR_PULSE
	i.envelope = 0xA8
	i.panning = lsdj_PAN_LEFT_RIGHT
	i.table = lsdj_NO_TABLE
	i.automate = 0

	i.pulse.pulseWidth = lsdj_PULSE_WAVE_PW_125
	i.pulse.length = lsdj_INSTRUMENT_UNLIMITED_LENGTH
	i.pulse.sweep = 0xFF
	i.pulse.plvibSpeed = lsdj_PLVIB_FAST
	i.pulse.vibShape = lsdj_VIB_TRIANGLE
	i.pulse.vibratoDirection = lsdj_VIB_UP
	i.pulse.transpose = 1
	i.pulse.drumMode = 0
	i.pulse.pulse2tune = 0
	i.pulse.fineTune = 0
}
