package liblsdj

import (
	"fmt"
)

const (
	tableCount                = 0x20
	tableLength               = 0x10
	tableAllocationLength     = 0x20
	tableTranspositionsLength = 0x200
	tableEnvelopesLength      = 0x200
)

type Table struct {
	envelopes     [tableLength]byte
	transposition [tableLength]byte
	col1, col2    struct {
		command [tableLength]byte
		value   [tableLength]byte
	}
}

func setTables(envelopes, transpositions []byte, col1Commands, col1Values, col2Commands, col2Values []byte) ([]Table, error) {
	totalLength := tableCount * tableLength

	if len(envelopes) != totalLength {
		return nil, fmt.Errorf("unexpected phrases length; expected: %v, got: %v", len(envelopes), totalLength)
	} else if len(transpositions) != totalLength {
		return nil, fmt.Errorf("unexpected phrase commands length; expected: %v, got: %v", len(transpositions), totalLength)
	} else if len(col1Commands) != totalLength {
		return nil, fmt.Errorf("unexpected phrase values length; expected: %v, got: %v", len(col1Commands), totalLength)
	} else if len(col1Values) != totalLength {
		return nil, fmt.Errorf("unexpected phrase instruments length; expected: %v, got: %v", len(col1Values), totalLength)
	} else if len(col2Commands) != totalLength {
		return nil, fmt.Errorf("unexpected phrase values length; expected: %v, got: %v", len(col2Commands), totalLength)
	} else if len(col2Values) != totalLength {
		return nil, fmt.Errorf("unexpected phrase instruments length; expected: %v, got: %v", len(col2Values), totalLength)
	}

	p := make([]Table, tableCount)
	for i := 0; i < tableCount; i++ {
		copy(p[i].envelopes[:], envelopes[i:tableLength*i])
		copy(p[i].transposition[:], transpositions[i:tableLength*i])
		copy(p[i].col1.command[:], col1Commands[i:tableLength*i])
		copy(p[i].col1.value[:], col1Values[i:tableLength*i])
		copy(p[i].col2.command[:], col2Commands[i:tableLength*i])
		copy(p[i].col2.value[:], col2Values[i:tableLength*i])
	}

	return p, nil
}
