package liblsdj

type Lsdj_channel_list_t struct {
	pulse1 byte
	pulse2 byte
	wave   byte
	noise  byte
}
type Lsdj_row_t struct {
	chanIndex Lsdj_channel_list_t
	channels  [LSDJ_CHANNEL_COUNT]byte
}

// Clear all row data to factory settings
func (row Lsdj_row_t) Clear() {
	row.chanIndex.pulse1 = 0xFF
	row.chanIndex.pulse2 = 0xFF
	row.chanIndex.wave = 0xFF
	row.chanIndex.noise = 0xFF
}
