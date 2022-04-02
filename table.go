package liblsdj

import (
	"errors"
	"fmt"
)

const (
	tableCount                = 0x20
	tableLength               = 0x10
	tableAllocationLength     = 0x20
	tableTranspositionsLength = 0x200
	tableEnvelopesLength      = 0x200
)

type Tables [tableCount]struct {
	Command [tableLength]byte
	Value   [tableLength]byte
}

type TableAllocationTable [tableAllocationLength]byte
type TableTranspositions [tableTranspositionsLength]byte
type TableEnvelopes [tableEnvelopesLength]byte

func (t *Tables) Set(command, value []byte) error {
	if len(command) != tableCount*tableLength {
		return errors.New(fmt.Sprintf("unexpected phrase length: %v, %v", len(command), tableCount*tableLength))
	}

	if len(value) != tableCount*tableLength {
		return errors.New(fmt.Sprintf("unexpected phrase length: %v, %v", len(value), tableCount*tableLength))
	}

	for i := 0; i < 4; i++ {
		copy(t[i].Command[:], command[tableLength*i:tableLength*(i+1)])
		copy(t[i].Value[:], value[tableLength*i:tableLength*(i+1)])
	}

	return nil
}

func (ta *TableAllocationTable) Set(b []byte) error {
	if len(b) != tableAllocationLength {
		return errors.New(fmt.Sprintf("unexpected length: %v, %v", len(b), tableAllocationLength))
	}

	copy(ta[:], b[:])

	return nil
}

func (tt *TableTranspositions) Set(b []byte) error {
	if len(b) != tableTranspositionsLength {
		return errors.New(fmt.Sprintf("unexpected length: %v, %v", len(b), tableTranspositionsLength))
	}

	copy(tt[:], b[:])

	return nil
}

func (te *TableEnvelopes) Set(b []byte) error {
	if len(b) != tableEnvelopesLength {
		return errors.New(fmt.Sprintf("unexpected length: %v, %v", len(b), tableEnvelopesLength))
	}

	copy(te[:], b[:])

	return nil
}
