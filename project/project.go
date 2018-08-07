package project

import (
	"fmt"
	"github.com/8lall0/liblsdj/compression"
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

func (p *Project) ReadLsdsng() {
	r := new(vio.Vio)
	w := new(vio.Vio)
	/*
		TODO: improve API. Maybe make a reader and a stupid writer?
	*/

	r.Open("3billetes.lsdsng")
	p.name = r.Read(LSDJ_PROJECT_NAME_LENGTH)
	p.version = r.ReadSingle()
	compression.Decompress(r, w)
	//w.Finalize(song.LSDJ_SONG_DECOMPRESSED_SIZE)
	//p.song.Read(w.Get())

	return

	//

	//fmt.Println(writeVio);
	fmt.Println(p.name)
	return
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
