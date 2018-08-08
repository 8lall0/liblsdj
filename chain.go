package liblsdj

const lsdj_CHAIN_LENGTH int = 16

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
