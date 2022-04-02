package liblsdj

import (
	"errors"
	"fmt"
)

const (
	instrumentCount               = 0x40 //! The amount of instruments in a song
	instrumentByteCount           = 16   //! The amount of bytes an instrument takes
	instrumentNameLength          = 5    //! The amount of bytes an instrument name takes
	instrumentPulseLengthInfinite = 0x40 //! The value of an infinite pulse length
	instrumentKitLengthAuto       = 0x0  //! The value of a kit length set to AUTO
	instrumentNoiseLengthInfinite = 0x40 //! The value of an infinite noise length
)

type InstrumentAllocationTable [instrumentCount]byte
type InstrumentParams [instrumentCount * instrumentByteCount]byte
type InstrumentNames [instrumentCount][instrumentNameLength]byte

func (in *InstrumentAllocationTable) Set(b []byte) error {
	if len(b) != instrumentCount {
		return errors.New(fmt.Sprintf("unexpected length: %v, %v", len(b), instrumentCount))
	}

	copy(in[:], b[:])

	return nil
}

func (in *InstrumentNames) Set(b []byte) error {
	if len(b) != instrumentCount*instrumentNameLength {
		return errors.New(fmt.Sprintf("unexpected length: %v, %v", len(b), instrumentCount*instrumentNameLength))
	}

	for i := 0; i < len(b)/instrumentNameLength; i++ {
		copy(in[i][:], b[instrumentNameLength*i:instrumentNameLength*(i+1)])
	}

	return nil
}

func (in *InstrumentParams) Set(b []byte) error {
	if len(b) != instrumentCount*instrumentByteCount {
		return errors.New(fmt.Sprintf("unexpected length: %v, %v", len(b), instrumentCount*instrumentByteCount))
	}

	copy(in[:], b[:])

	return nil
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
