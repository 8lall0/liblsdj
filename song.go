package liblsdj

const (
	LSDJ_SONG_DECOMPRESSED_SIZE byte = 0x8000 
	LSDJ_ROW_COUNT byte = 256 
	LSDJ_CHAIN_COUNT byte = 128 
	LSDJ_PHRASE_COUNT byte = 0xFF 
	LSDJ_INSTRUMENT_COUNT byte = 64 
	LSDJ_SYNTH_COUNT byte = 16 
	LSDJ_TABLE_COUNT byte = 32 
	LSDJ_WAVE_COUNT byte = 256 
	LSDJ_GROOVE_COUNT byte = 32 
	LSDJ_WORD_COUNT byte = 42 
	LSDJ_BOOKMARK_POSITION_COUNT byte = 16 
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


// An LSDJ song
struct lsdj_song_t
{
	unsigned char formatVersion;
	unsigned char tempo;
	unsigned char transposition;
	unsigned char drumMax;
	// The sequences of chains in the song
	lsdj_row_t rows[LSDJ_ROW_COUNT];
	// The chains in the song
	lsdj_chain_t* chains[LSDJ_CHAIN_COUNT];
	// The prases in the song
	lsdj_phrase_t* phrases[LSDJ_PHRASE_COUNT];
	// Instruments of the song
	lsdj_instrument_t* instruments[LSDJ_INSTRUMENT_COUNT];
	// Soft synths of the song
	lsdj_synth_t synths[LSDJ_SYNTH_COUNT];
	// Wave frames of the song
	lsdj_wave_t waves[LSDJ_WAVE_COUNT];
	// The tables in the song
	lsdj_table_t* tables[LSDJ_TABLE_COUNT];
	// The grooves in the song
	lsdj_groove_t grooves[LSDJ_GROOVE_COUNT];
	// The speech synth words in the song
	lsdj_word_t words[LSDJ_WORD_COUNT];
	char wordNames[LSDJ_WORD_COUNT][LSDJ_WORD_NAME_LENGTH];

	// Bookmarks
	union
	{
		struct
		{
			unsigned char pulse1[LSDJ_BOOKMARK_POSITION_COUNT];
			unsigned char pulse2[LSDJ_BOOKMARK_POSITION_COUNT];
			unsigned char wave[LSDJ_BOOKMARK_POSITION_COUNT];
			unsigned char noise[LSDJ_BOOKMARK_POSITION_COUNT];
		};
		unsigned char channels[LSDJ_BOOKMARK_POSITION_COUNT][LSDJ_CHANNEL_COUNT];
	} bookmarks;

	struct
	{
		unsigned char keyDelay;
		unsigned char keyRepeat;
		unsigned char font;
		unsigned char sync;
		unsigned char colorSet;
		unsigned char clone;
		unsigned char fileChangedFlag;
		unsigned char powerSave;
		unsigned char preListen;

		struct
		{
			unsigned char days;
			unsigned char hours;
			unsigned char minutes;
		} totalTime;

		struct
		{
		unsigned char hours;
		unsigned char minutes;
		} workTime;
	} meta;

	unsigned char reserved1030[96];
	unsigned char reserved1fba[70];
	unsigned char reserved2000[32];
	unsigned char reserved3fbf;
	unsigned char reserved3fb9;
	unsigned char reserved3fc6[10];
	unsigned char reserved3fd1[47];
	unsigned char reserved5fe0[32];
	unsigned char reserved7ff2[13];
};

// Create/free projects
lsdj_song_t* lsdj_song_new(lsdj_error_t** error);
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
