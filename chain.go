package liblsdj

const chainLen = 16

type chain struct {
	phrases        [chainLen]byte
	transpositions [chainLen]byte
}

func (c *chain) clear() {
	for i := 0; i < chainLen; i++ {
		c.phrases[i] = 0xFF
		c.transpositions[i] = 0
	}
}
