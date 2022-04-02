package liblsdj

import (
	"errors"
	"fmt"
)

const (
	bookmarkPerChannelCount = 16
	noBookmarkValue         = 0xFF
)

// TODO trova costante del numero di canali, giusto per orientarti dal codice c.
type Bookmarks [4][bookmarkPerChannelCount]byte

func (bo *Bookmarks) Set(b []byte) error {
	if len(b) != (4)*bookmarkPerChannelCount {
		return errors.New(fmt.Sprintf("unexpected phrase length: %v, %v", len(b), (4)*bookmarkPerChannelCount))
	}

	for i := 0; i < 4; i++ {
		copy(bo[i][:], b[i:bookmarkPerChannelCount*i])
	}

	return nil
}
