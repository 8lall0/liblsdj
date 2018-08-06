package liblsdj

const LSDJ_PHRASE_LENGTH int = 16

type Lsdj_phrase_t struct {
	notes       [LSDJ_PHRASE_LENGTH]byte
	instruments [LSDJ_PHRASE_LENGTH]byte
	commands    [LSDJ_PHRASE_LENGTH]Lsdj_command_t
}

func Lsdj_phrase_copy(ph *Lsdj_phrase_t) *Lsdj_phrase_t {
	var newPh Lsdj_phrase_t

	for i := 0; i < LSDJ_PHRASE_LENGTH; i++ {
		newPh.commands[i] = ph.commands[i]
		newPh.instruments[i] = ph.instruments[i]
		newPh.notes[i] = ph.notes[i]
	}

	return &newPh
}

func Lsdj_phrase_clear(ph *Lsdj_phrase_t) {
	for i := 0; i < LSDJ_PHRASE_LENGTH; i++ {
		ph.notes[i] = 0
		ph.instruments[i] = 0xFF
		Lsdj_command_clear(&ph.commands[i])
	}
}
