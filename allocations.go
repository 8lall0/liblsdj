package liblsdj

import (
	"fmt"
)

type AllocationTable struct {
	Phrases     [phraseAllocationsLength]byte
	Chains      [16]byte
	Instruments [64]byte
	Tables      [32]byte
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

	copy(a.Phrases[:], phrases)
	copy(a.Chains[:], chains)
	copy(a.Instruments[:], instruments)
	copy(a.Tables[:], tables)

	return nil
}
