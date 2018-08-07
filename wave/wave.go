package wave

// Structure represening a wave for the wave synthesizer
type Wave struct {
	data [LSDJ_WAVE_LENGTH]byte
}

// Clear all wave data to factory settings
func (wave *Wave) Clear() {
	wave.data = defaultWave
}
