package wave

// Structure represening a wave for the wave synthesizer
type Wave struct {
	data []byte
}

// Clear all wave data to factory settings
func (wave *Wave) Clear() {
	wave.data = LSDJ_DEFAULT_WAVE
}
