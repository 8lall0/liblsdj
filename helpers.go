package liblsdj

import (
	"io"
)

func readByte(r io.ReadSeeker) (b byte, err error) {
	var p [1]byte
	_, err = r.Read(p[:])
	if err == nil {
		b = p[0]
	}
	return
}

func writeByte(b byte, w io.WriteSeeker) (err error) {
	var p [1]byte

	p[0] = b
	_, err = w.Write(p[:])

	return
}
