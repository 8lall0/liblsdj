package compression

import (
	"fmt"
	"github.com/8lall0/liblsdj/vio"
	"github.com/8lall0/liblsdj/wave"
)

func decompressRleByte(r *vio.Vio, w *vio.Vio) {
	var b, cnt byte
	b = r.ReadSingle()
	if b == RUN_LENGTH_ENCODING_BYTE {
		w.WriteSingle(b)
	} else {
		cnt = r.ReadSingle()
		for i := 0; i < int(cnt); i++ {
			w.WriteSingle(b)
		}
	}
}

func decompressDefWaveByte(r *vio.Vio, w *vio.Vio) {
	var cnt byte

	cnt = r.ReadSingle()
	for i := 0; i < int(cnt); i++ {
		w.Write(wave.LSDJ_DEFAULT_WAVE)
	}
}

func decompressDefInstrumentByte(r *vio.Vio, w *vio.Vio) {
	var cnt byte

	cnt = r.ReadSingle()
	for i := 0; i < int(cnt); i++ {
		w.Write(LSDJ_DEFAULT_INSTRUMENT_COMPRESSION)
	}
}

func decompressSaByte(r *vio.Vio, w *vio.Vio, reading *bool) {
	var b byte

	b = r.ReadSingle()
	if b == SPECIAL_ACTION_BYTE {
		w.WriteSingle(b)
	} else if b == LSDJ_DEFAULT_WAVE_BYTE {
		decompressDefWaveByte(r, w)
	} else if b == LSDJ_DEFAULT_INSTRUMENT_BYTE {
		decompressDefInstrumentByte(r, w)
	} else if b == END_OF_FILE_BYTE {
		*reading = false
	} else {
		fmt.Println(b - 1)
	}
}

func Decompress(r *vio.Vio, w *vio.Vio) {
	var b byte
	reading := true

	for reading {
		b = r.ReadSingle()
		if b == RUN_LENGTH_ENCODING_BYTE {
			decompressRleByte(r, w)
		} else if b == SPECIAL_ACTION_BYTE {
			decompressSaByte(r, w, &reading)
		} else {
			w.WriteSingle(b)
		}
	}
	fmt.Println(len(w.Get()), w.Cur())

}
