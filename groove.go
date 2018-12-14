package lsdj

const grooveLen = 16

var defaultGroove = [grooveLen]byte{0x06, 0x06, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

type groove [grooveLen]byte

func (g *groove) clear() {
	*g = defaultGroove
}
