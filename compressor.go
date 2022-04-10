package liblsdj

import (
	"bytes"
	"github.com/orcaman/writerseeker"
	"io"
)

const blockSize = 0x200

func Compress(b []byte) (io.Reader, error) {
	return compress(b), nil
}

type BlockStruct struct {
	currentBlock     byte
	currentBlockSize int
	buffer           *writerseeker.WriterSeeker
}

func compress(b []byte) io.Reader {
	output := &BlockStruct{
		currentBlock:     1,
		currentBlockSize: 0,
		buffer:           &writerseeker.WriterSeeker{},
	}

	for i := 0; i < len(b); {
		cnt := 0
		for j := i; cnt < 0xFF && j+len(defWave) < len(b); j += len(defWave) {
			if bytes.Compare(b[j:j+len(defWave)], defWave) == 0 {
				cnt++
			} else {
				break
			}
		}
		if cnt > 0 {
			writeTo := []byte{specialActionByte, defaultWaveByte, byte(cnt)}
			output.writeToBuffer(writeTo)
			i += len(defWave) * cnt
			continue
		}

		cnt = 0
		for j := i; cnt < 0xFF && j+len(defInstr) < len(b); j += len(defInstr) {
			if bytes.Compare(b[j:j+len(defInstr)], defInstr) == 0 {
				cnt++
			} else {
				break
			}
		}
		if cnt > 0 {
			writeTo := []byte{specialActionByte, defaultInstrumentByte, byte(cnt)}
			output.writeToBuffer(writeTo)
			i += len(defInstr) * cnt
			continue
		}

		cnt = 1
		for j := i + 1; j < len(b) && cnt < 0xFF && b[j-1] == b[j]; j++ {
			cnt++
		}
		if cnt > 2 {
			writeTo := []byte{runLengthEncodingByte, b[i], byte(cnt)}
			output.writeToBuffer(writeTo)
			i += cnt
			continue
		}

		writeTo := []byte{b[i]}
		output.writeToBuffer(writeTo)
		i++
	}

	output.writeToBuffer([]byte{specialActionByte, endOfFileByte})

	return output.buffer.BytesReader()
}

func (b *BlockStruct) writeToBuffer(input []byte) error {
	if b.currentBlockSize+len(input)+2 >= blockSize {
		if _, err := b.buffer.Write([]byte{specialActionByte, b.currentBlock}); err != nil {
			return err
		}
		b.currentBlockSize += 2
		zeroes := make([]byte, blockSize-b.currentBlockSize)
		if _, err := b.buffer.Write(zeroes); err != nil {
			return err
		}
		b.currentBlock++
		b.currentBlockSize = 0
	}

	b.buffer.Write(input)
	b.currentBlockSize += len(input)

	return nil
}
