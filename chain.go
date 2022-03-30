package liblsdj

import (
	"errors"
	"fmt"
)

const (
	chainCount    = 0x7F //! The amount of chains in a song, UNUSED MAY BE 0X80
	chainLength   = 16   //! The number of steps in a chain
	chainNoPhrase = 0xFF //! The value of "no phrase" at a given step
)

// Format: [0..15 for 00, 0..15 for 01 etc]
type ChainPhrases []byte
type ChainTranspositions []byte
type ChainAllocations []byte

type ChainAssignments struct {
	Pulse1 []byte
	Pulse2 []byte
	Wave   []byte
	Noise  []byte
}

// La struttura dei ChainAssignments [P1 P2 W N] e si ripete per tutte le righe possibili
func (c *ChainAssignments) Set(b []byte) error {
	if len(b) != 1024 {
		return errors.New(fmt.Sprintf("unexpected length: %v, %v", len(b), 1024))
	}

	c.Pulse1 = make([]byte, 0)
	c.Pulse2 = make([]byte, 0)
	c.Wave = make([]byte, 0)
	c.Noise = make([]byte, 0)

	for i := 0; i < len(b)/4; i++ {
		c.Pulse1 = append(c.Pulse1, b[4*i])
		c.Pulse2 = append(c.Pulse1, b[4*i+1])
		c.Wave = append(c.Pulse1, b[4*i+2])
		c.Noise = append(c.Pulse1, b[4*i+3])
	}

	return nil
}
