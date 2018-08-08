package liblsdj

const (
	lsdj_SYNTH_WAVEFORM_SAWTOOTH byte = 0
	lsdj_SYNTH_WAVEFORM_SQUARE   byte = 1
	lsdj_SYNTH_WAVEFORM_TRIANGLE byte = 2

	lsdj_SYNTH_FILTER_LOW_PASS  byte = 0
	lsdj_SYNTH_FILTER_HIGH_PASS byte = 1
	lsdj_SYNTH_FILTER_BAND_PASS byte = 2
	lsdj_SYNTH_FILTER_ALL_PASS  byte = 3

	lsdj_SYNTH_DISTORTION_CLIP byte = 0
	lsdj_SYNTH_DISTORTION_WRAP byte = 1
	lsdj_SYNTH_DISTORTION_FOLD byte = 2

	lsdj_SYNTH_PHASE_NORMAL  byte = 0
	lsdj_SYNTH_PHASE_RESYNC  byte = 1
	lsdj_SYNTH_PHASE_RESYNC2 byte = 2
)

// Structure representing soft synth groove
type synth struct {
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

// clear all soft synth groove to factory settings
func (s *synth) clear() {
	s.waveform = lsdj_SYNTH_WAVEFORM_SAWTOOTH
	s.filter = lsdj_SYNTH_FILTER_LOW_PASS
	s.resonanceStart = 0
	s.resonanceEnd = 0
	s.distortion = lsdj_SYNTH_DISTORTION_CLIP
	s.phase = lsdj_SYNTH_PHASE_NORMAL
	s.volumeStart = 0x10
	s.cutOffStart = 0xFF
	s.phaseStart = 0
	s.vshiftStart = 0
	s.volumeEnd = 0x10
	s.cutOffEnd = 0xFF
	s.phaseEnd = 0
	s.vshiftEnd = 0
	s.limitStart = 0xF
	s.limitEnd = 0xF
	s.reserved[0] = 0
	s.reserved[1] = 0
	s.overwritten = false
}
