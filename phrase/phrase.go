package phrase

import "github.com/8lall0/liblsdj/command"

type Phrase struct {
	notes       [LSDJ_PHRASE_LENGTH]byte
	instruments [LSDJ_PHRASE_LENGTH]byte
	commands    [LSDJ_PHRASE_LENGTH]*command.Command
}

func (ph *Phrase) Clear() {
	for i := 0; i < LSDJ_PHRASE_LENGTH; i++ {
		ph.notes[i] = 0
		ph.instruments[i] = 0xFF
		ph.commands[i].Clear()
	}
}

func Copy(src *Phrase) *Phrase {
	return &(*src)
}
