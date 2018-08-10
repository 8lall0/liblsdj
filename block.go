package liblsdj

type block struct {
	block [BLOCK_COUNT][]byte
	cur   int
}

func (b *block) canContain(size int) bool {
	if len(b.block[b.cur])+size > BLOCK_SIZE {
		return false
	}
	return true
}

func (b *block) nextBlock() {
	b.cur++
}
