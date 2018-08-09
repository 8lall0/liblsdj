package liblsdj

const (
	lsdj_WORD_LENGTH int = 16
	// The constant length of all word names
	lsdj_WORD_NAME_LENGTH int = 4
)

type wordA []*word

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

func (w *wordA) initialize() {
	*w = make([]*word, lsdj_WORD_COUNT)
	for i := 0; i < lsdj_WORD_COUNT; i++ {
		(*w)[i] = new(word)
	}
}

func (w wordA) write(r *vio) {
	for i := 0; i < lsdj_WORD_COUNT; i++ {
		w[i].allophones = r.read(lsdj_WORD_LENGTH)
		w[i].lengths = r.read(lsdj_WORD_LENGTH)
	}
}
