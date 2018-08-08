package liblsdj

const (
	lsdj_PROJECT_NAME_LENGTH int = 8
)

// Representation of a project within an LSDJ sav file
type Project struct {
	// The name of the project
	name []byte
	// The version of the project
	version byte
	// The song belonging to this project
	/*! If this is NULL, the project isn't in use */
	song *song
}

func (p *Project) ReadLsdsng() {
	r := new(vio)
	w := new(vio)
	/*
		TODO: improve API. Maybe make a reader and a stupid writer?
	*/

	r.open("3billetes.lsdsng")
	p.name = r.read(lsdj_PROJECT_NAME_LENGTH)
	p.version = r.readSingle()
	// decompress the songg
	decompress(r, w)
	p.song.Read(w)
}

func (p *Project) WriteLsdsng() {

}

// Change groove in a project
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
