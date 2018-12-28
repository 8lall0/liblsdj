package liblsdj

import "io"

const channelCnt = 4

type row struct {
	pulse1 byte
	pulse2 byte
	wave   byte
	noise  byte
}

func (ro *row) clear() {
	ro.pulse1 = 0xFF
	ro.pulse2 = 0xFF
	ro.wave = 0xFF
	ro.noise = 0xFF
}

func (ro *row) write(w io.WriteSeeker) {
	_ = writeByte(ro.pulse1, w)
	_ = writeByte(ro.pulse2, w)
	_ = writeByte(ro.wave, w)
	_ = writeByte(ro.noise, w)
}

func (ro *row) read(r io.ReadSeeker) {
	ro.pulse1, _ = readByte(r)
	ro.pulse2, _ = readByte(r)
	ro.wave, _ = readByte(r)
	ro.noise, _ = readByte(r)
}
