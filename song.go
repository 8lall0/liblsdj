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

func (s *song) Clear() {
	s.formatVersion = 4
	s.tempo = 128
	s.transposition = 0
	s.drumMax = 0x6C

	s.rows = make([]*row, lsdj_ROW_COUNT)
	s.chains = make([]*chain, lsdj_CHAIN_COUNT)
	s.synths = make([]*synth, lsdj_SYNTH_COUNT)
	s.waves = make([]*wave, lsdj_WAVE_COUNT)
	s.tables = make([]*table, lsdj_TABLE_COUNT)
	s.grooves = make([]*groove, lsdj_GROOVE_COUNT)
	s.words = make([]*word, lsdj_WORD_COUNT)

	copy(s.wordNames, DEFAULT_WORD_NAMES)

	s.metadata.keyDelay = 7
	s.metadata.keyRepeat = 2
	s.metadata.font = 0
	s.metadata.sync = 0
	s.metadata.colorSet = 0
	s.metadata.clone = 0
	s.metadata.powerSave = 0
	s.metadata.preListen = 1
}
func (s *song) Copy() *song {
	return &(*s)
}
func (s *song) readBank0(r *vio) {
	for i := 0; i < lsdj_PHRASE_COUNT; i++ {
		if s.phrases[i] != nil {
			s.phrases[i].notes = r.read(lsdj_PHRASE_LENGTH)
		} else {
			r.seek(r.getCur() + lsdj_PHRASE_LENGTH)
		}
	}
	// TODO: Ask for this! How to handle union???
	// ALERT: ignoring bookChannels
	s.bookmarks.pulse1 = r.read(lsdj_BOOKMARK_POSITION_COUNT)
	s.bookmarks.pulse2 = r.read(lsdj_BOOKMARK_POSITION_COUNT)
	s.bookmarks.wave = r.read(lsdj_BOOKMARK_POSITION_COUNT)
	s.bookmarks.noise = r.read(lsdj_BOOKMARK_POSITION_COUNT)

	s.reserved1030 = r.read(reserved_1030)

	for i := 0; i < lsdj_GROOVE_COUNT; i++ {
		s.grooves[i].groove = r.read(lsdj_GROOVE_LENGTH)
	}
	// ALERT: ignoring channels
	for i := 0; i < lsdj_ROW_COUNT; i++ {
		s.rows[i].channelList.pulse1 = r.readSingle()
		s.rows[i].channelList.pulse2 = r.readSingle()
		s.rows[i].channelList.wave = r.readSingle()
		s.rows[i].channelList.noise = r.readSingle()
		//s.rows[i].channels = r.read(lsdj_CHANNEL_COUNT)
	}

	for i := 0; i < lsdj_TABLE_COUNT; i++ {
		if s.tables[i] != nil {
			s.tables[i].volumes = r.read(lsdj_TABLE_LENGTH)
		} else {
			r.seek(r.getCur() + lsdj_TABLE_LENGTH)
		}
	}

	for i := 0; i < lsdj_WORD_COUNT; i++ {
		s.words[i].allophones = r.read(lsdj_WORD_LENGTH)
		s.words[i].lengths = r.read(lsdj_WORD_LENGTH)
	}

	for i := 0; i < lsdj_WORD_COUNT; i++ {
		s.wordNames[i] = r.read(lsdj_WORD_NAME_LENGTH)
	}
	// jumping RB
	r.seek(r.getCur() + 2)

	for i := 0; i < lsdj_INSTRUMENT_COUNT; i++ {
		if s.instruments[i] != nil {
			s.instruments[i].name = r.read(lsdj_INSTRUMENT_NAME_LENGTH)
		} else {
			r.seek(r.getCur() + lsdj_INSTRUMENT_NAME_LENGTH)
		}
	}

	s.reserved1fba = r.read(reserved_1fba)
}
func (song *song) writeBank0() {

}
func (s *song) readBank1(r *vio) {
	s.reserved2000 = r.read(reserved_2000)

	// table and instr alloc tables already read at beginning
	r.seek(r.getCur() + lsdj_TABLE_ALLOC_TABLE_SIZE + lsdj_INSTR_ALLOC_TABLE_SIZE)

	for i := 0; i < lsdj_CHAIN_COUNT; i++ {
		if s.chains[i] != nil {
			s.chains[i].transpositions = r.read(lsdj_CHAIN_LENGTH)
		} else {
			r.seek(r.getCur() + lsdj_CHAIN_LENGTH)
		}
	}
	for i := 0; i < lsdj_INSTRUMENT_COUNT; i++ {
		if s.instruments[i] != nil {
			// Qua mi serve instrument_read
		} else {
			r.seek(r.getCur() + lsdj_INSTRUMENT_COUNT)
		}
	}
	for i := 0; i < lsdj_TABLE_COUNT; i++ {
		if s.tables[i] != nil {
			// table get command1
		} else {
			r.seek(r.getCur() + lsdj_TABLE_COUNT)
		}
	}
	// ALERT: duplicates in the original code!!!
	for i := 0; i < lsdj_TABLE_COUNT; i++ {
		if s.tables[i] != nil {
			// table get command2
		} else {
			r.seek(r.getCur() + lsdj_TABLE_COUNT)
		}
	}
	for i := 0; i < lsdj_TABLE_COUNT; i++ {
		if s.tables[i] != nil {
			// table get command2
		} else {
			r.seek(r.getCur() + lsdj_TABLE_COUNT)
		}
	}

	// jumping RB
	r.seek(r.getCur() + 2)
	// Already read at the beginning
	r.seek(r.getCur() + lsdj_PHRASE_ALLOC_TABLE_SIZE + lsdj_CHAIN_ALLOC_TABLE_SIZE)

	for i := 0; i < lsdj_SYNTH_COUNT; i++ {
		s.synths[i].readSoftSynthParam(r)
	}

	s.metadata.workTime.hours = r.readSingle()
	s.metadata.workTime.minutes = r.readSingle()
	s.tempo = r.readSingle()
	s.transposition = r.readSingle()
	s.metadata.totalTime.days = r.readSingle()
	s.metadata.totalTime.hours = r.readSingle()
	s.metadata.totalTime.minutes = r.readSingle()
	s.reserved3fb9 = r.readSingle()
	s.metadata.keyDelay = r.readSingle()
	s.metadata.keyRepeat = r.readSingle()
	s.metadata.font = r.readSingle()
	s.metadata.sync = r.readSingle()
	s.metadata.colorSet = r.readSingle()
	s.reserved3fbf = r.readSingle()
	s.metadata.clone = r.readSingle()
	s.metadata.fileChangedFlag = r.readSingle()
	s.metadata.powerSave = r.readSingle()
	s.metadata.preListen = r.readSingle()

	var waveSynthOverwriteLocks []byte // 2
	waveSynthOverwriteLocks = r.read(2)
	var i uint8
	for i = 0; i < uint8(lsdj_SYNTH_COUNT); i++ {
		s.synths[i].overwritten = (waveSynthOverwriteLocks[1-(i/8)] >> (i % 8)) & 1
	}
	s.reserved3fc6 = r.read(reserved_3fc6)
	s.drumMax = r.readSingle()
	s.reserved3fd1 = r.read(reserved_3fd1)
}
func (song *song) writeBank1() {

}
func (s *song) readBank2(r *vio) {
	for i := 0; i < lsdj_PHRASE_COUNT; i++ {
		if s.phrases[i] != nil {
			for j := 0; j < lsdj_PHRASE_LENGTH; j++ {
				s.phrases[i].commands[j].command = r.readSingle()
			}
		} else {
			r.seek(r.getCur() + lsdj_PHRASE_LENGTH)
		}
	}
	for i := 0; i < lsdj_PHRASE_COUNT; i++ {
		if s.phrases[i] != nil {
			for j := 0; j < lsdj_PHRASE_LENGTH; j++ {
				s.phrases[i].commands[j].value = r.readSingle()
			}
		} else {
			r.seek(r.getCur() + lsdj_PHRASE_LENGTH)
		}
	}
	s.reserved5fe0 = r.read(reserved_5fe0)
}
func (song *song) writeBank2() {

}
func (s *song) readBank3(r *vio) {
	for i := 0; i < lsdj_WAVE_COUNT; i++ {
		s.waves[i].data = r.read(lsdj_WAVE_LENGTH)
	}

	for i := 0; i < lsdj_PHRASE_COUNT; i++ {
		if s.phrases[i] != nil {
			s.phrases[i].instruments = r.read(lsdj_PHRASE_LENGTH)
		} else {
			r.seek(r.getCur() + lsdj_PHRASE_LENGTH)
		}
	}
	// RB
	r.seek(r.getCur() + 2)
	s.reserved7ff2 = r.read(reserved_7ff2)
	// Version number already read
	r.seek(r.getCur() + 1)
}
func (song *song) writeBank3() {

}
func checkRB(r *vio, i int) {
	r.seek(i)
	fmt.Println(string(r.readSingle()), string(r.readSingle()))
}

/*
	Public
*/

func (s *song) Read(r *vio) {
	var instrAllocTable []byte
	var tableAllocTable []byte
	var chainAllocTable []byte
	var phraseAllocTable []byte

	fmt.Println("Check RB...")
	checkRB(r, 7800)
	checkRB(r, 16000)
	checkRB(r, 32752)
	//Everything is correct, so i initialize the s.
	s.Clear()
	r.seek(int(0x7fff))
	s.formatVersion = r.readSingle()

	r.seek(int(0x2020))
	tableAllocTable = r.read(lsdj_TABLE_ALLOC_TABLE_SIZE)
	instrAllocTable = r.read(lsdj_INSTR_ALLOC_TABLE_SIZE)

	r.seek(int(0x3E82))
	phraseAllocTable = r.read(lsdj_PHRASE_ALLOC_TABLE_SIZE)
	chainAllocTable = r.read(lsdj_CHAIN_ALLOC_TABLE_SIZE)

	for i := 0; i < lsdj_TABLE_ALLOC_TABLE_SIZE; i++ {
		if tableAllocTable[i] != 0 {
			s.tables[i] = new(table)
		} else {
			s.tables[i] = nil
		}
	}
	/*
		check instrument reset!!!
	*/
	for i := 0; i < lsdj_INSTRUMENT_COUNT; i++ {
		if instrAllocTable[i] != 0 {
			s.instruments[i] = new(instrument)
		} else {
			s.instruments[i] = nil
		}
	}
	var i uint8
	for i = 0; i < uint8(lsdj_CHAIN_COUNT); i++ {
		if (chainAllocTable[i/8]>>(i%8))&1 == 1 {
			s.chains[i] = new(chain)
		} else {
			s.chains[i] = nil
		}
	}
	for i = 0; i < uint8(lsdj_PHRASE_COUNT); i++ {
		if (phraseAllocTable[i/8]>>(i%8))&1 == 1 {
			s.phrases[i] = new(phrase)
		} else {
			s.phrases[i] = nil
		}
	}

	r.seek(0)
	s.readBank0(r)
	s.readBank1(r)
	s.readBank2(r)
	s.readBank3(r)
}

func (song *song) Write() {

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

func (song *song) SetWordName() {}
func (song *song) GetWordName() {}
func (song *song) GetBookmark() {}
func (song *song) SetBookMark() {}
