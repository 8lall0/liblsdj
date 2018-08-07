package song

import (
	"github.com/8lall0/liblsdj/phrase"
	"github.com/8lall0/liblsdj/table"
	"github.com/8lall0/liblsdj/word"
)

const (
	LSDJ_SONG_DECOMPRESSED_SIZE  int  = 32768 //0x8000
	LSDJ_ROW_COUNT               int  = 256
	LSDJ_CHAIN_COUNT             int  = 128
	LSDJ_PHRASE_COUNT            byte = 0xFF
	LSDJ_INSTRUMENT_COUNT        int  = 64
	LSDJ_SYNTH_COUNT             int  = 16
	LSDJ_TABLE_COUNT             int  = 32
	LSDJ_WAVE_COUNT              int  = 256
	LSDJ_GROOVE_COUNT            int  = 32
	LSDJ_WORD_COUNT              int  = 42
	LSDJ_BOOKMARK_POSITION_COUNT int  = 16
	LSDJ_NO_BOOKMARK             byte = 0xFF

	LSDJ_CLONE_DEEP byte = 0
	LSDJ_CLONE_SLIM byte = 1
)

var DEFAULT_WORD_NAMES = [LSDJ_WORD_COUNT][word.LSDJ_WORD_NAME_LENGTH]byte{
	{'C', ' ', '2', ' '},
	{'C', ' ', '2', ' '},
	{'D', ' ', '2', ' '},
	{'D', ' ', '2', ' '},
	{'E', ' ', '2', ' '},
	{'F', ' ', '2', ' '},
	{'F', ' ', '2', ' '},
	{'G', ' ', '2', ' '},
	{'G', ' ', '2', ' '},
	{'A', ' ', '2', ' '},
	{'A', ' ', '2', ' '},
	{'B', ' ', '2', ' '},
	{'C', ' ', '3', ' '},
	{'C', ' ', '3', ' '},
	{'D', ' ', '3', ' '},
	{'D', ' ', '3', ' '},
	{'E', ' ', '3', ' '},
	{'F', ' ', '3', ' '},
	{'F', ' ', '3', ' '},
	{'G', ' ', '3', ' '},
	{'G', ' ', '3', ' '},
	{'A', ' ', '3', ' '},
	{'A', ' ', '3', ' '},
	{'B', ' ', '3', ' '},
	{'C', ' ', '4', ' '},
	{'C', ' ', '4', ' '},
	{'D', ' ', '4', ' '},
	{'D', ' ', '4', ' '},
	{'E', ' ', '4', ' '},
	{'F', ' ', '4', ' '},
	{'F', ' ', '4', ' '},
	{'G', ' ', '4', ' '},
	{'G', ' ', '4', ' '},
	{'A', ' ', '4', ' '},
	{'A', ' ', '4', ' '},
	{'B', ' ', '4', ' '},
	{'C', ' ', '5', ' '},
	{'C', ' ', '5', ' '},
	{'D', ' ', '5', ' '},
	{'D', ' ', '5', ' '},
	{'E', ' ', '5', ' '},
	{'F', ' ', '5', ' '},
}

const (
	INSTR_ALLOC_TABLE_SIZE  int = 64
	TABLE_ALLOC_TABLE_SIZE  int = 32
	CHAIN_ALLOC_TABLE_SIZE  int = 16
	PHRASE_ALLOC_TABLE_SIZE int = 32
)

var LSDJ_TABLE_LENGTH_ZERO = [table.LSDJ_TABLE_LENGTH]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
var LSDJ_CHAIN_LENGTH_ZERO = [table.LSDJ_TABLE_LENGTH]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
var LSDJ_CHAIN_LENGTH_FF = [table.LSDJ_TABLE_LENGTH]byte{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF}
var LSDJ_PHRASE_LENGTH_ZERO = [phrase.LSDJ_PHRASE_LENGTH]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
var LSDJ_PHRASE_LENGTH_FF = [phrase.LSDJ_PHRASE_LENGTH]byte{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF}
