package liblsdj

const (
	LSDJ_SONG_DECOMPRESSED_SIZE byte = 0x8000 
	LSDJ_ROW_COUNT int = 256
	LSDJ_CHAIN_COUNT int = 128
	LSDJ_PHRASE_COUNT byte = 0xFF
	LSDJ_INSTRUMENT_COUNT int = 64
	LSDJ_SYNTH_COUNT int = 16
	LSDJ_TABLE_COUNT int = 32
	LSDJ_WAVE_COUNT int = 256
	LSDJ_GROOVE_COUNT int = 32
	LSDJ_WORD_COUNT int = 42
	LSDJ_BOOKMARK_POSITION_COUNT int = 16
	LSDJ_NO_BOOKMARK byte = 0xFF

	LSDJ_CLONE_DEEP byte = 0
	LSDJ_CLONE_SLIM byte = 1
)

var DEFAULT_WORD_NAMES = [LSDJ_WORD_COUNT][LSDJ_WORD_NAME_LENGTH]byte {
	{'C',' ','2',' '},
	{'C',' ','2',' '},
	{'D',' ','2',' '},
	{'D',' ','2',' '},
	{'E',' ','2',' '},
	{'F',' ','2',' '},
	{'F',' ','2',' '},
	{'G',' ','2',' '},
	{'G',' ','2',' '},
	{'A',' ','2',' '},
	{'A',' ','2',' '},
	{'B',' ','2',' '},
	{'C',' ','3',' '},
	{'C',' ','3',' '},
	{'D',' ','3',' '},
	{'D',' ','3',' '},
	{'E',' ','3',' '},
	{'F',' ','3',' '},
	{'F',' ','3',' '},
	{'G',' ','3',' '},
	{'G',' ','3',' '},
	{'A',' ','3',' '},
	{'A',' ','3',' '},
	{'B',' ','3',' '},
	{'C',' ','4',' '},
	{'C',' ','4',' '},
	{'D',' ','4',' '},
	{'D',' ','4',' '},
	{'E',' ','4',' '},
	{'F',' ','4',' '},
	{'F',' ','4',' '},
	{'G',' ','4',' '},
	{'G',' ','4',' '},
	{'A',' ','4',' '},
	{'A',' ','4',' '},
	{'B',' ','4',' '},
	{'C',' ','5',' '},
	{'C',' ','5',' '},
	{'D',' ','5',' '},
	{'D',' ','5',' '},
	{'E',' ','5',' '},
	{'F',' ','5',' '},
}

const(
	INSTR_ALLOC_TABLE_SIZE int = 64
	TABLE_ALLOC_TABLE_SIZE int = 32
	CHAIN_ALLOC_TABLE_SIZE int = 16
	PHRASE_ALLOC_TABLE_SIZE int = 32
)

var LSDJ_TABLE_LENGTH_ZERO = [LSDJ_TABLE_LENGTH]byte{ 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0 }
var LSDJ_CHAIN_LENGTH_ZERO = [LSDJ_TABLE_LENGTH]byte{ 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0 }
var LSDJ_CHAIN_LENGTH_FF = [LSDJ_TABLE_LENGTH]byte{ 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF}
var LSDJ_PHRASE_LENGTH_ZERO = [LSDJ_PHRASE_LENGTH]byte{ 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0 }
var LSDJ_PHRASE_LENGTH_FF = [LSDJ_PHRASE_LENGTH]byte{ 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF}


type lsdj_song_bookmark_chans struct {
	pulse1 [LSDJ_BOOKMARK_POSITION_COUNT]byte
	pulse2 [LSDJ_BOOKMARK_POSITION_COUNT]byte
	wave [LSDJ_BOOKMARK_POSITION_COUNT]byte
	noise [LSDJ_BOOKMARK_POSITION_COUNT]byte
}

type lsdj_song_totalTime struct {
	days byte
	hours byte
	minutes byte
}

type lsdj_song_workTime struct {
	hours byte
	minutes byte
}

type lsdj_song_metadata struct {
	keyDelay byte
	keyRepeat byte
	font byte
	sync byte
	colorSet byte
	clone byte
	fileChangedFlag byte
	powerSave byte
	preListen byte
	workTime *lsdj_song_workTime
	totalTime *lsdj_song_totalTime
}

// An LSDJ song
type Lsdj_song_t struct {
	formatVersion byte
	tempo byte
	transposition byte
	drumMax byte
	// The sequences of chains in the song
	rows [LSDJ_ROW_COUNT]*Row
	// The chains in the song
	chains [LSDJ_CHAIN_COUNT]*Chain
	// The prases in the song
	phrases [LSDJ_PHRASE_COUNT]*Phrase
	// Instruments of the song
	instruments [LSDJ_INSTRUMENT_COUNT]Lsdj_instrument_t
	// Soft synths of the song
	synths [LSDJ_SYNTH_COUNT]*Synth
	// Wave frames of the song
	waves [LSDJ_WAVE_COUNT]*Wave
	// The tables in the song
	tables[LSDJ_TABLE_COUNT]*Table
	// The grooves in the song
	grooves [LSDJ_GROOVE_COUNT]*Groove
	// The speech synth words in the song
	words [LSDJ_WORD_COUNT]*Word
	wordNames [LSDJ_WORD_COUNT][LSDJ_WORD_NAME_LENGTH]string

	// Bookmarks
	bookmark *lsdj_song_bookmark_chans
	bookChannels [LSDJ_BOOKMARK_POSITION_COUNT][LSDJ_CHANNEL_COUNT]byte

	meta *lsdj_song_metadata

	reserved1030 [96]byte
	reserved1fba [70]byte
	reserved2000 [32]byte
	reserved3fbf byte
	reserved3fb9 byte
	reserved3fc6 [10]byte
	reserved3fd1 [47]byte
	reserved5fe0 [32]byte
	reserved7ff2 [13]byte
};

// Create/free projects
func lsdj_song_new() {
	var song Lsdj_song_t

	song.formatVersion = 4
	song.tempo = 128
	song.transposition = 0
	song.drumMax = 0x6C
	for i := 0; i < LSDJ_ROW_COUNT; i++ {
		song.rows[i].Clear()
	}
	for i := 0; i < LSDJ_CHAIN_COUNT; i++ {
		song.chains[i].Clear()
	}
	for i := 0; i < LSDJ_PHRASE_COUNT; i++ {
		song.chains[i].Clear()
	}
}
lsdj_song_t* lsdj_song_copy(const lsdj_song_t* song, lsdj_error_t** error);
void lsdj_song_free(lsdj_song_t* song);

// Deserialize a song
lsdj_song_t* lsdj_song_read(lsdj_vio_t* vio, lsdj_error_t** error);
lsdj_song_t* lsdj_song_read_from_memory(const unsigned char* data, size_t size, lsdj_error_t** error);

// Serialize a song
void lsdj_song_write(const lsdj_song_t* song, lsdj_vio_t* vio, lsdj_error_t** error);
void lsdj_song_write_to_memory(const lsdj_song_t* song, unsigned char* data, size_t size, lsdj_error_t** error);

// Change data in a song
void lsdj_song_set_format_version(lsdj_song_t* song, unsigned char version);
unsigned char lsdj_song_get_format_version(const lsdj_song_t* song);
void lsdj_song_set_tempo(lsdj_song_t* song, unsigned char tempo);
unsigned char lsdj_song_get_tempo(const lsdj_song_t* song);
void lsdj_song_set_transposition(lsdj_song_t* song, unsigned char transposition);
unsigned char lsdj_song_get_transposition(const lsdj_song_t* song);
unsigned char lsdj_song_get_file_changed_flag(const lsdj_song_t* song);
void lsdj_song_set_drum_max(lsdj_song_t* song, unsigned char drumMax);
unsigned char lsdj_song_get_drum_max(const lsdj_song_t* song);

lsdj_row_t* lsdj_song_get_row(lsdj_song_t* song, size_t index);
lsdj_chain_t* lsdj_song_get_chain(lsdj_song_t* song, size_t index);
lsdj_phrase_t* lsdj_song_get_phrase(lsdj_song_t* song, size_t index);
lsdj_instrument_t* lsdj_song_get_instrument(lsdj_song_t* song, size_t index);
lsdj_synth_t* lsdj_song_get_synth(lsdj_song_t* song, size_t index);
lsdj_wave_t* lsdj_song_get_wave(lsdj_song_t* song, size_t index);
lsdj_table_t* lsdj_song_get_table(lsdj_song_t* song, size_t index);
lsdj_groove_t* lsdj_song_get_groove(lsdj_song_t* song, size_t index);
lsdj_word_t* lsdj_song_get_word(lsdj_song_t* song, size_t index);
void lsdj_song_set_word_name(lsdj_song_t* song, size_t index, const char* data, size_t size);
void lsdj_song_get_word_name(lsdj_song_t* song, size_t index, char* data, size_t size);
void lsdj_song_set_bookmark(lsdj_song_t* song, lsdj_channel_t channel, size_t position, unsigned char bookmark);
unsigned char lsdj_song_get_bookmark(lsdj_song_t* song, lsdj_channel_t channel, size_t position);
