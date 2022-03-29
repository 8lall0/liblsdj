package liblsdj

import (
	"errors"
	"fmt"
)

type Groove [grooveLength]byte
type Grooves [grooveCount]Groove

const (
	grooveCount   = 0x1f // The amount of grooves in a song
	grooveLength  = 16   // The number of steps in a groove
	grooveNoValue = 0    // The value of an empty (unused) step
)

func (g *Grooves) Set(b []byte) error {
	// Adding +1, have to learn about grooves
	if len(b) != (grooveCount+1)*grooveLength {
		return errors.New(fmt.Sprintf("unexpected phrase length: %v, %v", len(b), (grooveCount+1)*grooveLength))
	}

	for i := 0; i < grooveCount; i++ {
		copy(g[i][:], b[i:grooveLength*i])
	}

	return nil
}
