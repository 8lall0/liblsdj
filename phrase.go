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

type Phrases [phraseCount][phraseLength]byte
type Phrase [phraseCount]struct {
	Command [phraseLength]byte
	Value   [phraseLength]byte
}
type PhraseInstruments [phraseCount * phraseLength]byte
type PhraseAllocations [phraseAllocationsLength]byte

func (p *Phrases) Set(b []byte) error {
	if len(b) != phraseCount*phraseLength {
		return errors.New(fmt.Sprintf("unexpected phrase length: %v, %v", len(b), phraseCount*phraseLength))
	}

	for i := 0; i < phraseCount; i++ {
		copy(p[i][:], b[i:phraseLength*i])
	}

	return nil
}

func (p *PhraseInstruments) Set(b []byte) error {
	if len(b) != phraseCount*phraseLength {
		return errors.New(fmt.Sprintf("unexpected phrase length: %v, %v", len(b), phraseCount*phraseLength))
	}

	copy(p[:], b[:])

	return nil
}

func (p *Phrase) SetCommand(b []byte) error {
	if len(b) != phraseCount*phraseLength {
		return errors.New(fmt.Sprintf("unexpected phrase length: %v, %v", len(b), phraseCount*phraseLength))
	}

	for i := 0; i < 4; i++ {
		copy(p[i].Command[:], b[i:phraseLength*i])
	}

	return nil
}

func (p *Phrase) SetValue(b []byte) error {
	if len(b) != phraseCount*phraseLength {
		return errors.New(fmt.Sprintf("unexpected phrase length: %v, %v", len(b), phraseCount*phraseLength))
	}

	for i := 0; i < 4; i++ {
		copy(p[i].Value[:], b[i:phraseLength*i])
	}

	return nil
}

func (pa *PhraseAllocations) Set(b []byte) error {
	if len(b) != phraseAllocationsLength {
		return errors.New(fmt.Sprintf("unexpected phrase length: %v, %v", len(b), phraseAllocationsLength))
	}

	copy(pa[:], b[:])

	return nil
}
