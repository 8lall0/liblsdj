package row

import "github.com/8lall0/liblsdj/channel"

type Row struct {
	channelList struct {
		pulse1 byte
		pulse2 byte
		wave   byte
		noise  byte
	}
	channels [channel.LSDJ_CHANNEL_COUNT]byte
}

func (row Row) Clear() {
	row.channelList.pulse1 = 0xFF
	row.channelList.pulse2 = 0xFF
	row.channelList.wave = 0xFF
	row.channelList.noise = 0xFF
}
