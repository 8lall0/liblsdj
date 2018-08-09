package liblsdj

const lsdj_PHRASE_LENGTH int = 16

var lsdj_PHRASE_LENGTH_ZERO = []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
var lsdj_PHRASE_LENGTH_FF = []byte{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF}

type phraseA []*phrase

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

func (p *phraseA) initialize(allocTable []byte) {
	var i uint8

	*p = make([]*phrase, lsdj_PHRASE_COUNT)
	for i = 0; i < uint8(lsdj_PHRASE_COUNT); i++ {
		if (allocTable[i/8]>>(i%8))&1 == 1 {
			(*p)[i] = new(phrase)
		} else {
			(*p)[i] = nil
		}
	}
}

func (p phraseA) readNote(r *vio) {
	for i := 0; i < lsdj_PHRASE_COUNT; i++ {
		if p[i] != nil {
			p[i].notes = r.read(lsdj_PHRASE_LENGTH)
		} else {
			r.seekCur(lsdj_PHRASE_LENGTH)
		}
	}
}

func (p phraseA) writeNote(w *vio) {
	for i := 0; i < lsdj_PHRASE_COUNT; i++ {
		if p[i] != nil {
			w.write(p[i].notes)
		} else {
			w.write(lsdj_PHRASE_LENGTH_ZERO)
		}
	}
}

func (p phraseA) readCommand(r *vio) {
	for i := 0; i < lsdj_PHRASE_COUNT; i++ {
		if p[i] != nil {
			p[i].commands = make([]*command, lsdj_PHRASE_LENGTH)
			for j := 0; j < lsdj_PHRASE_LENGTH; j++ {
				p[i].commands[j] = new(command)
				p[i].commands[j].command = r.readByte()
			}
		} else {
			r.seekCur(lsdj_PHRASE_LENGTH)
		}
	}
	for i := 0; i < lsdj_PHRASE_COUNT; i++ {
		if p[i] != nil {
			for j := 0; j < lsdj_PHRASE_LENGTH; j++ {
				p[i].commands[j].value = r.readByte()
			}
		} else {
			r.seekCur(lsdj_PHRASE_LENGTH)
		}
	}
}

func (p phraseA) readInstrument(r *vio) {
	for i := 0; i < lsdj_PHRASE_COUNT; i++ {
		if p[i] != nil {
			p[i].instruments = r.read(lsdj_PHRASE_LENGTH)
		} else {
			r.seekCur(lsdj_PHRASE_LENGTH)
		}
	}
}

func (p phraseA) writePhraseAllocTable(w *vio) {
	table := make([]byte, lsdj_PHRASE_ALLOC_TABLE_SIZE)
	for i := 0; i < lsdj_PHRASE_COUNT; i++ {
		if p[i] != nil {
			table[i] = 1
		} else {
			table[i] = 0
		}
	}
	w.write(table)
}
