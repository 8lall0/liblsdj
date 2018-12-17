package liblsdj

const phraseLen = 16

type phrase struct {
	notes       [phraseLen]byte
	instruments [phraseLen]byte
	commands    [phraseLen]command
}

func (p *phrase) clear() {
	for i := 0; i < phraseLen; i++ {
		p.notes[i] = 0
		p.instruments[i] = 0xFF
		p.commands[i].clear()
	}
}
