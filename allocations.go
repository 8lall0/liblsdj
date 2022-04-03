package liblsdj

import (
	"fmt"
)

type AllocationTable struct {
	Phrases     []byte
	Chains      []byte
	Instruments []byte
	Tables      []byte
}

func (a *AllocationTable) Set(phrases, chains, instruments, tables []byte) error {
	if len(phrases) != phraseAllocationsLength {
		return fmt.Errorf("unexpected Phrase length: %v, %v", len(phrases), phraseAllocationsLength)
	} else if len(chains) != 16 { // TODO TROVA
		return fmt.Errorf("unexpected Phrase length: %v, %v", len(chains), 16)
	} else if len(instruments) != 64 { // TODO trova
		return fmt.Errorf("unexpected Phrase length: %v, %v", len(instruments), 64)
	} else if len(tables) != 32 { // TODO trova
		return fmt.Errorf("unexpected Phrase length: %v, %v", len(tables), 32)
	}

	a.Phrases = phrases
	a.Chains = chains
	a.Instruments = instruments
	a.Tables = tables

	return nil
}
