package liblsdj

const tableLen = 16

type table struct {
	volumes        [tableLen]byte
	transpositions [tableLen]byte
	commands1      [tableLen]command
	commands2      [tableLen]command
}

func (t *table) clear() {
	for i := 0; i < tableLen; i++ {
		t.volumes[i] = 0
		t.transpositions[i] = 0
		t.commands1[i].clear()
		t.commands2[i].clear()
	}
}

func (t *table) getCommand1() (out []byte) {
	for i := 0; i < tableLen; i++ {
		out = append(out, t.commands1[i].get()...)
	}
	return out
}

func (t *table) getCommand2() (out []byte) {
	for i := 0; i < tableLen; i++ {
		out = append(out, t.commands2[i].get()...)
	}
	return out
}
