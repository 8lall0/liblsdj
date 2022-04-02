package liblsdj

import "fmt"

const wordCount = 14
const wordValueLength = 0x60
const wordNameLength = 0xC

type Words []byte
type WordNames []byte

type Word struct {
	name  []byte
	value []byte
}

func setWords(names, values []byte) ([]Word, error) {
	if len(values) != wordCount*wordValueLength {
		return nil, fmt.Errorf("unexpected pippo length: %v, %v", len(values), wordCount*wordValueLength)
	} else if len(names) != wordCount*wordNameLength {
		return nil, fmt.Errorf("unexpected pluto length: %v, %v", len(names), wordCount*wordNameLength)
	}

	wo := make([]Word, wordCount)

	for i := 0; i < wordCount; i++ {
		wo[i].name = names[i*wordNameLength : (i+1)*wordNameLength]
		wo[i].value = names[i*wordValueLength : (i+1)*wordValueLength]
	}

	return wo, nil
}
