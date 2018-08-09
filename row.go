package liblsdj

type rowA []*row

type row struct {
	pulse1 byte
	pulse2 byte
	wave   byte
	noise  byte
}

func (r row) clear() {
	r.pulse1 = 0xFF
	r.pulse2 = 0xFF
	r.wave = 0xFF
	r.noise = 0xFF
}

func (r *rowA) initialize() {
	*r = make([]*row, lsdj_ROW_COUNT)
	for i := 0; i < lsdj_ROW_COUNT; i++ {
		(*r)[i] = new(row)
		//r[i].clear()
	}
}

func (row rowA) write(r *vio) {
	for i := 0; i < lsdj_ROW_COUNT; i++ {
		row[i].pulse1 = r.readByte()
		row[i].pulse2 = r.readByte()
		row[i].wave = r.readByte()
		row[i].noise = r.readByte()
	}
}
