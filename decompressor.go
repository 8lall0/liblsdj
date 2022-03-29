package liblsdj

import (
	"github.com/orcaman/writerseeker"
	"io"
)

const (
	runLengthEncodingByte = 0xc0
	specialActionByte     = 0xe0
	endOfFileByte         = 0xff
	defaultWaveByte       = 0xf0
	defaultInstrumentByte = 0xf1
)

var defInstr = []byte{0xa8, 0, 0, 0xff, 0, 0, 3, 0, 0, 0xd0, 0, 0, 0, 0xf3, 0, 0}
var defWave = []byte{
	0x8E, 0xCD, 0xCC, 0xBB, 0xAA, 0xA9, 0x99, 0x88, 0x87, 0x76, 0x66, 0x55, 0x54, 0x43, 0x32, 0x31}

func decompressRLEByte(r io.ReadSeeker, w io.WriteSeeker) error {
	b, err := readByte(r)
	if err != nil {
		return err
	}

	if b != runLengthEncodingByte {
		cnt, err := readByte(r)
		if err != nil {
			return err
		}

		for i := 0; i < int(cnt); i++ {
			err := writeByte(b, w)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func decompressDefaultWaveByte(r io.ReadSeeker, w io.WriteSeeker) error {
	cnt, err := readByte(r)
	if err != nil {
		return err
	}

	for i := 0; i < int(cnt); i++ {
		_, err = w.Write(defWave[:])
		if err != nil {
			return err
		}
	}

	return nil
}

func decompressDefaultInstrumentByte(r io.ReadSeeker, w io.WriteSeeker) error {
	cnt, err := readByte(r)
	if err != nil {
		return err
	}

	for i := 0; i < int(cnt); i++ {
		_, err := w.Write(defInstr[:])
		if err != nil {
			return err
		}
	}

	return nil
}

func decompressSAByte(r io.ReadSeeker, w io.WriteSeeker, flag *bool) error {
	b, err := readByte(r)
	if err != nil {
		return err
	}

	switch b {
	case specialActionByte:
	case defaultWaveByte:
		err = decompressDefaultWaveByte(r, w)
	case defaultInstrumentByte:
		err = decompressDefaultInstrumentByte(r, w)
	case endOfFileByte:
		*flag = false
	default:
		_, err = r.Seek((int64(b)-1)*0x200, io.SeekStart)
	}

	if err != nil {
		return err
	}

	return nil
}

func decompress(r io.ReadSeeker, w io.WriteSeeker) error {
	var b byte

	_, err := r.Seek(0, io.SeekCurrent)
	if err != nil {
		return err
	}

	for loop := true; loop; {
		b, err = readByte(r)
		if err != nil {
			return err
		}

		switch b {
		case runLengthEncodingByte:
			err = decompressRLEByte(r, w)
		case specialActionByte:
			err = decompressSAByte(r, w, &loop)
		default:
			err = writeByte(b, w)
		}
		if err != nil {
			return err
		}
	}

	return nil
}

func Decompress(seeker io.ReadSeeker) (io.Reader, error) {
	out := &writerseeker.WriterSeeker{}

	err := decompress(seeker, out)
	if err != nil {
		return nil, err
	}

	return out.BytesReader(), nil
}
