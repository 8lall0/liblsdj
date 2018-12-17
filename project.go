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

// Forse è qui che inizia il writer
func readLsdsng(r io.ReadSeeker) (p Project) {
	_, _ = io.ReadFull(r, p.name[:])
	p.version, _ = readByte(r)

	// Decidi dove salvare questo writeseeker
	b := new(writerseeker.WriterSeeker)
	decompress(r, b, nil)

	p.song, _ = Read(b.BytesReader())

	return p
}
