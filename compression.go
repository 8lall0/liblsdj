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
		fmt.Println("Decompressed size: ", wEnd-wStart, " Normal size: ", songDecompressedSize)
	}

}

func compress(r io.ReadSeeker, w io.WriteSeeker, startBlock int) int {
	if startBlock == blockCnt+1 {
		return 0
	}
	var sel byte

	curBlock := startBlock
	curBlockSize := 0

	wStart, _ := w.Seek(0, io.SeekCurrent)
	rStart, _ := r.Seek(0, io.SeekCurrent)
	rEnd := rStart + songDecompressedSize

	// TODO controlla sto cazzo di loop
	for posRead, _ := r.Seek(0, io.SeekCurrent); posRead < rEnd; posRead, _ = r.Seek(0, io.SeekCurrent) {
		nextEvent := []byte{0, 0, 0}

		// Are we reading a default wave? If so, we can compress these!
		var defWaveLengthCnt byte
		var readWave [waveLen]byte
		for posWave, _ := r.Seek(0, io.SeekCurrent); (posWave+waveLen < rEnd) && (defWaveLengthCnt != 0xff); posWave, _ = r.Seek(0, io.SeekCurrent) {
			_, _ = r.Read(readWave[:])
			if readWave == defaultWave {
				defWaveLengthCnt++
			} else {
				_, _ = r.Seek(-waveLen, io.SeekCurrent)
				break
			}
		}

		if defWaveLengthCnt > 0 {
			nextEvent = []byte{specialActionByte, defaultWaveByte, defWaveLengthCnt}
		} else {
			// Are we reading a default instrument? If so, we can compress these!
			var defInstrumentLengthCnt byte
			var readInstr [instrumentDefaultLen]byte
			for posInstr, _ := r.Seek(0, io.SeekCurrent); (posInstr+instrumentDefaultLen < rEnd) && (defInstrumentLengthCnt != 0xff); posInstr, _ = r.Seek(0, io.SeekCurrent) {
				_, _ = r.Read(readInstr[:])
				if readInstr == instrumentDefault {
					defInstrumentLengthCnt++
				} else {
					_, _ = r.Seek(-instrumentDefaultLen, io.SeekCurrent)
					break
				}
			}

			if defInstrumentLengthCnt > 0 {
				nextEvent = []byte{specialActionByte, defaultInstrumentByte, defInstrumentLengthCnt}
			} else {
				// Not a default wave, time to do "normal" compression
				//TODO INGHIPPO

				sel, _ = readByte(r)
				switch sel {
				case runLengthEncodingByte:
					nextEvent = []byte{runLengthEncodingByte, runLengthEncodingByte}
					sel, _ = readByte(r)
				case specialActionByte:
					nextEvent = []byte{specialActionByte, specialActionByte}
					sel, _ = readByte(r)
				default:
					// See if we can do run-length encoding
					posRun, _ := r.Seek(0, io.SeekCurrent)
					if posRun+3 < rEnd {
						read1, _ := readByte(r)
						read2, _ := readByte(r)
						read3, _ := readByte(r)
						_, _ = r.Seek(-3, io.SeekCurrent)

						// TODO sto codice fa SCHIFO e non so nemmeno se è CORRETTO PORCODIO
						if read1 == sel && read2 == sel && read3 == sel {
							var cnt byte
							for ; (read3 == sel) && (cnt != 0xff); cnt++ {
								read3, _ = readByte(r)
							}
							nextEvent = []byte{runLengthEncodingByte, sel, cnt}
						}
					} else {
						tmp, _ := readByte(r)
						nextEvent = []byte{tmp}
					}
				}
			}
		}

		// See if the event would still fit in this block
		// If not, move to a new block
		if curBlockSize+len(nextEvent)+2 >= blockSize {
			// Write the "next block" command
			_ = writeByte(specialActionByte, w)
			_ = writeByte(byte(curBlock)+1, w)

			curBlockSize += 2
			// assert curblocksize <= blockSize
			zeroes := make([]byte, blockSize-curBlockSize)
			_, _ = w.Write(zeroes)

			curBlock += 1
			curBlockSize = 0

			// Have we reached the maximum block count?
			// If so, roll back
			if curBlock == blockCnt+1 {
				posRollback, _ := w.Seek(0, io.SeekCurrent)
				_, _ = w.Seek(wStart, io.SeekStart)

				zeroes := make([]byte, posRollback-wStart)
				_, _ = w.Write(zeroes)

				// seek back
				_, _ = w.Seek(wStart, io.SeekStart)
				return 0
			}
		}
		_, _ = w.Write(nextEvent)
		curBlockSize += len(nextEvent)
	}

	_ = writeByte(specialActionByte, w)
	_ = writeByte(endOfFileByte, w)

	zeroes := make([]byte, blockSize-curBlockSize)
	_, _ = w.Write(zeroes)

	return int(curBlock - startBlock + 1)
}
