package groove

type Groove struct {
	data [LSDJ_GROOVE_LENGTH]byte
}

func (t *Groove) Clear() {
	for i := 0; i < LSDJ_GROOVE_LENGTH; i++ {
		t.data[i] = defaultGroove[i]
	}
}
