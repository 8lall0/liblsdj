package liblsdj

const (
	LSDJ_NO_ACTIVE_PROJECT byte = 0xFF
	LSDJ_SAV_PROJECT_COUNT      = 32
	HEADER_START                = lsdj_SONG_DECOMPRESSED_SIZE
	BLOCK_COUNT            int  = 191
	BLOCK_SIZE             int  = int(0x200)
)

type Sav struct {
	// The projects
	projects [LSDJ_SAV_PROJECT_COUNT]*Project
	// Index of the project that is currently being edited
	/*! Indices start at 0, a value of 0xFF means there is no active project */
	activeProject byte
	// The song in active working memory
	song *song
	//! Reserved empty memory
	reserved8120 []byte

	header Header
}

type Header struct {
	project_names  []byte //LSDJ_SAV_PROJECT_COUNT * 8
	versions       []byte //LSDJ_SAV_PROJECT_COUNT * 1
	empty          []byte
	init           []byte
	active_project byte
}

func (s *Sav) GetWorkingMemorySong() *song {
	return s.song
}
func (s *Sav) SetWorkingMemorySong(song *song, activeProject byte) {
	s.song = song
	s.activeProject = activeProject
}

// Per ora va bene, devi ancora fare la return
func (s *Sav) SetWorkingMemorySongFromProject(index int) {
	//s.SetWorkingMemorySong(s.projects[index].GetSong())
}

func (s *Sav) GetActiveProject() byte {
	return s.activeProject
}
func (s *Sav) SetActiveProject() {}

func (s *Sav) GetProjectCount() int {
	return LSDJ_SAV_PROJECT_COUNT
}

func (s *Sav) GetProject(index int) *Project {
	return s.projects[index]
}

func (s *Sav) SetProject(p *Project, index int) {
	s.projects[index] = p
}

func (s *Sav) SavWrite(r *vio, w *vio) {
	var blockAllocCur int
	var blockAllocTable []byte
	var blocks blockA

	blockAllocTable = make([]byte, BLOCK_COUNT)
	for i := 0; i < BLOCK_COUNT; i++ {
		blockAllocTable[i] = 0xFF
	}

	s.activeProject = 0xFF
	s.header.init = []byte("jk")
	s.header.active_project = s.activeProject
	s.header.empty = s.reserved8120

	for true {
		i := 0
		// Lunghezza fissa di almeno 8
		for j := 0; j < len(s.projects[i].name); j++ {
			s.header.project_names = append(s.header.project_names, s.projects[i].name[j])
		}
		for j := len(s.projects[i].name); j < 8; j++ {
			s.header.versions = append(s.header.versions, 0)
		}

		// TODO: error handling
		writtenBlocks := compress(r, &blocks)
		if writtenBlocks == 0 {
			panic("Non abbastanza spazio")
		}
		for j := 0; j < writtenBlocks; j++ {
			blockAllocTable[blockAllocCur] = byte(i)
			blockAllocCur++
		}
		break
	}
	w.write(s.header.project_names)
	w.write(s.header.versions)
	w.write(s.header.empty)
	w.write(s.header.init)
	w.writeByte(s.header.active_project)
	w.write(blockAllocTable)
	w.write(blocks.readAll())
}
