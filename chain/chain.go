package chain

type Chain struct {
	phrases        [LSDJ_CHAN_LENGTH]byte
	transpositions [LSDJ_CHAN_LENGTH]byte
}

func (ch *Chain) Clear() {
	for i := 0; i < LSDJ_CHAN_LENGTH; i++ {
		ch.phrases[i] = 0xFF
		ch.transpositions[i] = 0
	}
}

func Copy(source *Chain) *Chain {
	return &(*source)
}
