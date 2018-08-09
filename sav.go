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

/*
func (s *Sav) SavWrite(r *vio, w *vio) {
	header := new(Header)

	header.init = make([]byte, 2)

	header.init[0] = []byte("j")[0]
	header.init[1] = []byte("j")[1]
	header.active_project = s.activeProject
	header.empty = s.reserved8120

	var blockAllocTable []byte
	var block [BLOCK_COUNT][]byte
	for i := range block {
		block[i] = make([]byte, BLOCK_SIZE)
	}

	blockAllocTable = make([]byte, BLOCK_COUNT)
	for i := 0; i < BLOCK_COUNT; i++ {
		blockAllocTable[i] = 0xFF
	}


	for i := 0; i < LSDJ_SAV_PROJECT_COUNT; i++ {
		name := make([]byte, 8)
		copy(name, s.projects[i].name)
		header.project_names = append(header.project_names, name...)

		header.versions[i] = s.projects[i].version

		song := s.projects[i].song
		if song != nil {
			 //block[currentblock - 1]
			 // TODO: compress wav

			 if (written_block_count == 0) {
			 	error
			 }

			 currentblock += written_block_count
		}
	}
	w.write(header.project_names)
	w.write(header.versions)
	w.write(header.empty)
	w.write(header.init)
	w.writeByte(header.active_project)
	w.write(blockAllocTable)
	for i:= range block {
		w.write(block[i])
	}
}*/

// Create a project that contains the working memory song
//lsdj_project_t* lsdj_project_new_from_working_memory_song(const lsdj_sav_t* sav, lsdj_error_t** error);
