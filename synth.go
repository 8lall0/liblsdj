package liblsdj

const (
	LSDJ_SYNTH_WAVEFORM_SAWTOOTH byte = 0
	LSDJ_SYNTH_WAVEFORM_SQUARE   byte = 1
	LSDJ_SYNTH_WAVEFORM_TRIANGLE byte = 2

	LSDJ_SYNTH_FILTER_LOW_PASS  byte = 0
	LSDJ_SYNTH_FILTER_HIGH_PASS byte = 1
	LSDJ_SYNTH_FILTER_BAND_PASS byte = 2
	LSDJ_SYNTH_FILTER_ALL_PASS  byte = 3

	LSDJ_SYNTH_DISTORTION_CLIP byte = 0
	LSDJ_SYNTH_DISTORTION_WRAP byte = 1
	LSDJ_SYNTH_DISTORTION_FOLD byte = 2

	LSDJ_SYNTH_PHASE_NORMAL  byte = 0
	LSDJ_SYNTH_PHASE_RESYNC  byte = 1
	LSDJ_SYNTH_PHASE_RESYNC2 byte = 2
)

// Structure representing soft synth data
type Lsdj_synth_t struct {
	waveform       byte
	filter         byte
	resonanceStart byte
	resonanceEnd   byte
	distortion     byte
	phase          byte

	volumeStart byte
	volumeEnd   byte
	cutOffStart byte
	cutOffEnd   byte

	phaseStart  byte
	phaseEnd    byte
	vshiftStart byte
	vshiftEnd   byte

	limitStart byte
	limitEnd   byte

	reserved [2]byte

	overwritten bool
}

// Clear all soft synth data to factory settings
func (synth Lsdj_synth_t) Clear() {
	synth.waveform = LSDJ_SYNTH_WAVEFORM_SAWTOOTH
	synth.filter = LSDJ_SYNTH_FILTER_LOW_PASS
	synth.resonanceStart = 0
	synth.resonanceEnd = 0
	synth.distortion = LSDJ_SYNTH_DISTORTION_CLIP
	synth.phase = LSDJ_SYNTH_PHASE_NORMAL
	synth.volumeStart = 0x10
	synth.cutOffStart = 0xFF
	synth.phaseStart = 0
	synth.vshiftStart = 0
	synth.volumeEnd = 0x10
	synth.cutOffEnd = 0xFF
	synth.phaseEnd = 0
	synth.vshiftEnd = 0
	synth.limitStart = 0xF
	synth.limitEnd = 0xF
	synth.reserved[0] = 0
	synth.reserved[1] = 0

	synth.overwritten = false
}
