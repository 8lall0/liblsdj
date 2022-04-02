package liblsdj

import (
	"errors"
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
		return errors.New(fmt.Sprintf("unexpected phrase length: %v, %v", len(phrases), phraseAllocationsLength))
	} else if len(chains) != phraseAllocationsLength { // TODO TROVA
		return errors.New(fmt.Sprintf("unexpected phrase length: %v, %v", len(chains), phraseAllocationsLength))
	} else if len(instruments) != phraseAllocationsLength { // TODO trova
		return errors.New(fmt.Sprintf("unexpected phrase length: %v, %v", len(instruments), phraseAllocationsLength))
	}

	a.Phrases = phrases
	a.Chains = chains
	a.Instruments = instruments

	return nil
}
