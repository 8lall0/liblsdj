package liblsdj

const lsdj_WAVE_LENGTH int = 16

var lsdj_DEFAULT_WAVE = []byte{0x8E, 0xCD, 0xCC, 0xBB, 0xAA, 0xA9, 0x99, 0x88, 0x87, 0x76, 0x66, 0x55, 0x54, 0x43, 0x32, 0x31} //lsdj_WAVE_LENGTH

// Structure represening a wave for the wave synthesizer
type wave struct {
	data []byte //lsdj_WAVE_LENGTH
}

// clear all wave groove to factory settings
func (w *wave) clear() {
	w.data = make([]byte, lsdj_GROOVE_LENGTH)
	copy(w.data, lsdj_DEFAULT_WAVE)
}
