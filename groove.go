package liblsdj

const lsdj_GROOVE_LENGTH int = 16

var lsdj_DEFAULT_GROOVE = []byte{0x06, 0x06, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

type grooveA []*groove

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

func (g *grooveA) initialize() {
	*g = make([]*groove, lsdj_GROOVE_COUNT)
	for i := 0; i < lsdj_GROOVE_COUNT; i++ {
		(*g)[i] = new(groove)
	}
}

func (g grooveA) write(r *vio) {
	for i := 0; i < lsdj_GROOVE_COUNT; i++ {
		g[i].groove = r.read(lsdj_GROOVE_LENGTH)
	}
}
