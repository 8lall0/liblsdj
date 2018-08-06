package sav

import (
	"github.com/8lall0/liblsdj/project"
	"github.com/8lall0/liblsdj/song"
)

type Sav struct {
	// The projects
	projects [LSDJ_SAV_PROJECT_COUNT]*project.Project
	// Index of the project that is currently being edited
	/*! Indices start at 0, a value of 0xFF means there is no active project */
	activeProject byte
	// The song in active working memory
	song *song.Song
	//! Reserved empty memory
	reserved8120 [30]byte
}

type Header struct {
	project_names  [LSDJ_SAV_PROJECT_COUNT * 8]string
	versions       [LSDJ_SAV_PROJECT_COUNT * 1]byte
	empty          [30]byte
	init           [2]byte
	active_project byte
}

func (s *Sav) GetWorkingMemorySong() *song.Song {
	return s.song
}
func (s *Sav) SetWorkingMemorySong(song *song.Song, activeProject byte) {
	s.song = song
	s.activeProject = activeProject
}

// Per ora va bene, devi ancora fare la return
func (s *Sav) SetWorkingMemorySongFromProject(index int) {
	s.SetWorkingMemorySong(s.projects[index].GetSong())
}

func (s *Sav) GetActiveProject() byte {
	return s.activeProject
}
func (s *Sav) SetActiveProject() {}

func (s *Sav) GetProjectCount() int {
	return LSDJ_SAV_PROJECT_COUNT
}

func (s *Sav) GetProject(index int) *project.Project {
	return s.projects[index]
}

func (s *Sav) SetProject(p *project.Project, index int) {
	s.projects[index] = p
}

// Create a project that contains the working memory song
//lsdj_project_t* lsdj_project_new_from_working_memory_song(const lsdj_sav_t* sav, lsdj_error_t** error);
