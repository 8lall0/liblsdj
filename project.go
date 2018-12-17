package liblsdj

import "io"

const projectNameLen = 8

type Project struct {
	name    [projectNameLen]byte
	version byte
	song    *Song
}

func readLsdsng(r io.ReadSeeker) Project {
	var p Project
	_, _ = io.ReadFull(r, p.name[:])
	p.version, _ = readByte(r)

	//TODO decompression
	decompress(r, w)

	p.song = Read()

	return p
}
