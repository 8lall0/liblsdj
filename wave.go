package liblsdj

import (
	"errors"
	"fmt"
)

const (
	waveCount         = 0xFF //! The amount of waves in a song
	wavePerSynthCount = 0xF  //! The amount of waves per synth
	waveByteCount     = 16   //! The number of bytes a wave takes /*! Do note that each step is represented by 4 bits, so the step count is twice this */
)

type Waves [4096]byte

func (w *Waves) Set(b []byte) error {
	if len(b) != 4096 {
		return errors.New(fmt.Sprintf("unexpected phrase length: %v, %v", len(b), 4096))
	}

	copy(w[:], b[:])

	return nil
}
