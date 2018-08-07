package compression

import (
	"fmt"
	"github.com/8lall0/liblsdj/vio"
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
		w.WriteSingle(LSDJ_DEFAULT_WAVE_BYTE)
	}
}

func decompressDefInstrumentByte(r *vio.Vio, w *vio.Vio) {
	var cnt byte

	cnt = r.ReadSingle()
	for i := 0; i < int(cnt); i++ {
		w.Write(LSDJ_DEFAULT_INSTRUMENT_COMPRESSION)
	}
}

func decompressSaByte(r *vio.Vio, w *vio.Vio, reading *bool, nBlock *int) {
	var b byte

	b = r.ReadSingle()
	if b == SPECIAL_ACTION_BYTE {
		fmt.Println("Special action byte")
		w.WriteSingle(b)
	} else if b == LSDJ_DEFAULT_WAVE_BYTE {
		decompressDefWaveByte(r, w)
		fmt.Println("Wave default byte")
	} else if b == LSDJ_DEFAULT_INSTRUMENT_BYTE {
		fmt.Println("Instr default byte")
		decompressDefInstrumentByte(r, w)
	} else if b == END_OF_FILE_BYTE {
		fmt.Println("EOF byte")
		*reading = false
	} else {
		fmt.Println("ELSE")
		*nBlock++
	}
}

func Decompress(r *vio.Vio, w *vio.Vio) {
	var b byte
	reading := true
	nBlock := 1

	for reading {
		b = r.ReadSingle()
		//fmt.Println(b)
		if b == RUN_LENGTH_ENCODING_BYTE {
			decompressRleByte(r, w)
			fmt.Println("rel")
		} else if b == SPECIAL_ACTION_BYTE {
			decompressSaByte(r, w, &reading, &nBlock)
			fmt.Println("sa")
		} else {
			w.WriteSingle(b)
			//fmt.Println(w.Get())
		}
	}
	fmt.Println(len(w.Get()), w.Cur())
}
