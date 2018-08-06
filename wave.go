package liblsdj

const LSDJ_WAVE_LENGTH int = 16
var LSDJ_DEFAULT_WAVE [LSDJ_WAVE_LENGTH]byte = [0x8E, 0xCD, 0xCC, 0xBB, 0xAA, 0xA9, 0x99, 0x88, 0x87, 0x76, 0x66, 0x55, 0x54, 0x43, 0x32, 0x31]

// Structure represening a wave for the wave synthesizer
type Lsdj_wave_t struct {
	data [LSDJ_WAVE_LENGTH]byte
}

// Clear all wave data to factory settings
func (wave Lsdj_wave_t) Clear() {
	for i:=0;i<LSDJ_WAVE_LENGTH;i++ {
		wave.data[i] = LSDJ_DEFAULT_WAVE[i]
	}
}
