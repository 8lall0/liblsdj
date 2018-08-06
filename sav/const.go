package sav

import "github.com/8lall0/liblsdj/song"

const (
	LSDJ_NO_ACTIVE_PROJECT byte = 0xFF
	LSDJ_SAV_PROJECT_COUNT      = 32
	HEADER_START                = song.LSDJ_SONG_DECOMPRESSED_SIZE
	BLOCK_COUNT            int  = 191
	BLOCK_SIZE             byte = 0x200
)
