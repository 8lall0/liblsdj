package liblsdj

const waveLen = 16

var defaultWave = [waveLen]byte{0x8E, 0xCD, 0xCC, 0xBB, 0xAA, 0xA9, 0x99, 0x88, 0x87, 0x76, 0x66, 0x55, 0x54, 0x43, 0x32, 0x31}

type wave [waveLen]byte

func (w *wave) clear() {
	*w = defaultWave
}
