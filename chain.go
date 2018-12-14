package lsdj

const chainLen = 16

type chain struct {
	prases         [chainLen]byte
	transpositions [chainLen]byte
}

func (c *chain) clear() {
	for i := 0; i < chainLen; i++ {
		c.prases[i] = 0xFF
		c.transpositions[i] = 0
	}
}
