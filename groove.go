package liblsdj

import (
	"fmt"
)

type Groove [grooveLength]byte
type Grooves [grooveCount]Groove

const (
	grooveCount   = 0x1f // The amount of grooves in a song
	grooveLength  = 16   // The number of steps in a groove
	grooveNoValue = 0    // The value of an empty (unused) step
)

func setGrooves(grooves []byte) ([]Groove, error) {
	// Adding +1, have to learn about grooves
	if len(grooves) != (grooveCount+1)*grooveLength {
		return nil, fmt.Errorf("unexpected Phrase length: %v, %v", len(grooves), (grooveCount+1)*grooveLength)
	}

	gr := make([]Groove, grooveCount+1)
	for i := 0; i < grooveCount+1; i++ {
		copy(gr[i][:], grooves[i*grooveLength:grooveLength*(i+1)])
	}

	return gr, nil
}
