package liblsdj

const grooveLen = 16

type groove [grooveLen]byte

var defaultGroove = groove{0x06, 0x06, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

func (g *groove) clear() {
	*g = defaultGroove
}
