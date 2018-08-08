package liblsdj

const lsdj_PHRASE_LENGTH int = 16

type phrase struct {
	notes       []byte     //lsdj_PHRASE_LENGTH
	instruments []byte     //lsdj_PHRASE_LENGTH
	commands    []*command //lsdj_PHRASE_LENGTH
}

func (p *phrase) clear() {
	p.notes = make([]byte, lsdj_PHRASE_LENGTH)
	p.instruments = make([]byte, lsdj_PHRASE_LENGTH)
	p.commands = make([]*command, lsdj_PHRASE_LENGTH)

	for i := range p.instruments {
		p.instruments[i] = 0xFF
	}
	for i := range p.commands {
		p.commands[i].clear()
	}
}

func (p *phrase) copy() *phrase {
	return &(*p)
}
