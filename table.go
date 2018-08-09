package liblsdj

const lsdj_TABLE_LENGTH int = 16

type tableA []*table

type table struct {
	// The volume column of the table
	volumes []byte //lsdj_TABLE_LENGTH
	// The transposition column of the table
	transpositions []byte //lsdj_TABLE_LENGTH
	// The first effect command column of the table
	commands1 []*command //lsdj_TABLE_LENGTH
	// The second effect command column of the table
	commands2 []*command //lsdj_TABLE_LENGTH
}

func (t *table) clear() {
	t.volumes = make([]byte, lsdj_TABLE_LENGTH)
	t.transpositions = make([]byte, lsdj_TABLE_LENGTH)
	t.commands1 = make([]*command, lsdj_TABLE_LENGTH)
	t.commands2 = make([]*command, lsdj_TABLE_LENGTH)

	for i := 0; i < lsdj_TABLE_LENGTH; i++ {
		t.commands1[i].clear()
		t.commands2[i].clear()
	}
}

func (t *table) copy() *table {
	return &(*t)
}

func (t *table) SetVolume(volume byte, index int) {
	t.volumes[index] = volume
}

func (t *table) SetVolumes(volumes []byte) {
	t.volumes = volumes
}

func (t *table) GetVolume(index int) byte {
	return t.volumes[index]
}

func (t *table) SetTransposition(transposition byte, index int) {
	t.transpositions[index] = transposition
}

func (t *table) SetTranspositions(transpositions []byte) {
	t.transpositions = transpositions
}

func (t *table) GetTranspositions(index int) byte {
	return t.transpositions[index]
}

func (t *table) GetCommand1(index int) *command {
	return t.commands1[index]
}

func (t *table) GetCommand2(index int) *command {
	return t.commands2[index]
}

func (t *tableA) initialize(allocTable []byte) {
	*t = make([]*table, lsdj_TABLE_COUNT)
	for i := 0; i < lsdj_TABLE_COUNT; i++ {
		if allocTable[i] != 0 {
			(*t)[i] = new(table)
		} else {
			(*t)[i] = nil
		}
	}
}

func (t tableA) write(r *vio) {
	for i := 0; i < lsdj_TABLE_COUNT; i++ {
		if t[i] != nil {
			t[i].volumes = r.read(lsdj_TABLE_LENGTH)
		} else {
			r.seekCur(lsdj_TABLE_LENGTH)
		}
	}
}
