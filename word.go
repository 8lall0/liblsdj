package liblsdj

const (
	lsdj_WORD_LENGTH int = 16
	// The constant length of all word names
	lsdj_WORD_NAME_LENGTH int = 4
)

// Structure representing word groove for the speech synthesizer
type word struct {
	allophones []byte //lsdj_WORD_LENGTH
	lengths    []byte //lsdj_WORD_LENGTH
}

// clear all word groove to factory settings
func (w *word) clear() {
	w.allophones = make([]byte, lsdj_WORD_LENGTH)
	w.lengths = make([]byte, lsdj_WORD_LENGTH)
}
