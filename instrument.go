package liblsdj

import (
	"fmt"
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
	read(in *instrument, r io.ReadSeeker)
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
	i.instrument.clear()
}

func (i *instrument) writeName(r io.ReadSeeker) {
	if _, err := io.ReadFull(r, i.name[:]); err != nil {
		panic(err)
	}
}

// TODO prima di leggere strumenti specifici, dobbiamo leggere il tipo
func (i *instrument) read(r io.ReadSeeker) {
	i.insType, _ = readByte(r)

	pos1, _ := r.Seek(0, io.SeekCurrent)

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

	pos2, _ := r.Seek(0, io.SeekCurrent)
	// TODO dovrebbe fare 15, controlla analogo di assert per golang
	fmt.Println(pos2 - pos1)

	i.envelopeVolume, _ = readByte(r)
}
