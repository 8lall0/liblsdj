package song

import (
	"github.com/8lall0/liblsdj/chain"
	"github.com/8lall0/liblsdj/channel"
	"github.com/8lall0/liblsdj/groove"
	"github.com/8lall0/liblsdj/instrument"
	"github.com/8lall0/liblsdj/phrase"
	"github.com/8lall0/liblsdj/row"
	"github.com/8lall0/liblsdj/synth"
	"github.com/8lall0/liblsdj/table"
	"github.com/8lall0/liblsdj/wave"
	"github.com/8lall0/liblsdj/word"
)

// An LSDJ song
type Song struct {
	formatVersion byte
	tempo         byte
	transposition byte
	drumMax       byte
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
	tables [LSDJ_TABLE_COUNT]*table.Table
	// The grooves in the song
	grooves [LSDJ_GROOVE_COUNT]*groove.Groove
	// The speech synth words in the song
	words     [LSDJ_WORD_COUNT]*word.Word
	wordNames [LSDJ_WORD_COUNT][word.LSDJ_WORD_NAME_LENGTH]byte

	// Bookmarks
	bookmarks struct {
		pulse1 [LSDJ_BOOKMARK_POSITION_COUNT]byte
		pulse2 [LSDJ_BOOKMARK_POSITION_COUNT]byte
		wave   [LSDJ_BOOKMARK_POSITION_COUNT]byte
		noise  [LSDJ_BOOKMARK_POSITION_COUNT]byte
	}
	bookChannels [LSDJ_BOOKMARK_POSITION_COUNT][channel.LSDJ_CHANNEL_COUNT]byte

	metadata struct {
		keyDelay        byte
		keyRepeat       byte
		font            byte
		sync            byte
		colorSet        byte
		clone           byte
		fileChangedFlag byte
		powerSave       byte
		preListen       byte
		workTime        struct {
			hours   byte
			minutes byte
		}
		totalTime struct {
			days    byte
			hours   byte
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
}

func (song *Song) Clear() {
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
	for i := 0; i < LSDJ_SYNTH_COUNT; i++ {
		song.synths[i].Clear()
	}
	for i := 0; i < LSDJ_WAVE_COUNT; i++ {
		song.waves[i].Clear()
	}
	for i := 0; i < LSDJ_TABLE_COUNT; i++ {
		song.tables[i].Clear()
	}
	for i := 0; i < LSDJ_GROOVE_COUNT; i++ {
		song.grooves[i].Clear()
	}
	for i := 0; i < LSDJ_WORD_COUNT; i++ {
		song.words[i].Clear()
	}
	song.wordNames = DEFAULT_WORD_NAMES

	song.metadata.keyDelay = 7
	song.metadata.keyRepeat = 2
	song.metadata.font = 0
	song.metadata.sync = 0
	song.metadata.colorSet = 0
	song.metadata.clone = 0
	song.metadata.powerSave = 0
	song.metadata.preListen = 1

	/*
		TODO: do phrases and instruments
		TODO: bookmarks
	*/
}

func (song *Song) Copy() *Song {
	return &(*song)
}

/*
	TODO: All below
	TODO: finish VIO and SAV
*/
func (song *Song) readBank0() {
	for i := 0; byte(i) < LSDJ_PHRASE_COUNT; i++ {
		if song.phrases[i] != nil {

		}
	}
}
func (song *Song) writeBank0() {

}
func (song *Song) readBank1() {

}
func (song *Song) writeBank1() {

}
func (song *Song) readBank2() {

}
func (song *Song) writeBank2() {

}
func (song *Song) readBank3() {

}
func (song *Song) writeBank3() {

}
func (song *Song) checkRB() {

}

func (song *Song) readSoftSynthParam() {

}

/*
	Public
*/

func (song *Song) Read() {

}
func (song *Song) ReadFromMemory() {

}

func (song *Song) Write() {

}
func (song *Song) WriteToMemory() {

}

func (song *Song) GetFormatVersion() byte {
	return song.formatVersion
}
func (song *Song) SetFormatVersion(formatVersion byte) {
	song.formatVersion = formatVersion
}

func (song *Song) GetTempo() byte {
	return song.tempo
}
func (song *Song) SetTempo(tempo byte) {
	song.tempo = tempo
}

func (song *Song) GetTransposition() byte {
	return song.transposition
}
func (song *Song) SetTransposition(transposition byte) {
	song.transposition = transposition
}

func (song *Song) GetFileChangedFlag() byte {
	return song.metadata.fileChangedFlag
}

func (song *Song) GetDrumMax(drumMax byte) {
	song.drumMax = drumMax
}
func (song *Song) SetDrumMax() byte {
	return song.drumMax
}

func (song *Song) GetRow(index int) *row.Row {
	return song.rows[index]
}
func (song *Song) GetChain(index int) *chain.Chain {
	return song.chains[index]
}
func (song *Song) GetPhrase(index int) *phrase.Phrase {
	return song.phrases[index]
}
func (song *Song) GetInstrument(index int) *instrument.Instrument {
	return song.instruments[index]
}
func (song *Song) GetSynth(index int) *synth.Synth {
	return song.synths[index]
}
func (song *Song) GetWave(index int) *wave.Wave {
	return song.waves[index]
}
func (song *Song) GetTable(index int) *table.Table {
	return song.tables[index]
}
func (song *Song) GetGroove(index int) *groove.Groove {
	return song.grooves[index]
}
func (song *Song) GetWord(index int) *word.Word {
	return song.words[index]
}

/*
BOH!
*/

func (song *Song) SetWordName() {

}
func (song *Song) GetWordName() {

}

func (song *Song) GetBookmark() {

}
func (song *Song) SetBookMark() {}
