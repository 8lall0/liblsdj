package liblsdj

import (
	"fmt"
)

const (
	waveCount         = 0xFF //! The amount of waves in a song
	wavePerSynthCount = 0xF  //! The amount of waves per synth
	waveByteCount     = 16   //! The number of bytes a wave takes /*! Do note that each step is represented by 4 bits, so the step count is twice this */
)

type Wave [16]byte

func setWaves(waves []byte) ([]Wave, error) {
	if len(waves) != 4096 {
		return nil, fmt.Errorf("unexpected Phrase length: %v, %v", len(waves), 4096)
	}

	wv := make([]Wave, 4096/16)
	for i := 0; i < 4096/16; i++ {
		copy(wv[i][:], waves[waveByteCount*i:waveByteCount*(i+1)])
	}

	return wv, nil
}
