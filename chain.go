package liblsdj

import (
	"errors"
	"fmt"
)

const (
	chainCount    = 0x7F + 1 //! The amount of chains in a song, UNUSED MAY BE 0X80
	chainLength   = 16       //! The number of steps in a chain
	chainNoPhrase = 0xFF     //! The value of "no Phrase" at a given step

	chainAssignLength = 0x100
)

type Chain struct {
	phrase        [chainLength]byte
	transposition [chainLength]byte
}

func setChains(phrases, transpositions []byte) ([]Chain, error) {
	if len(phrases) != chainCount*chainLength {
		return nil, fmt.Errorf("unexpected chain phrases length; expected: %v, got: %v", len(phrases), chainCount*chainLength)
	} else if len(transpositions) != chainCount*chainLength {
		return nil, fmt.Errorf("unexpected chain transpositions length; expected: %v, got: %v", len(transpositions), chainCount*chainLength)
	}

	// ChainPhrases Format: [0..15 for 00, 0..15 for 01 etc]
	c := make([]Chain, chainCount)
	for i := 0; i < chainCount; i++ {
		copy(c[i].phrase[:], phrases[i*chainLength:chainLength*(i+1)])
		copy(c[i].transposition[:], transpositions[i*chainLength:chainLength*(i+1)])
	}

	return c, nil
}

type ChainAssignments struct {
	Pulse1 []byte
	Pulse2 []byte
	Wave   []byte
	Noise  []byte
}

// Set La struttura dei ChainAssignments [P1 P2 W N] e si ripete per tutte le righe possibili
func (c *ChainAssignments) Set(b []byte) error {
	if len(b) != 4*chainAssignLength {
		return errors.New(fmt.Sprintf("unexpected length: %v, %v", len(b), 4*chainAssignLength))
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

func (c *ChainAssignments) Get() []byte {
	b := make([]byte, 0)

	for i := 0; i < chainAssignLength; i++ {
		b = append(b, c.Pulse1[i])
		b = append(b, c.Pulse2[i])
		b = append(b, c.Wave[i])
		b = append(b, c.Noise[i])
	}

	return b
}
