package liblsdj

type row struct {
	channelList struct {
		pulse1 byte
		pulse2 byte
		wave   byte
		noise  byte
	}
	channels []byte //lsdj_CHANNEL_COUNT
}

func (row row) clear() {
	row.channelList.pulse1 = 0xFF
	row.channelList.pulse2 = 0xFF
	row.channelList.wave = 0xFF
	row.channelList.noise = 0xFF
}
