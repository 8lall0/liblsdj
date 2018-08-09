package liblsdj

const lsdj_CHAIN_LENGTH int = 16

type chainA []*chain

type chain struct {
	phrases        []byte //lsdj_CHAIN_LENGTH
	transpositions []byte //lsdj_CHAIN_LENGTH
}

func (c *chain) clear() {
	c.transpositions = make([]byte, lsdj_CHAIN_LENGTH)
	c.phrases = make([]byte, lsdj_CHAIN_LENGTH)
	for i := range c.phrases {
		c.phrases[i] = 0xFF
	}
}

func (c *chain) copy() *chain {
	return &(*c)
}

func (c *chainA) initialize(allocTable []byte) {
	var i uint8

	*c = make([]*chain, lsdj_CHAIN_COUNT)
	for i = 0; i < uint8(lsdj_CHAIN_COUNT); i++ {
		if (allocTable[i/8]>>(i%8))&1 == 1 {
			(*c)[i] = new(chain)
		} else {
			(*c)[i] = nil
		}
	}
}

func (c chainA) write(r *vio) {
	for i := 0; i < lsdj_CHAIN_COUNT; i++ {
		if c[i] != nil {
			c[i].transpositions = r.read(lsdj_CHAIN_LENGTH)
		} else {
			r.seekCur(lsdj_CHAIN_LENGTH)
		}
	}
}
