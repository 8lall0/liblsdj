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
			*curBlockPosition += int64(blockSize)
		}
		_, _ = r.Seek(*curBlockPosition, io.SeekStart)
	}
}

func decompress(r io.ReadSeeker, w io.WriteSeeker, block1position *int64) {
	var b byte

	wStart, _ := w.Seek(0, io.SeekCurrent)
	currentBlockPos, _ := r.Seek(0, io.SeekCurrent)

	for loop := true; loop; {
		b, _ = readByte(r)
		fmt.Println(b)

		switch b {
		case runLengthEncodingByte:
			fmt.Println("RLE")
			decompressRLEByte(r, w)
		case specialActionByte:
			fmt.Println("SA")
			decompressSAByte(r, w, &currentBlockPos, block1position, &loop)
		default:
			fmt.Println("DEF")
			_ = writeByte(b, w)
		}
	}

	wEnd, _ := w.Seek(0, io.SeekCurrent)
	if (wEnd - wStart) != songDecompressedSize {
		fmt.Println("Decompressed size: ", wEnd-wStart, " Normal size: ", songDecompressedSize)
	}

}

func compress(r io.ReadSeeker, w io.WriteSeeker, startBlock byte, blocksize int, blockCount int) int {
	if startBlock == byte(blockCount+1) {
		return 0
	}

	var b byte
	// TODO temporanea, poi controlla meglio
	var readWave [waveLen]byte
	var readInstr [instrumentDefaultLen]byte

	nextEvent := []byte{0, 0, 0}
	wStart, _ := w.Seek(0, io.SeekCurrent)
	curBlockSize := 0
	currentBlock := startBlock

	// TODO poiché non ho aritmetica dei puntatori, devo incrementare la i per i cazzi miei
	// TODO forse posso usare la positione con seek, ma me la devo controllare un attimo
	for i := 0; i < songDecompressedSize; {
		// Are we reading a default wave? If so, we can compress these!
		// TODO Questo legge le default wave
		defWaveLengthCnt := byte(0)
		_, _ = r.Read(readWave[:])
		i += waveLen // manuale incremento i
		for (i+waveLen < songDecompressedSize) &&
			(readWave == defaultWave) &&
			(defWaveLengthCnt != 0xff) {

			defWaveLengthCnt++
			_, _ = r.Read(readWave[:])
			i += waveLen // manuale incremento i
		}

		//Forse devo seekare uno indietro

		if defWaveLengthCnt > 0 {
			nextEvent = []byte{specialActionByte, defaultWaveByte, defWaveLengthCnt}
		} else {
			// Are we reading a default instrument? If so, we can compress these!
			// TODO Questo legge le default instr
			defInstrumentLengthCnt := byte(0)
			_, _ = r.Read(readInstr[:])
			i += instrumentDefaultLen // manuale incremento i
			for (i+instrumentDefaultLen < songDecompressedSize) &&
				(readInstr == instrumentDefault) &&
				(defInstrumentLengthCnt != 0xff) {

				defInstrumentLengthCnt++
				_, _ = r.Read(readInstr[:])
				i += instrumentDefaultLen // manuale incremento i
			}
			//Forse devo seekare uno indietro

			if defInstrumentLengthCnt > 0 {
				nextEvent = []byte{specialActionByte, defaultInstrumentByte, defInstrumentLengthCnt}
			} else {
				// Not a default wave, time to do "normal" compression
				b, _ = readByte(r)
				i++ // manuale incremento i
				switch b {
				case runLengthEncodingByte:
					nextEvent = []byte{runLengthEncodingByte, runLengthEncodingByte}
					b, _ = readByte(r)
					i++ // manuale incremento i
				case specialActionByte:
					nextEvent = []byte{specialActionByte, specialActionByte}
					b, _ = readByte(r)
					i++ // manuale incremento i
				default:
					c := b
					if i+3 < songDecompressedSize {
						read1, _ := readByte(r)
						i++ // manuale incremento i
						read2, _ := readByte(r)
						i++ // manuale incremento i
						read3, _ := readByte(r)
						i++ // manuale incremento i

						//TODO caso END
						// TODO sto codice fa SCHIFO e non so nemmeno se è CORRETTO PORCODIO
						if read1 == c && read2 == c && read3 == c {
							var cnt byte
							for cnt = 0; (read3 == c) && (cnt != 0xff); {
								read3, _ = readByte(r)
								i++ // manuale incremento i
								cnt++
							}
							nextEvent = []byte{runLengthEncodingByte, c, cnt}
						} else {
							nextEvent = []byte{read3}
							read3, _ = readByte(r)
							i++ // manuale incremento i
						}
					}

				}
			}

			// blocksize o GLOBAL blocksize???
			if curBlockSize+len(nextEvent)+2 >= blocksize {
				_ = writeByte(specialActionByte, w)
				_ = writeByte(specialActionByte+currentBlock+1, w)

				curBlockSize += 2
				// assert curblocksize <= blockSize
				zeroes := make([]byte, curBlockSize-blocksize)
				_, _ = w.Write(zeroes)

				currentBlock++
				curBlockSize = 0

				// Have we reached the maximum block count?
				// If so, roll back
				if currentBlock == byte(blockCount)+1 {
					pos, _ := w.Seek(0, io.SeekCurrent)
					zeroes := make([]byte, pos-wStart)
					_, _ = w.Write(zeroes)
				}

				// seek back
				_, _ = w.Seek(wStart, io.SeekStart)
				return 0
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
