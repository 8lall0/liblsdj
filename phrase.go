package liblsdj

import (
	"fmt"
)

const (
	phraseCount        = 0xFF //! The amount of phrases in a song
	phraseLength       = 16   //! The number of steps in a Phrase
	phraseNoNote       = 0    //! The value of "no note" at a given step
	phraseNoInstrument = 0xFF //! The value of "no instrument" at a given step
)

const phraseAllocationsLength = 0x20     // 20
const phraseCommandsLength = 0x0ff0      // 4080
const phraseCommandValuesLength = 0x0ff0 // 4080

type Phrase struct {
	Phrase      [phraseLength]byte
	Command     [phraseLength]byte
	Value       [phraseLength]byte
	Instruments [phraseLength]byte
}

func setPhrases(phrases, commands, values, instruments []byte) ([]Phrase, error) {
	totalLength := phraseCount * phraseLength

	if len(phrases) != totalLength {
		return nil, fmt.Errorf("unexpected phrases length; expected: %v, got: %v", len(phrases), totalLength)
	} else if len(commands) != totalLength {
		return nil, fmt.Errorf("unexpected Phrase commands length; expected: %v, got: %v", len(commands), totalLength)
	} else if len(values) != totalLength {
		return nil, fmt.Errorf("unexpected Phrase values length; expected: %v, got: %v", len(values), totalLength)
	} else if len(instruments) != totalLength {
		return nil, fmt.Errorf("unexpected Phrase instruments length; expected: %v, got: %v", len(values), totalLength)
	}

	p := make([]Phrase, phraseCount)
	for i := 0; i < phraseCount; i++ {
		copy(p[i].Phrase[:], phrases[i:phraseLength*i])
		copy(p[i].Command[:], commands[i:phraseLength*i])
		copy(p[i].Value[:], values[i:phraseLength*i])
		copy(p[i].Instruments[:], instruments[i:phraseLength*i])
	}

	return p, nil
}
