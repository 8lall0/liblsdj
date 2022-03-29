package liblsdj

import (
	"errors"
	"fmt"
)

const (
	tableCount  = 0x20
	tableLength = 0x10
)

type Tables [tableCount]struct {
	Command [tableLength]byte
	Value   [tableLength]byte
}

func (t *Tables) SetCommand(b []byte) error {
	if len(b) != tableCount*tableLength {
		return errors.New(fmt.Sprintf("unexpected phrase length: %v, %v", len(b), tableCount*tableLength))
	}

	for i := 0; i < 4; i++ {
		copy(t[i].Command[:], b[i:tableLength*i])
	}

	return nil
}

func (t *Tables) SetValue(b []byte) error {
	if len(b) != tableCount*tableLength {
		return errors.New(fmt.Sprintf("unexpected phrase length: %v, %v", len(b), tableCount*tableLength))
	}

	for i := 0; i < 4; i++ {
		copy(t[i].Value[:], b[i:tableLength*i])
	}

	return nil
}
