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

func decompressSAByte(r io.ReadSeeker, w io.WriteSeeker, curBlockPosition *int64, block1Position *int64, flag *bool) {
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
		if block1Position != nil {
			*curBlockPosition = *block1Position + (int64(b)-1)*int64(blockSize)
		} else {
			*curBlockPosition += blockSize
		}
		_, _ = r.Seek(*curBlockPosition, io.SeekCurrent)
	}
}

func decompress(r io.ReadSeeker, w io.WriteSeeker, block1position *int64) {
	wStart, _ := w.Seek(0, io.SeekCurrent)

	currentBlockPos, _ := r.Seek(0, io.SeekCurrent)

	b, _ := readByte(r)
	for loop := true; loop; {
		switch b {
		case runLengthEncodingByte:
			decompressRLEByte(r, w)
		case specialActionByte:
			decompressSAByte(r, w, &currentBlockPos, block1position, &loop)
		default:
			_ = writeByte(b, w)
		}
	}

	wEnd, _ := w.Seek(0, io.SeekCurrent)
	if (wEnd - wStart) != songDecompressedSize {
		fmt.Println("Decompressed size: ", wEnd-wStart)
	}

}

func compress(r io.ReadSeeker, w io.WriteSeeker, startBlock byte, blocksize int, blockCount int) int {
	if startBlock == byte(blockCount+1) {
		return 0
	}

	var b byte

	nextEvent := []byte{0, 0, 0}
	wStart, _ := w.Seek(0, io.SeekCurrent)
	curBlockSize := 0
	currentBlock := startBlock

	for i := 0; i < songDecompressedSize; i++ {
		// TODO capire while
		defWaveLengthCnt := byte(0)
		for j := 0; j < 0xff; j++ {

		}

		if defWaveLengthCnt > 0 {
			nextEvent = []byte{specialActionByte, defaultWaveByte, defWaveLengthCnt}
		} else {
			// Are we reading a default instrument? If so, we can compress these!
			// TODO capire while
			defInstrumentLengthCnt := byte(0)
			for j := 0; j < 0xff; j++ {

			}
			if defInstrumentLengthCnt > 0 {
				nextEvent = []byte{specialActionByte, defaultInstrumentByte, defInstrumentLengthCnt}
			} else {
				// Not a default wave, time to do "normal" compression

				b, _ = readByte(r)
				switch b {
				case runLengthEncodingByte:
					nextEvent = []byte{runLengthEncodingByte, runLengthEncodingByte}
					b, _ = readByte(r)
				case specialActionByte:
					nextEvent = []byte{specialActionByte, specialActionByte}
					b, _ = readByte(r)
				default:
					//pos, _ := r.Seek(0, io.SeekCurrent)
					c := b
					read1, _ := readByte(r)
					read2, _ := readByte(r)
					read3, _ := readByte(r)

					//TODO caso END
					// TODO sto codice fa SCHIFO e non so nemmeno se è CORRETTO PORCODIO
					if read1 == c && read2 == c && read3 == c {
						var cnt byte
						for cnt = 0; (read3 == c) && (cnt != 0xff); read3, _ = readByte(r) {
							cnt++
						}
						nextEvent = []byte{runLengthEncodingByte, c, cnt}
					} else {
						nextEvent = []byte{read3}
						read3, _ = readByte(r)
					}
				}
			}

			// blocksize o GLOBAL blocksize???
			if curBlockSize+len(nextEvent)+2 >= blocksize {
				_ = writeByte(specialActionByte, w)
				_ = writeByte(specialActionByte+currentBlock+1, w)

				curBlockSize += 2
				// assert curblocksize <= blockSize
				var zeroes [curBlockSize - blocksize]byte
				_, _ = w.Write(zeroes[:])

				currentBlock++
				curBlockSize = 0

				if currentBlock == blockCount+1 {
				}
			}
		}
	}
	_ = writeByte(specialActionByte, w)
	_ = writeByte(endOfFileByte, w)

	if curBlockSize > 0 {
		for curBlockSize += 2; curBlockSize < blockSize; curBlockSize++ {
			_ = writeByte(0, w)
		}
	}

	return int(currentBlock - startBlock + 1)
}
