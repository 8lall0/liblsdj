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
// TODO capire a cosa servano e come strutturarli meglio
type Bookmark [bookmarkPerChannelCount]byte

func setBookmarks(b []byte) ([]Bookmark, error) {
	if len(b) != (4)*bookmarkPerChannelCount {
		return nil, errors.New(fmt.Sprintf("unexpected phrase length: %v, %v", len(b), (4)*bookmarkPerChannelCount))
	}

	bo := make([]Bookmark, 4)
	for i := 0; i < 4; i++ {
		copy(bo[i][:], b[i:bookmarkPerChannelCount*i])
	}

	return bo, nil
}
