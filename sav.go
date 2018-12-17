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
