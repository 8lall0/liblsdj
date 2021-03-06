package liblsdj

import (
	"io"
)

type plVibSpeed byte
type vibShape byte
type vibDirection byte

const (
	instrumentNameLen    = 5
	instrumentDefaultLen = 16

	noTable                   = 0x20
	instrumentUnlimitedLength = 0x40
	kitLengthAuto             = 0x0

	instrPulse byte = iota
	instrWave
	instrKit
	instrNoise

	plVibFast plVibSpeed = iota
	plVibTick
	plVibStep

	vibTriangle vibShape = iota
	vibSawtooth
	vibSquare

	vipUp vibDirection = iota
	vibDown
)

type instrumentT interface {
	read(in *instrument, r io.ReadSeeker, version byte)
	write(in *instrument, w io.WriteSeeker, version byte)
	clear()
}

var instrumentDefault = [instrumentDefaultLen]byte{0, 0xA8, 0, 0, 0xFF, 0, 0, 3, 0, 0, 0xD0, 0, 0, 0, 0xF3, 0}
var instrumentNameEmpty = [instrumentNameLen]byte{0, 0, 0, 0, 0}

type instrument struct {
	name           [instrumentNameLen]byte
	insType        byte
	envelopeVolume byte
	panning        panning
	table          byte // 0x20 or higher = LSDJ_NO_TABLE
	automate       byte
	instrument     instrumentT
}

func (i *instrument) clearAsPulse() {
	i.insType = instrPulse
	i.envelopeVolume = 0xA8
	i.panning = panLeftRight
	i.table = noTable
	i.automate = 0
	i.instrument = new(instrumentPulse)
	i.instrument.clear()
}

func (i *instrument) read(r io.ReadSeeker, version byte) {
	i.insType, _ = readByte(r)

	pos1, _ := r.Seek(0, io.SeekCurrent)

	i.envelopeVolume, _ = readByte(r)

	switch i.insType {
	case 0:
		i.instrument = new(instrumentPulse)
	case 1:
		i.instrument = new(instrumentWave)
	case 2:
		i.instrument = new(instrumentKit)
	case 3:
		i.instrument = new(instrumentNoise)
	default:
		panic("Strumento non conosciuto")
	}

	i.instrument.read(i, r, version)

	pos2, _ := r.Seek(0, io.SeekCurrent)

	// TODO ritorna errore per ste cagate
	if pos2-pos1 != 15 {
		panic("Non è 15")
	}

}
