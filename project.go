package liblsdj

import (
	"fmt"
	"github.com/orcaman/writerseeker"
	"io"
)

const projectNameLen = 8

type Project struct {
	name    [projectNameLen]byte
	version byte
	song    *Song
}

func ReadLsdsng(r io.ReadSeeker) (p Project) {
	_, _ = io.ReadFull(r, p.name[:])
	p.version, _ = readByte(r)

	b := new(writerseeker.WriterSeeker)
	decompress(r, b, nil)

	p.song, _ = ReadSong(b.BytesReader(), p.version)

	return p
}

func WriteLsdsng(w io.WriteSeeker, p Project) {
	// controlla se c'è la canzona
	// if song == null
	// if name
	// if version

	_, _ = w.Write(p.name[:])
	_ = writeByte(p.version, w)

	tmp := make([]byte, songDecompressedSize)
	for i := 0; i < songDecompressedSize; i++ {
		tmp[i] = 0x34
	}
	b := new(writerseeker.WriterSeeker)
	_, _ = b.Write(tmp)
	_, _ = b.Seek(0, io.SeekStart)
	WriteSong(b, p.song)
	asd := compress(b.BytesReader(), w, 1)

	fmt.Println(asd)
}
