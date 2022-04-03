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
	Envelopes     [tableLength]byte
	Transposition [tableLength]byte
	Col1, Col2    struct {
		Command [tableLength]byte
		Value   [tableLength]byte
	}
}

func setTables(envelopes, transpositions []byte, col1Commands, col1Values, col2Commands, col2Values []byte) ([]Table, error) {
	totalLength := tableCount * tableLength

	if len(envelopes) != totalLength {
		return nil, fmt.Errorf("unexpected phrases length; expected: %v, got: %v", len(envelopes), totalLength)
	} else if len(transpositions) != totalLength {
		return nil, fmt.Errorf("unexpected Phrase commands length; expected: %v, got: %v", len(transpositions), totalLength)
	} else if len(col1Commands) != totalLength {
		return nil, fmt.Errorf("unexpected Phrase values length; expected: %v, got: %v", len(col1Commands), totalLength)
	} else if len(col1Values) != totalLength {
		return nil, fmt.Errorf("unexpected Phrase instruments length; expected: %v, got: %v", len(col1Values), totalLength)
	} else if len(col2Commands) != totalLength {
		return nil, fmt.Errorf("unexpected Phrase values length; expected: %v, got: %v", len(col2Commands), totalLength)
	} else if len(col2Values) != totalLength {
		return nil, fmt.Errorf("unexpected Phrase instruments length; expected: %v, got: %v", len(col2Values), totalLength)
	}

	p := make([]Table, tableCount)
	for i := 0; i < tableCount; i++ {
		copy(p[i].Envelopes[:], envelopes[i*tableLength:tableLength*(i+1)])
		copy(p[i].Transposition[:], transpositions[i*tableLength:tableLength*(i+1)])
		copy(p[i].Col1.Command[:], col1Commands[i*tableLength:tableLength*(i+1)])
		copy(p[i].Col1.Value[:], col1Values[i*tableLength:tableLength*(i+1)])
		copy(p[i].Col2.Command[:], col2Commands[i*tableLength:tableLength*(i+1)])
		copy(p[i].Col2.Value[:], col2Values[i*tableLength:tableLength*(i+1)])
	}

	return p, nil
}
