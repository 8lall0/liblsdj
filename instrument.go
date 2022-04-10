package liblsdj

import (
	"fmt"
)

const (
	instrumentCount      = 0x40 //! The amount of instruments in a song
	instrumentByteCount  = 16   //! The amount of bytes an instrument takes
	instrumentNameLength = 5    //! The amount of bytes an instrument name takes
)

const (
	//! The kind of instrument types that exist
	instrumentTypePulse = iota
	instrumentTypeWave
	instrumentTypeKit
	instrumentTypeNoise
)

type InstrumentInterface interface {
	setParams(b []byte)
	getParamsBytes() []byte
}

type Instrument struct {
	Name       [instrumentNameLength]byte
	Params     [instrumentByteCount]byte
	Instrument InstrumentInterface
}

func setInstruments(names, params []byte) ([]Instrument, error) {
	if len(names) != instrumentCount*instrumentNameLength {
		return nil, fmt.Errorf("unexpected instruments name length: %v, %v", len(names), instrumentCount*instrumentNameLength)
	} else if len(params) != instrumentCount*instrumentByteCount {
		return nil, fmt.Errorf("unexpected instruments name length: %v, %v", len(params), instrumentCount*instrumentByteCount)
	}

	in := make([]Instrument, instrumentCount)
	for i := 0; i < len(names)/instrumentNameLength; i++ {
		copy(in[i].Name[:], names[instrumentNameLength*i:instrumentNameLength*(i+1)])
	}
	for i := 0; i < len(params)/instrumentByteCount; i++ {
		copy(in[i].Params[:], params[instrumentByteCount*i:instrumentByteCount*(i+1)])

		var instr InstrumentInterface
		switch params[instrumentByteCount*i] {
		case instrumentTypePulse:
			instr = new(PulseInstrument)
		case instrumentTypeNoise:
			instr = new(NoiseInstrument)
		case instrumentTypeWave:
			instr = new(WaveInstrument)
		case instrumentTypeKit:
			instr = new(KitInstrument)
		}

		instr.setParams(params[instrumentByteCount*i : instrumentByteCount*(i+1)])

		in[i].Instrument = instr
	}

	return in, nil
}

const (
	instrumentTablePlay = iota
	instrumentTableStep
)

const (
	instrumentVibratoTriangle = iota
	instrumentVibratoSawtooth
	instrumentVibratoSquare
)

const (
	instrumentVibratoDown = iota
	instrumentVibratoUp
)

const (
	instrumentPlvFast = iota
	instrumentPlvTick
	instrumentPlvStep
	instrumentPlvDrum
)
