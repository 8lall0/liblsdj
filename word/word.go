package word

// Structure representing word data for the speech synthesizer
type Word struct {
	allophones [LSDJ_WORD_LENGTH]byte
	lengths    [LSDJ_WORD_LENGTH]byte
}

// Clear all word data to factory settings
func (word *Word) Clear() {
	for i := 0; i < LSDJ_WORD_LENGTH; i++ {
		word.allophones[i] = 0
		word.lengths[i] = 0
	}
}
