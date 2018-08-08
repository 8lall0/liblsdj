package liblsdj

import "fmt"

const (
	lsdj_SONG_DECOMPRESSED_SIZE  int  = 32768 //0x8000
	lsdj_ROW_COUNT               int  = 256
	lsdj_CHAIN_COUNT             int  = 128
	lsdj_PHRASE_COUNT            int  = int(0xFF)
	lsdj_INSTRUMENT_COUNT        int  = 64
	lsdj_SYNTH_COUNT             int  = 16
	lsdj_TABLE_COUNT             int  = 32
	lsdj_WAVE_COUNT              int  = 256
	lsdj_GROOVE_COUNT            int  = 32
	lsdj_WORD_COUNT              int  = 42
	lsdj_BOOKMARK_POSITION_COUNT int  = 16
	lsdj_NO_BOOKMARK             byte = 0xFF

	lsdj_CLONE_DEEP byte = 0
	lsdj_CLONE_SLIM byte = 1

	lsdj_INSTR_ALLOC_TABLE_SIZE  int = 64
	lsdj_TABLE_ALLOC_TABLE_SIZE  int = 32
	lsdj_CHAIN_ALLOC_TABLE_SIZE  int = 16
	lsdj_PHRASE_ALLOC_TABLE_SIZE int = 32
)

const (
	reserved_1030 int = 96
	reserved_1fba int = 70
	reserved_2000 int = 32
	reserved_3fc6 int = 10
	reserved_3fd1 int = 47
	reserved_5fe0 int = 32
	reserved_7ff2 int = 13
)

//[lsdj_WORD_COUNT][lsdj_WORD_NAME_LENGTH]
var DEFAULT_WORD_NAMES = [][]byte{
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

const ()

var lsdj_TABLE_LENGTH_ZERO = [lsdj_TABLE_LENGTH]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
var lsdj_CHAIN_LENGTH_ZERO = [lsdj_TABLE_LENGTH]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
var lsdj_CHAIN_LENGTH_FF = [lsdj_TABLE_LENGTH]byte{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF}
var lsdj_PHRASE_LENGTH_ZERO = [lsdj_PHRASE_LENGTH]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
var lsdj_PHRASE_LENGTH_FF = [lsdj_PHRASE_LENGTH]byte{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF}

// An LSDJ song
type song struct {
	formatVersion byte
	tempo         byte
	transposition byte
	drumMax       byte
	// The sequences of chains in the song
	rows []*row //lsdj_ROW_COUNT
	// The chains in the song
	chains []*chain //lsdj_CHAIN_COUNT
	// The prases in the song
	phrases []*phrase //lsdj_PHRASE_COUNT
	// Instruments of the song
	instruments []*instrument //lsdj_INSTRUMENT_COUNT
	// Soft synths of the song
	synths []*synth //lsdj_SYNTH_COUNT
	// wave frames of the song
	waves []*wave //lsdj_WAVE_COUNT
	// The tables in the song
	tables []*table //lsdj_TABLE_COUNT
	// The grooves in the song
	grooves []*groove //lsdj_GROOVE_COUNT
	// The speech synth words in the song
	words     []*word  //lsdj_WORD_COUNT
	wordNames [][]byte //[lsdj_WORD_COUNT][lsdj_WORD_NAME_LENGTH]

	// Bookmarks
	bookmarks struct {
		pulse1 []byte //lsdj_BOOKMARK_POSITION_COUNT
		pulse2 []byte //lsdj_BOOKMARK_POSITION_COUNT
		wave   []byte //lsdj_BOOKMARK_POSITION_COUNT
		noise  []byte //lsdj_BOOKMARK_POSITION_COUNT
	}
	bookChannels [lsdj_BOOKMARK_POSITION_COUNT][lsdj_CHANNEL_COUNT]byte

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

	reserved1030 []byte
	reserved1fba []byte
	reserved2000 []byte
	reserved3fbf byte
	reserved3fb9 byte
	reserved3fc6 []byte
	reserved3fd1 []byte
	reserved5fe0 []byte
	reserved7ff2 []byte
}

func (song *song) Clear() {
	song.formatVersion = 4
	song.tempo = 128
	song.transposition = 0
	song.drumMax = 0x6C

	song.rows = make([]*row, lsdj_ROW_COUNT)
	song.chains = make([]*chain, lsdj_CHAIN_COUNT)
	song.synths = make([]*synth, lsdj_SYNTH_COUNT)
	song.waves = make([]*wave, lsdj_WAVE_COUNT)
	song.tables = make([]*table, lsdj_TABLE_COUNT)
	song.grooves = make([]*groove, lsdj_GROOVE_COUNT)
	song.words = make([]*word, lsdj_WORD_COUNT)

	copy(song.wordNames, DEFAULT_WORD_NAMES)

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
func (song *song) Copy() *song {
	return &(*song)
}
func (song *song) readBank0(r *vio) {
	for i := 0; i < lsdj_PHRASE_COUNT; i++ {
		if song.phrases[i] != nil {
			song.phrases[i].notes = r.read(lsdj_PHRASE_LENGTH)
		} else {
			r.seek(r.getCur() + lsdj_PHRASE_LENGTH)
		}
	}
	song.bookmarks.pulse1 = r.read(lsdj_BOOKMARK_POSITION_COUNT)
	song.bookmarks.pulse2 = r.read(lsdj_BOOKMARK_POSITION_COUNT)
	song.bookmarks.wave = r.read(lsdj_BOOKMARK_POSITION_COUNT)
	song.bookmarks.noise = r.read(lsdj_BOOKMARK_POSITION_COUNT)

	/* Ask for this! How to handle union??? */
	for i := 0; i < lsdj_BOOKMARK_POSITION_COUNT; i++ {
		song.bookmarks.pulse1 = r.read(lsdj_BOOKMARK_POSITION_COUNT)
		song.bookmarks.pulse2 = r.read(lsdj_BOOKMARK_POSITION_COUNT)
		song.bookmarks.wave = r.read(lsdj_BOOKMARK_POSITION_COUNT)
		song.bookmarks.noise = r.read(lsdj_BOOKMARK_POSITION_COUNT)
	}
	song.reserved1030 = r.read(reserved_1030)

	for i := 0; i < lsdj_GROOVE_COUNT; i++ {
		song.grooves[i].groove = r.read(lsdj_GROOVE_LENGTH)
	}
	for i := 0; i < lsdj_ROW_COUNT; i++ {
		song.rows[i].channelList.pulse1 = r.readSingle()
		song.rows[i].channelList.pulse2 = r.readSingle()
		song.rows[i].channelList.wave = r.readSingle()
		song.rows[i].channelList.noise = r.readSingle()
		song.rows[i].channels = r.read(lsdj_CHANNEL_COUNT)
	}

	for i := 0; i < lsdj_TABLE_COUNT; i++ {
		if song.tables[i] != nil {
			song.tables[i].volumes = r.read(lsdj_TABLE_LENGTH)
		} else {
			r.seek(r.getCur() + lsdj_TABLE_LENGTH)
		}
	}

	for i := 0; i < lsdj_WORD_COUNT; i++ {
		song.words[i].allophones = r.read(lsdj_WORD_LENGTH)
		song.words[i].lengths = r.read(lsdj_WORD_LENGTH)
	}

	for i := 0; i < lsdj_WORD_COUNT; i++ {
		song.wordNames[i] = r.read(lsdj_WORD_NAME_LENGTH)
	}
	// jumping RB
	r.seek(r.getCur() + 2)

	for i := 0; i < lsdj_INSTRUMENT_COUNT; i++ {
		if song.instruments[i] != nil {
			song.instruments[i].name = r.read(lsdj_INSTRUMENT_NAME_LENGTH)
		} else {
			r.seek(r.getCur() + lsdj_INSTRUMENT_NAME_LENGTH)
		}
	}

	song.reserved1fba = r.read(reserved_1fba)
}
func (song *song) writeBank0() {

}
func (song *song) readBank1(r *vio) {
	song.reserved2000 = r.read(reserved_2000)

	r.seek(r.getCur() + lsdj_TABLE_ALLOC_TABLE_SIZE + lsdj_INSTR_ALLOC_TABLE_SIZE)

}
func (song *song) writeBank1() {

}
func (song *song) readBank2() {

}
func (song *song) writeBank2() {

}
func (song *song) readBank3() {

}
func (song *song) writeBank3() {

}
func (song *song) checkRB(r *vio, i int) {
	r.seek(i)
	fmt.Println(string(r.readSingle()), string(r.readSingle()))
}
func (song *song) readSoftSynthParam() {
}

/*
	Public
*/

func (song *song) Read(r *vio) {
	var instrAllocTable []byte
	var tableAllocTable []byte
	var chainAllocTable []byte
	var phraseAllocTable []byte

	fmt.Println("Check RB...")
	song.checkRB(r, 7800)
	song.checkRB(r, 16000)
	song.checkRB(r, 32752)
	//Everything is correct, so i initialize the song.
	song.Clear()
	r.seek(int(0x7fff))
	song.formatVersion = r.readSingle()

	r.seek(int(0x2020))
	tableAllocTable = r.read(lsdj_TABLE_ALLOC_TABLE_SIZE)
	instrAllocTable = r.read(lsdj_INSTR_ALLOC_TABLE_SIZE)

	r.seek(int(0x3E82))
	phraseAllocTable = r.read(lsdj_PHRASE_ALLOC_TABLE_SIZE)
	chainAllocTable = r.read(lsdj_CHAIN_ALLOC_TABLE_SIZE)

	// Probabilmente allocazioni, probabilmente non necessaria
	for i := 0; i < lsdj_TABLE_ALLOC_TABLE_SIZE; i++ {
		if tableAllocTable[i] != 0 {
			song.tables[i] = new(table)
		} else {
			song.tables[i] = nil
		}
	}
	/*
		check instrument reset!!!
	*/
	for i := 0; i < lsdj_INSTRUMENT_COUNT; i++ {
		if instrAllocTable[i] != 0 {
			song.instruments[i] = new(instrument)
		} else {
			song.instruments[i] = nil
		}
	}
	// Capire perché porcodio ste robe
	for i := 0; i < lsdj_CHAIN_COUNT; i++ {
		if (chainAllocTable[i/8] >> (i % 8)) & 1 {
			song.chains[i] = new(chain)
		} else {
			song.chains[i] = nil
		}
	}
	for i := 0; i < int(lsdj_PHRASE_COUNT); i++ {
		if (phraseAllocTable[i/8] >> (i % 8)) & 1 {
			song.phrases[i] = new(phrase)
		} else {
			song.phrases[i] = nil
		}
	}

	r.seek(0)
	song.readBank0(r)
}

func (song *song) Write() {

}
func (song *song) WriteToMemory() {

}

func (song *song) GetFormatVersion() byte {
	return song.formatVersion
}
func (song *song) SetFormatVersion(formatVersion byte) {
	song.formatVersion = formatVersion
}

func (song *song) GetTempo() byte {
	return song.tempo
}
func (song *song) SetTempo(tempo byte) {
	song.tempo = tempo
}

func (song *song) GetTransposition() byte {
	return song.transposition
}
func (song *song) SetTransposition(transposition byte) {
	song.transposition = transposition
}

func (song *song) GetFileChangedFlag() byte {
	return song.metadata.fileChangedFlag
}

func (song *song) GetDrumMax(drumMax byte) {
	song.drumMax = drumMax
}
func (song *song) SetDrumMax() byte {
	return song.drumMax
}

func (song *song) GetRow(index int) *row {
	return song.rows[index]
}
func (song *song) GetChain(index int) *chain {
	return song.chains[index]
}
func (song *song) GetPhrase(index int) *phrase {
	return song.phrases[index]
}
func (song *song) GetInstrument(index int) *instrument {
	return song.instruments[index]
}
func (song *song) GetSynth(index int) *synth {
	return song.synths[index]
}
func (song *song) GetWave(index int) *wave {
	return song.waves[index]
}
func (song *song) GetTable(index int) *table {
	return song.tables[index]
}
func (song *song) GetGroove(index int) *groove {
	return song.grooves[index]
}
func (song *song) GetWord(index int) *word {
	return song.words[index]
}

/*
BOH!
*/

func (song *song) SetWordName() {

}
func (song *song) GetWordName() {

}

func (song *song) GetBookmark() {

}
func (song *song) SetBookMark() {}
