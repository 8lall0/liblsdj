package liblsdj

import (
	"errors"
	"fmt"
)

type AllocationTable struct {
	Phrases     [phraseAllocationsLength]byte
	Chains      [16]byte
	Instruments [64]byte
	Tables      [32]byte
}

func (a *AllocationTable) Set(phrases, chains, instruments, tables []byte) error {
	errs := make([]error, 0)
	if len(phrases) != phraseAllocationsLength {
		errs = append(errs, fmt.Errorf("unexpected Phrase length: %v, %v", len(phrases), phraseAllocationsLength))
	}
	// TODO i todo qui sotto sono da verificare, in realtà dovrebbe essere giusto
	if len(chains) != 16 { // TODO TROVA VALORE CORRETTO
		errs = append(errs, fmt.Errorf("unexpected Phrase length: %v, %v", len(chains), 16))
	}
	if len(instruments) != 64 { // TODO TROVA VALORE CORRETTO
		errs = append(errs, fmt.Errorf("unexpected Phrase length: %v, %v", len(instruments), 64))
	}
	if len(tables) != 32 { // TODO TROVA VALORE CORRETTO
		errs = append(errs, fmt.Errorf("unexpected Phrase length: %v, %v", len(tables), 32))
	}

	err := errors.Join(errs...)
	if err != nil {
		return err
	}

	copy(a.Phrases[:], phrases)
	copy(a.Chains[:], chains)
	copy(a.Instruments[:], instruments)
	copy(a.Tables[:], tables)

	return nil
}
