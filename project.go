package liblsdj

import (
	"github.com/orcaman/writerseeker"
	"io"
)

const projectNameLen = 8

type Project struct {
	name    [projectNameLen]byte
	version byte
	song    Song
}

func ReadLsdsng(r io.ReadSeeker) (p Project) {
	_, _ = io.ReadFull(r, p.name[:])
	p.version, _ = readByte(r)

	b := new(writerseeker.WriterSeeker)
	decompress(r, b, nil)

	p.song, _ = ReadSong(b.BytesReader(), p.version)

	return p
}
