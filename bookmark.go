package liblsdj

import (
	"fmt"
)

const (
	bookmarkPerChannelCount = 16
	noBookmarkValue         = 0xFF
)

type Bookmark []byte

func setBookmarks(b []byte) ([]Bookmark, error) {
	if len(b) != (4)*bookmarkPerChannelCount {
		return nil, fmt.Errorf("unexpected Phrase length: %v, %v", len(b), (4)*bookmarkPerChannelCount)
	}

	bo := make([]Bookmark, 4)
	for i := 0; i < 4; i++ {
		bo[i] = b[i*bookmarkPerChannelCount : bookmarkPerChannelCount*(i+1)]
	}

	return bo, nil
}
