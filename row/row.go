package row

import "github.com/8lall0/liblsdj/channel"

type bookMarkChannel struct {
	pulse1 byte
	pulse2 byte
	wave   byte
	noise  byte
}

type Row struct {
	chanIndex bookMarkChannel
	channels  [channel.LSDJ_CHANNEL_COUNT]byte
}

func (row Row) Clear() {
	row.chanIndex.Clear()
}

func (bookmark bookMarkChannel) Clear() {
	bookmark.pulse1 = 0xFF
	bookmark.pulse2 = 0xFF
	bookmark.wave = 0xFF
	bookmark.noise = 0xFF
}
