package liblsdj

import "io"

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

func (wo *word) write(w io.WriteSeeker) {
	_, _ = w.Write(wo.allophones[:])
	_, _ = w.Write(wo.lenghts[:])
}
