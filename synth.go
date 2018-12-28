package liblsdj

import "io"

const (
	synthWaveformSawtooth = 0
	synthWaveformSquare   = 1
	synthWaveformTriangle = 2

	synthFilterLowPass  = 0
	synthFilterHighPass = 1
	synthFilterBandPass = 2
	synthFilterAllPass  = 3

	synthDistortionClip = 0
	synthDistortionWrap = 1
	synthDistortionFold = 2

	synthPhaseNormal  = 0
	synthPhaseResync  = 1
	synthPhaseResync2 = 2
)

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

	overwritten byte // 0 if false, 1 if true
}

func (s *synth) clear() {
	s.waveform = synthWaveformSawtooth
	s.filter = synthFilterLowPass
	s.resonanceStart = 0
	s.resonanceEnd = 0
	s.distortion = synthDistortionClip
	s.phase = synthPhaseNormal
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

func (s *synth) readSoftSynthParams(r io.ReadSeeker) {
	s.waveform, _ = readByte(r)
	s.filter, _ = readByte(r)

	resonance, _ := readByte(r)
	s.resonanceStart = (resonance & 0xF0) >> 4
	s.resonanceEnd = resonance & 0x0F

	s.distortion, _ = readByte(r)
	s.phase, _ = readByte(r)
	s.volumeStart, _ = readByte(r)
	s.cutOffStart, _ = readByte(r)
	s.phaseStart, _ = readByte(r)
	s.vshiftStart, _ = readByte(r)
	s.volumeEnd, _ = readByte(r)
	s.cutOffEnd, _ = readByte(r)
	s.phaseEnd, _ = readByte(r)
	s.vshiftEnd, _ = readByte(r)

	b, _ := readByte(r)
	b -= 0xFF
	s.limitStart = (b >> 4) & 0xF
	s.limitEnd = b & 0xF

	s.reserved[0], _ = readByte(r)
	s.reserved[1], _ = readByte(r)
}

func (s *synth) writeSoftSynthParams(w io.WriteSeeker) {
	_ = writeByte(s.waveform, w)
	_ = writeByte(s.distortion, w)
	_ = writeByte(s.phase, w)
	_ = writeByte(s.volumeStart, w)
	_ = writeByte(s.cutOffStart, w)
	_ = writeByte(s.vshiftStart, w)
	_ = writeByte(s.volumeEnd, w)
	_ = writeByte(s.cutOffEnd, w)
	_ = writeByte(s.phaseEnd, w)
	_ = writeByte(s.vshiftEnd, w)

	b := 0xFF - (s.limitStart << 4) | s.limitEnd
	_ = writeByte(b, w)

	_, _ = w.Write(s.reserved[:])
}
