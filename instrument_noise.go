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

func (i *instrumentNoise) write(in *instrument, w io.WriteSeeker, version byte) {

	_ = writeByte(3, w)
	_ = writeByte(in.envelopeVolume, w)
	_ = writeByte(createScommandByte(i.sCommand), w)
	_ = writeByte(createLengthByte(i.length), w)
	_ = writeByte(i.shape, w)
	_ = writeByte(createAutomateByte(in.automate), w)
	_ = writeByte(createTableByte(in.table), w)
	_ = writeByte(createPanningByte(in.panning), w)

	empty := []byte{0, 0, 0xD0, 0, 0, 0, 0xF3, 0}
	_, _ = w.Write(empty) // Bytes 8-15 are empty
}
