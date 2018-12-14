package lsdj

const (
	wordLen     = 16
	wordNameLen = 4
)

type word struct {
	allophones [wordLen]byte
	lenghts    [wordLen]byte
}

func (w *word) clear() {
	for i := 0; i < wordLen; i++ {
		w.allophones[i] = 0
		w.lenghts[i] = 0
	}
}
