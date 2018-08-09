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

type synthA []*synth

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

	reserved []byte //2

	overwritten byte
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
	s.overwritten = 0
}

func (s *synth) readSoftSynthParam(r *vio) {
	s.waveform = r.readByte()
	s.filter = r.readByte()

	resonance := r.readByte()
	s.resonanceStart = (resonance & 0xF0) >> 4
	s.resonanceEnd = resonance & 0x0F

	s.distortion = r.readByte()
	s.phase = r.readByte()

	s.volumeStart = r.readByte()
	s.cutOffStart = r.readByte()
	s.phaseStart = r.readByte()
	s.vshiftStart = r.readByte()

	s.volumeEnd = r.readByte()
	s.cutOffEnd = r.readByte()
	s.phaseEnd = r.readByte()
	s.vshiftEnd = r.readByte()

	byte := 0xFF - r.readByte()
	s.limitStart = (byte >> 4) & 0xF
	s.limitEnd = byte & 0xF

	s.reserved = r.read(2)
}

func (s *synth) writeSoftSynthParam(w *vio) {
	w.writeByte(s.waveform)
	w.writeByte(s.filter)
	w.writeByte(((s.resonanceStart & 0x0F) << 4) | (s.resonanceEnd & 0x0F)) //resonance
	w.writeByte(s.distortion)
	w.writeByte(s.phase)
	w.writeByte(s.volumeStart)
	w.writeByte(s.cutOffStart)
	w.writeByte(s.phaseStart)
	w.writeByte(s.vshiftStart)
	w.writeByte(s.volumeEnd)
	w.writeByte(s.cutOffEnd)
	w.writeByte(s.phaseEnd)
	w.writeByte(s.vshiftEnd)
	w.writeByte(0xFF - ((s.limitStart << 4) | s.limitEnd))
	w.write(s.reserved)
}
func (s *synthA) initialize() {
	*s = make([]*synth, lsdj_SYNTH_COUNT)
	for i := 0; i < lsdj_SYNTH_COUNT; i++ {
		(*s)[i] = new(synth)
	}
}

func (s synthA) writeSynthParam(w *vio) {
	for i := 0; i < lsdj_SYNTH_COUNT; i++ {
		s[i].writeSoftSynthParam(w)
	}
}

func (s synthA) readSynthParam(r *vio) {
	for i := 0; i < lsdj_SYNTH_COUNT; i++ {
		s[i].readSoftSynthParam(r)
	}
}

func (s synthA) readSynthOverwritten(r *vio) {
	var i uint8

	waveSynthOverwriteLocks := r.read(2)
	for i = 0; i < uint8(lsdj_SYNTH_COUNT); i++ {
		s[i].overwritten = (waveSynthOverwriteLocks[1-(i/8)] >> (i % 8)) & 1
	}
}
