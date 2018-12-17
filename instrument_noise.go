package liblsdj

import "io"

type sCommand byte

const (
	sCommandFree sCommand = iota
	sCommandStable
)

type instrumentNoise struct {
	length   byte
	shape    byte
	sCommand sCommand
}

func (i *instrumentNoise) clear() {
	i.length = instrumentUnlimitedLength
	i.shape = 0xFF
	i.sCommand = sCommandFree
}

func (i *instrumentNoise) read(in *instrument, r io.ReadSeeker) {
	var b byte

	b, _ = readByte(r)
	i.sCommand = parseScommand(b)

	b, _ = readByte(r)
	i.length = parseLength(b)

	i.shape, _ = readByte(r)

	b, _ = readByte(r)
	in.automate = parseAutomate(b)

	b, _ = readByte(r)
	in.table = parseTable(b)

	b, _ = readByte(r)
	in.panning = parsePanning(b)

	_, _ = r.Seek(8, io.SeekCurrent) // Bytes 8-15 are empty
}
