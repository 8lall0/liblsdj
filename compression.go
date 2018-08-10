package liblsdj

import (
	"fmt"
	"github.com/tango-contrib/cache"
)

const (
	RUN_LENGTH_ENCODING_BYTE     byte = 0xC0
	SPECIAL_ACTION_BYTE          byte = 0xE0
	END_OF_FILE_BYTE             byte = 0xFF
	LSDJ_DEFAULT_WAVE_BYTE       byte = 0xF0
	LSDJ_DEFAULT_INSTRUMENT_BYTE byte = 0xF1
)

var lsdj_DEFAULT_INSTRUMENT_COMPRESSION = []byte{0xA8, 0, 0, 0xFF, 0, 0, 3, 0, 0, 0xD0, 0, 0, 0, 0xF3, 0, 0}

func decompressRleByte(r *vio, w *vio) {
	var b, cnt byte
	b = r.readByte()
	if b == RUN_LENGTH_ENCODING_BYTE {
		w.writeByte(b)
	} else {
		cnt = r.readByte()
		for i := 0; i < int(cnt); i++ {
			w.writeByte(b)
		}
	}
}

func decompressDefWaveByte(r *vio, w *vio) {
	var cnt byte

	cnt = r.readByte()
	for i := 0; i < int(cnt); i++ {
		w.write(lsdj_DEFAULT_WAVE)
	}
}

func decompressDefInstrumentByte(r *vio, w *vio) {
	var cnt byte

	cnt = r.readByte()
	for i := 0; i < int(cnt); i++ {
		w.write(lsdj_DEFAULT_INSTRUMENT_COMPRESSION)
	}
}

func decompressSaByte(r *vio, w *vio, reading *bool, offset int, nBlock *int) {
	var b byte

	b = r.readByte()
	if b == SPECIAL_ACTION_BYTE {
		w.writeByte(b)
	} else if b == LSDJ_DEFAULT_WAVE_BYTE {
		decompressDefWaveByte(r, w)
	} else if b == LSDJ_DEFAULT_INSTRUMENT_BYTE {
		decompressDefInstrumentByte(r, w)
	} else if b == END_OF_FILE_BYTE {
		*reading = false
	} else {
		r.seek(offset + (512 * (*nBlock)))
		*nBlock++
	}
}

func decompress(r *vio, w *vio) {
	var b byte
	reading := true
	nBlock := 1
	curPos := r.getCur()

	for reading {
		b = r.readByte()
		if b == RUN_LENGTH_ENCODING_BYTE {
			decompressRleByte(r, w)
		} else if b == SPECIAL_ACTION_BYTE {
			decompressSaByte(r, w, &reading, curPos, &nBlock)
		} else {
			w.writeByte(b)
		}
	}
	fmt.Println("Size: ", len(w.get()))
}

func compress(r *vio, w *vio) int {
	var end = r.getCur() + lsdj_SONG_DECOMPRESSED_SIZE
	var defWaveCnt, defInsCnt byte

	nextEvent := []byte{0, 0, 0}
	eventSize := 0

	// So già quante cazzo di volte devo ciclare
	for i := 0; i < lsdj_SONG_DECOMPRESSED_SIZE; i++ {
		// Controllo per la wave
		for isDefault := true; isDefault && defWaveCnt != 0xFF; {
			tmp := r.read(lsdj_WAVE_LENGTH)
			for j := 0; isDefault && j < lsdj_WAVE_LENGTH; j++ {
				isDefault = (tmp[j] == lsdj_DEFAULT_WAVE[j])
			}
			if isDefault {
				defWaveCnt++
			} else {
				// Torno indietro
				r.seekCur(-lsdj_WAVE_LENGTH)
			}
		}
		if defWaveCnt > 0 {
			nextEvent[0] = SPECIAL_ACTION_BYTE
			nextEvent[1] = LSDJ_DEFAULT_INSTRUMENT_BYTE
			nextEvent[2] = defWaveCnt
			eventSize = 3
		} else {
			//Controllo per gli strumenti
			for isDefault := true; isDefault && defInsCnt != 0xFF; {
				tmp := r.read(lsdj_DEFAULT_INSTRUMENT_LENGTH)
				for j := 0; isDefault && j < lsdj_DEFAULT_INSTRUMENT_LENGTH; j++ {
					isDefault = (tmp[j] == lsdj_DEFAULT_INSTRUMENT_COMPRESSION[j])
				}
				if isDefault {
					defInsCnt++
				} else {
					// Torno indietro
					r.seekCur(-lsdj_DEFAULT_INSTRUMENT_LENGTH)
				}
			}
			if defInsCnt > 0 {
				nextEvent[0] = SPECIAL_ACTION_BYTE
				nextEvent[1] = LSDJ_DEFAULT_INSTRUMENT_BYTE
				nextEvent[2] = defInsCnt
				eventSize = 3
			} else {
				b := r.readByte()
				if b == RUN_LENGTH_ENCODING_BYTE {
					nextEvent[0] = RUN_LENGTH_ENCODING_BYTE
					nextEvent[1] = RUN_LENGTH_ENCODING_BYTE
					eventSize = 2
				} else if b == RUN_LENGTH_ENCODING_BYTE {
					nextEvent[0] = SPECIAL_ACTION_BYTE
					nextEvent[1] = SPECIAL_ACTION_BYTE
					eventSize = 2
				} else {
					r.seekCur(-1)
					cur := r.getCur()
					if cur+3 < end &&
						r.readByte() == b &&
						r.readByte() == b &&
						r.readByte() == b {
						cnt := byte(0)
						for r.getCur() < end &&
							r.readByte() == b &&
							cnt != 0xFF {
							cnt++
						}
						nextEvent[0] = RUN_LENGTH_ENCODING_BYTE
						nextEvent[1] = b
						nextEvent[2] = cnt
						eventSize = 3
					} else {
						nextEvent[0] = b
						eventSize = 1
						r.seekCur(1)
					}
				}
			}
		}

		// Controlla se il blocco è esaurito
		// TODO: vio per gestire blocchi??
		if (w.getLen() + eventSize) >= BLOCK_SIZE {
			w.writeByte(SPECIAL_ACTION_BYTE)
			w.writeByte(SPECIAL_ACTION_BYTE)
		}

		w.writeByte(SPECIAL_ACTION_BYTE)
		w.writeByte(END_OF_FILE_BYTE)
		// Write all zeroes
	}

}
