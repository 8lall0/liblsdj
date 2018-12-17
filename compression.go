package liblsdj

import (
	"fmt"
	"io"
)

const (
	runLengthEncodingByte = 0xc0
	specialActionByte     = 0xe0
	endOfFileByte         = 0xff
	defaultWaveByte       = 0xf0
	defaultInstrumentByte = 0xf1
)

var defaultInstrumentCompression = [instrumentDefaultLen]byte{0xA8, 0, 0, 0xFF, 0, 0, 3, 0, 0, 0xD0, 0, 0, 0, 0xF3, 0, 0}

func decompressRLEByte(r io.ReadSeeker, w io.WriteSeeker) {
	b, _ := readByte(r)
	// può panicare

	if b == runLengthEncodingByte {
		_ = writeByte(b, w)
	} else {
		cnt, _ := readByte(r)
		for i := 0; i < int(cnt); i++ {
			_ = writeByte(b, w)
		}
	}
}

func decompressDefaultWaveByte(r io.ReadSeeker, w io.WriteSeeker) {
	cnt, _ := readByte(r)
	for i := 0; i < int(cnt); i++ {
		_, _ = w.Write(defaultWave[:])
	}
}

func decompressDefaultInstrumentByte(r io.ReadSeeker, w io.WriteSeeker) {
	cnt, _ := readByte(r)
	for i := 0; i < int(cnt); i++ {
		_, _ = w.Write(defaultInstrumentCompression[:])
	}
}

func decompressSAByte(r io.ReadSeeker, w io.WriteSeeker, flag *bool) {
	b, _ := readByte(r)

	switch b {
	case specialActionByte:
		_ = writeByte(b, w)
	case defaultWaveByte:
		decompressDefaultWaveByte(r, w)
	case defaultInstrumentByte:
		decompressDefaultInstrumentByte(r, w)
	case endOfFileByte:
		// qui finisce il suo percorso
		*flag = false
	default:
		// TODO currentblockposition
	}
}

func decompress(r io.ReadSeeker, w io.WriteSeeker, block1position int64, blocksize int) {
	wStart, _ := w.Seek(0, io.SeekCurrent)
	// TODO currentblockposition
	//currentBlockPos, _ := r.Seek(0, io.SeekCurrent)

	b, _ := readByte(r)
	for loop := true; loop; {
		switch b {
		case runLengthEncodingByte:
			decompressRLEByte(r, w)
		case specialActionByte:
			decompressSAByte(r, w, &loop)
		default:
			_ = writeByte(b, w)
		}
	}

	wEnd, _ := w.Seek(0, io.SeekCurrent)
	if (wEnd - wStart) != songDecompressedSize {
		fmt.Println("Decompressed size: ", wEnd-wStart)
	}

}

func Compress(r io.ReadSeeker, w io.WriteSeeker, startBlock byte, blocksize int, blockCount int) byte {
	if startBlock == byte(blockCount+1) {
		return 0
	}

	nextEvent := []byte{0, 0, 0}
	eventSize := 0
	b := 0

	wStart, _ := w.Seek(0, io.SeekCurrent)

	return currentBlock - startBlock + 1
}
