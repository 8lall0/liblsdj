package liblsdj

const (
	noActiveProject = 0xff
	savProjectCnt   = 32
	headerStart     = songDecompressedSize
	blockCnt        = 191
	blockSize       = 0x200
)

type sav struct {
	projects     [savProjectCnt]*Project
	active       byte
	activeSong   *Song
	reserved8120 [30]byte
}

type headerT struct {
	projectNames  [savProjectCnt * 8]byte
	versions      [savProjectCnt * 1]byte
	empty         [30]byte
	init          [2]byte
	activeProject byte
}

func (s *sav) setWorkingSong(so *Song, activeProject byte) {
	s.activeSong = so
	s.active = activeProject
}

func SavWrite(s *sav) {
	var header headerT

	header.init[0] = 'j'
	header.init[1] = 'k'
	header.activeProject = s.active

	//TODO empty and reservec

	// Create the block allocation table for writing
	var blockAllocTable [blockCnt]byte
	//var blocks [blockCnt][blockSize]byte

	for i := 0; i < blockCnt; i++ {
		blockAllocTable[i] = 0xff
	}

	for i := 0; i < savProjectCnt; i++ {

	}

}
