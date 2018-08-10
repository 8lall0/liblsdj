package liblsdj

type blockA struct {
	block   [BLOCK_COUNT][BLOCK_SIZE]byte
	cur     int
	written int
}

func (b *blockA) write(in []byte) {
	if len(b.block[b.cur])+len(in)+2 >= BLOCK_SIZE {
		b.block[b.cur][b.written+1] = SPECIAL_ACTION_BYTE
		if b.cur+1 >= BLOCK_COUNT {
			panic("MAX BLOCK")
		}
		b.cur++
		b.written = 0
	}

	for i := 0; i < len(in); i++ {
		b.block[b.cur][b.written+i] = in[i]
	}
	b.written += len(in)

}

func (b *blockA) writeByte(in byte) {
	if len(b.block[b.cur])+1+2 >= BLOCK_SIZE {
		b.block[b.cur][b.written+1] = SPECIAL_ACTION_BYTE
		if b.cur+1 >= BLOCK_COUNT {
			panic("MAX BLOCK")
		}
		b.cur++
		b.written = 0
	}

	b.block[b.cur][b.written+1] = in
	b.written++
}

func (b *blockA) readAll() []byte {
	var out []byte

	for i := 0; i < b.written; i++ {
		for j := 0; j < BLOCK_SIZE; j++ {
			out = append(out, b.block[i][j])
		}
	}
	return out
}
