package liblsdj

import (
	"fmt"
)

const (
	instrumentCount               = 0x40 //! The amount of instruments in a song
	instrumentByteCount           = 16   //! The amount of bytes an instrument takes
	instrumentNameLength          = 5    //! The amount of bytes an instrument name takes
	instrumentPulseLengthInfinite = 0x40 //! The value of an infinite pulse length
	instrumentKitLengthAuto       = 0x0  //! The value of a InstrumentKit length set to AUTO
	instrumentNoiseLengthInfinite = 0x40 //! The value of an infinite noise length
)

type InstrumentParams [instrumentCount * instrumentByteCount]byte
type InstrumentNames [instrumentCount][instrumentNameLength]byte

type Instrument struct {
	Name   [instrumentNameLength]byte
	Params [instrumentByteCount]byte
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
	}

	return in, nil
}

//! The kind of instrument types that exist
const (
	instrumentTypePulse = iota
	instrumentTypeWave
	instrumentTypeKit
	instrumentTypeNoise
)

const (
	instrumentTablePlay = iota
	instrumentTableStep
)

const (
	instrumentWaveVolume0 = 0x00
	instrumentWaveVolume1 = 0x60
	instrumentWaveVolume2 = 0x40
	instrumentWaveVolume3 = 0xA8
)

const (
	instrumentPulseWidth125 = iota
	instrumentPulseWidth25
	instrumentPulseWidth50
	instrumentPulseWidth75
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

const (
	instrumentWavePlayOnce = iota
	instrumentWavePlayLoop
	instrumentWavePlayPingPong
	instrumentWavePlayManual
)

const (
	instrumentKitLoopOff = iota
	instrumentKitLoopOn
	instrumentKitLoopAttack
)

const (
	instrumentKitDistortionClip = iota
	instrumentKitDistortionShape
	instrumentKitDistortionShape2
	instrumentKitDistortionWrap
)

const (
	instrumentNoiseFree = iota
	instrumentNoiseStable
)
