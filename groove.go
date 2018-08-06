package liblsdj

const LSDJ_GROOVE_LENGTH int = 16

var DEFAULT_GROOVE = [LSDJ_GROOVE_LENGTH]byte{0x06, 0x06, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

type Lsdj_groove_t struct {
	data [LSDJ_GROOVE_LENGTH]byte
}

func Lsdj_groove_clear(t *Lsdj_groove_t) {
	for i := 0; i < LSDJ_GROOVE_LENGTH; i++ {
		t.data[i] = DEFAULT_GROOVE[i]
	}
}
