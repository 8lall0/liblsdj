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

type Phrases [phraseCount][phraseLength]byte

func (p *Phrases) Set(b []byte) error {
	if len(b) != phraseCount*phraseLength {
		return errors.New(fmt.Sprintf("unexpected phrase length: %v, %v", len(b), phraseCount*phraseLength))
	}

	for i := 0; i < phraseCount; i++ {
		copy(p[i][:], b[i:phraseLength*i])
	}

	return nil
}
