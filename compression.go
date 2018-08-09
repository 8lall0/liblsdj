package liblsdj

import (
	"fmt"
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

func compress() {
	var defaultWaveLengthCount int

}
