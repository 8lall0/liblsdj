package synth

// Structure representing soft synth data
type Synth struct {
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
func (synth *Synth) Clear() {
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
