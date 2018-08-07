package groove

type Groove struct {
	data [LSDJ_GROOVE_LENGTH]byte
}

func (t *Groove) Clear() {
	t.data = defaultGroove
}
