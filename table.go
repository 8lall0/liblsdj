package liblsdj

import (
	"errors"
	"fmt"
)

const (
	tableCount            = 0x20
	tableLength           = 0x10
	allocationTableLength = 0x32
	contentLength         = 512
)

type Tables [tableCount]struct {
	Command [tableLength]byte
	Value   [tableLength]byte
}

type TableAllocationTables []byte

func (t *Tables) Set(command, value []byte) error {
	if len(command) != tableCount*tableLength {
		return errors.New(fmt.Sprintf("unexpected phrase length: %v, %v", len(command), tableCount*tableLength))
	}

	if len(value) != tableCount*tableLength {
		return errors.New(fmt.Sprintf("unexpected phrase length: %v, %v", len(value), tableCount*tableLength))
	}

	// Controllalo bene!
	for i := 0; i < 4; i++ {
		copy(t[i].Command[:], command[tableLength*i:tableLength*(i+1)])
		copy(t[i].Value[:], value[tableLength*i:tableLength*(i+1)])
	}

	return nil
}
