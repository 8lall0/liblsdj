package liblsdj

import (
	"errors"
	"fmt"
)

const (
	phraseCount        = 0xFF //! The amount of phrases in a song
	phraseLength       = 16   //! The number of steps in a phrase
	phraseNoNote       = 0    //! The value of "no note" at a given step
	phraseNoInstrument = 0xFF //! The value of "no instrument" at a given step
)

const phraseAllocationsLength = 0x20     // 20
const phraseCommandsLength = 0x0ff0      // 4080
const phraseCommandValuesLength = 0x0ff0 // 4080
const phraseInstrumentsLength = 0xFF0    // Inserito io

type Phrases [phraseCount][phraseLength]byte
type PhraseCommands [phraseCommandsLength]byte
type PhraseCommandValues [phraseCommandValuesLength]byte
type PhraseAllocations [phraseAllocationsLength]byte
type PhraseInstruments [phraseInstrumentsLength]byte

func (p *Phrases) Set(b []byte) error {
	if len(b) != phraseCount*phraseLength {
		return errors.New(fmt.Sprintf("unexpected phrase length: %v, %v", len(b), phraseCount*phraseLength))
	}

	for i := 0; i < phraseCount; i++ {
		copy(p[i][:], b[i:phraseLength*i])
	}

	return nil
}

func (p *PhraseAllocations) Set(b []byte) error {
	if len(b) != phraseAllocationsLength {
		return errors.New(fmt.Sprintf("unexpected length: %v, %v", len(b), phraseAllocationsLength))
	}

	copy(p[:], b[:])

	return nil
}

func (p *PhraseCommands) Set(b []byte) error {
	if len(b) != phraseCommandsLength {
		return errors.New(fmt.Sprintf("unexpected length: %v, %v", len(b), phraseCommandsLength))
	}

	copy(p[:], b[:])

	return nil
}

func (p *PhraseCommandValues) Set(b []byte) error {
	if len(b) != phraseCommandValuesLength {
		return errors.New(fmt.Sprintf("unexpected length: %v, %v", len(b), phraseCommandValuesLength))
	}

	copy(p[:], b[:])

	return nil
}

func (p *PhraseInstruments) Set(b []byte) error {
	if len(b) != phraseInstrumentsLength {
		return errors.New(fmt.Sprintf("unexpected length: %v, %v", len(b), phraseInstrumentsLength))
	}

	copy(p[:], b[:])

	return nil
}
