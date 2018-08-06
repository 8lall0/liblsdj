package project

import (
	"github.com/8lall0/liblsdj/song"
	"github.com/8lall0/liblsdj/vio"
)

// Representation of a project within an LSDJ sav file
type Project struct {
	// The name of the project
	name []byte
	// The version of the project
	version byte
	// The song belonging to this project
	/*! If this is NULL, the project isn't in use */
	song *song.Song
}

type Data struct {
	data []byte
	ndx  int
}

/* Require VIO */
func (p *Project) ReadLsdsng(f *Data) {
	var decompressed [song.LSDJ_SONG_DECOMPRESSED_SIZE]byte

	copy(f.data[0:LSDJ_PROJECT_NAME_LENGTH-1], p.name)
	p.version = f.data[LSDJ_PROJECT_NAME_LENGTH]

}

func (p *Project) WriteLsdsng() {

}

// Change data in a project
func (p *Project) GetName() {

}
func (p *Project) SetName() {

}

func (p *Project) GetVersion() {

}
func (p *Project) SetVersion() {

}

func (p *Project) GetSong() {

}
func (p *Project) SetSong() {

}
