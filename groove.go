package liblsdj

const lsdj_GROOVE_LENGTH int = 16

var lsdj_DEFAULT_GROOVE = []byte{0x06, 0x06, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0} //lsdj_GROOVE_LENGTH

type groove struct {
	groove []byte //lsdj_GROOVE_LENGTH
}

func (g *groove) setGroove(b []byte) {
	g.groove = b
}

func (g *groove) clear() {
	g.groove = make([]byte, lsdj_GROOVE_LENGTH)
	copy(g.groove, lsdj_DEFAULT_GROOVE)
}
