package liblsdj

import "fmt"

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
	readData := new(vio)
	decomprSong := new(vio)
	decomprSongDecoded := new(vio)
	/*savFile := new(vio)

	s := new (Sav)*/
	/*
		TODO: improve API. Maybe make a reader and a stupid writer?
	*/

	readData.open("3billetes.lsdsng")
	p.name = readData.read(lsdj_PROJECT_NAME_LENGTH)
	p.version = readData.readByte()
	p.song = new(song)
	// decompress the songg
	decompress(readData, decomprSong)

	p.song.Read(decomprSong)
	return
	p.song.Write(decomprSongDecoded)
	fmt.Println(decomprSongDecoded.getLen())
	/*s.SetProject(p, 0)
	s.SavWrite(decomprSong, savFile)
	ioutil.WriteFile("gino.sav", savFile.get(), 0755)*/
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
