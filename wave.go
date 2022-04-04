package liblsdj

import (
	"fmt"
)

const (
	waveCount         = 0xFF + 1 //! The amount of waves in a song
	wavePerSynthCount = 0xF      //! The amount of waves per synth
	waveByteCount     = 16       //! The number of bytes a wave takes /*! Do note that each step is represented by 4 bits, so the step count is twice this */
)

type Wave [waveByteCount]byte

func setWaves(waves []byte) ([]Wave, error) {
	if len(waves) != waveCount*waveByteCount {
		return nil, fmt.Errorf("unexpected Phrase length: %v, %v", len(waves), waveCount*waveByteCount)
	}

	wv := make([]Wave, waveCount)
	for i := 0; i < waveCount*waveByteCount/waveByteCount; i++ {
		copy(wv[i][:], waves[waveByteCount*i:waveByteCount*(i+1)])
	}

	return wv, nil
}
