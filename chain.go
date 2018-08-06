package liblsdj

const LSDJ_CHAN_LENGTH = 16

type Lsdj_chan_t struct {
	phrases        [LSDJ_CHAN_LENGTH]byte
	transpositions [LSDJ_CHAN_LENGTH]byte
}

func (ch Lsdj_chan_t) Clear() {
	for i := 0; i < LSDJ_CHAN_LENGTH; i++ {
		ch.phrases[i] = 0xFF
		ch.transpositions[i] = 0
	}
}

func (dest Lsdj_chan_t) CopyFrom(source *Lsdj_chan_t) {
	/*
		TODO: add copy function
	*/
	for i := 0; i < LSDJ_CHAN_LENGTH; i++ {
		dest.phrases[i] = source.phrases[i]
		dest.transpositions[i] = source.transpositions[i]
	}
}
