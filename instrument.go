package liblsdj

import "io"

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
	read(r io.ReadSeeker)
	clearInstrument()
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
	i.instrument.clearInstrument()
}

func (i *instrument) writeName(r io.ReadSeeker) {
	if _, err := io.ReadFull(r, i.name[:]); err != nil {
		panic(err)
	}
}
