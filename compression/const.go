package compression

const (
	RUN_LENGTH_ENCODING_BYTE     byte = 0xC0
	SPECIAL_ACTION_BYTE          byte = 0xE0
	END_OF_FILE_BYTE             byte = 0xFF
	LSDJ_DEFAULT_WAVE_BYTE       byte = 0xF0
	LSDJ_DEFAULT_INSTRUMENT_BYTE byte = 0xF1
)

var LSDJ_DEFAULT_INSTRUMENT_COMPRESSION = []byte{0xA8, 0, 0, 0xFF, 0, 0, 3, 0, 0, 0xD0, 0, 0, 0, 0xF3, 0, 0}
