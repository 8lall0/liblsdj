package lsdj

import "io"

const channelCnt = 4

type row struct {
	pulse1   byte
	pulse2   byte
	wave     byte
	noise    byte
}

func (ro *row) clear() {
	ro.pulse1 = 0xFF
	ro.pulse2 = 0xFF
	ro.wave = 0xFF
	ro.noise = 0xFF
}

func (ro *row) write(r io.ReadSeeker) {
	// TODO errori
	ro.pulse1, _ = readByte(r)
	ro.pulse2, _ = readByte(r)
	ro.wave, _ = readByte(r)
	ro.noise, _ = readByte(r)
}
