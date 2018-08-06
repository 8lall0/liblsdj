package song

import (
	"github.com/8lall0/liblsdj/row"
	"github.com/8lall0/liblsdj/chain"
	"github.com/8lall0/liblsdj/phrase"
	"github.com/8lall0/liblsdj/instrument"
	"github.com/8lall0/liblsdj/synth"
	"github.com/8lall0/liblsdj/wave"
	"github.com/8lall0/liblsdj/table"
	"github.com/8lall0/liblsdj/groove"
	"github.com/8lall0/liblsdj/word"
	"github.com/8lall0/liblsdj/channel"
)

// An LSDJ song
type Song struct {
	formatVersion byte
	tempo byte
	transposition byte
	drumMax byte
	// The sequences of chains in the song
	rows [LSDJ_ROW_COUNT]*row.Row
	// The chains in the song
	chains [LSDJ_CHAIN_COUNT]*chain.Chain
	// The prases in the song
	phrases [LSDJ_PHRASE_COUNT]*phrase.Phrase
	// Instruments of the song
	instruments [LSDJ_INSTRUMENT_COUNT]*instrument.Instrument
	// Soft synths of the song
	synths [LSDJ_SYNTH_COUNT]*synth.Synth
	// Wave frames of the song
	waves [LSDJ_WAVE_COUNT]*wave.Wave
	// The tables in the song
	tables[LSDJ_TABLE_COUNT]*table.Table
	// The grooves in the song
	grooves [LSDJ_GROOVE_COUNT]*groove.Groove
	// The speech synth words in the song
	words [LSDJ_WORD_COUNT]*word.Word
	wordNames [LSDJ_WORD_COUNT][word.LSDJ_WORD_NAME_LENGTH]string

	// Bookmarks
	bookmarks struct {
		pulse1 [LSDJ_BOOKMARK_POSITION_COUNT]byte
		pulse2 [LSDJ_BOOKMARK_POSITION_COUNT]byte
		wave [LSDJ_BOOKMARK_POSITION_COUNT]byte
		noise [LSDJ_BOOKMARK_POSITION_COUNT]byte
	}
	bookChannels [LSDJ_BOOKMARK_POSITION_COUNT][channel.LSDJ_CHANNEL_COUNT]byte

	metadata  struct {
		keyDelay byte
		keyRepeat byte
		font byte
		sync byte
		colorSet byte
		clone byte
		fileChangedFlag byte
		powerSave byte
		preListen byte
		workTime struct {
			hours byte
			minutes byte
		}
		totalTime struct {
			days byte
			hours byte
			minutes byte
		}
	}

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
	var song Song

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
