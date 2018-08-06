package liblsdj

const (
	LSDJ_WORD_LENGTH int = 16
	// The constant length of all word names
	LSDJ_WORD_NAME_LENGTH int = 4
)

// Structure representing word data for the speech synthesizer
type Lsdj_word_t struct {
	allophones [LSDJ_WORD_LENGTH]byte
	lengths    [LSDJ_WORD_LENGTH]byte
}

// Clear all word data to factory settings
func (word Lsdj_word_t) Clear() {
	for i := 0; i < LSDJ_WORD_LENGTH; i++ {
		word.allophones[i] = 0
		word.lengths[i] = 0
	}
}
