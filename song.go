package liblsdj

import "fmt"

const (
	lsdj_SONG_DECOMPRESSED_SIZE  int  = int(0x8000)
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

var lsdj_CHAIN_LENGTH_ZERO = []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
var lsdj_CHAIN_LENGTH_FF = []byte{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF}

// An LSDJ song
type song struct {
	formatVersion byte
	tempo         byte
	transposition byte
	drumMax       byte
	// The sequences of chains in the song
	rows rowA
	// The chains in the song
	chains chainA
	// The prases in the song
	phrases phraseA
	// Instruments of the song
	instruments instrumentContainerA
	// Soft synths of the song
	synths synthA
	// wave frames of the song
	waves waveA
	// The tables in the song
	tables tableA
	// The grooves in the song
	grooves grooveA
	// The speech synth words in the song
	words     wordA
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

	s.rows.initialize()
	s.synths.initialize()
	s.waves.initialize()
	s.grooves.initialize()
	s.words.initialize()

	s.wordNames = make([][]byte, lsdj_WORD_COUNT)
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
	s.phrases.readNote(r)

	s.bookmarks.pulse1 = r.read(lsdj_BOOKMARK_POSITION_COUNT)
	s.bookmarks.pulse2 = r.read(lsdj_BOOKMARK_POSITION_COUNT)
	s.bookmarks.wave = r.read(lsdj_BOOKMARK_POSITION_COUNT)
	s.bookmarks.noise = r.read(lsdj_BOOKMARK_POSITION_COUNT)

	s.reserved1030 = r.read(reserved_1030)

	s.grooves.readGroove(r)
	s.rows.readRow(r)
	s.tables.readVolume(r)
	s.words.readWord(r)

	for i := 0; i < lsdj_WORD_COUNT; i++ {
		s.wordNames[i] = r.read(lsdj_WORD_NAME_LENGTH)
	}

	// jumping RB
	r.seekCur(2)

	s.instruments.readInsName(r)

	s.reserved1fba = r.read(reserved_1fba)
}
func (s *song) writeBank0(w *vio) {
	s.phrases.writeNote(w)

	w.write(s.bookmarks.pulse1)
	w.write(s.bookmarks.pulse2)
	w.write(s.bookmarks.wave)
	w.write(s.bookmarks.noise)

	w.write(s.reserved1030)

	s.grooves.writeGroove(w)
	s.rows.writeRow(w)
	s.words.writeWord(w)
	for i := range s.wordNames {
		w.write(s.wordNames[i])
	}
	w.write([]byte("rb"))
	s.instruments.writeInsName(w)
	w.write(s.reserved1fba)
}

func (s *song) readBank1(r *vio) {
	s.reserved2000 = r.read(reserved_2000)

	// table and instr alloc tables already setInstrument at beginning
	r.seekCur(lsdj_TABLE_ALLOC_TABLE_SIZE + lsdj_INSTR_ALLOC_TABLE_SIZE)

	s.chains.readChain(r)
	s.instruments.readInstrument(r, s.formatVersion)

	s.tables.readCommand1(r)
	s.tables.readCommand2(r)
	s.tables.readCommand2(r)

	// jumping RB
	r.seekCur(2)
	// Already setInstrument at the beginning
	r.seekCur(lsdj_PHRASE_ALLOC_TABLE_SIZE + lsdj_CHAIN_ALLOC_TABLE_SIZE)

	s.synths.readSynthParam(r)

	s.metadata.workTime.hours = r.readByte()
	s.metadata.workTime.minutes = r.readByte()
	s.tempo = r.readByte()
	s.transposition = r.readByte()
	s.metadata.totalTime.days = r.readByte()
	s.metadata.totalTime.hours = r.readByte()
	s.metadata.totalTime.minutes = r.readByte()

	s.reserved3fb9 = r.readByte()

	s.metadata.keyDelay = r.readByte()
	s.metadata.keyRepeat = r.readByte()
	s.metadata.font = r.readByte()
	s.metadata.sync = r.readByte()
	s.metadata.colorSet = r.readByte()

	s.reserved3fbf = r.readByte()

	s.metadata.clone = r.readByte()
	s.metadata.fileChangedFlag = r.readByte()
	s.metadata.powerSave = r.readByte()
	s.metadata.preListen = r.readByte()

	s.synths.readSynthOverwritten(r)

	s.reserved3fc6 = r.read(reserved_3fc6)

	s.drumMax = r.readByte()

	s.reserved3fd1 = r.read(reserved_3fd1)
}
func (s *song) writeBank1(w *vio) {
	w.write(s.reserved2000)
	s.tables.writeTabAllocTable(w)
	s.instruments.writeInsAllocTable(w)
	s.chains.writeChain(w)
	s.instruments.writeInstrument(w, s.formatVersion)
	s.tables.writeTable(w, s.formatVersion)

	w.write([]byte("rb"))

	s.phrases.writePhraseAllocTable(w)
	s.chains.writeChainAllocTable(w)
	s.synths.writeSynthParam(w)

	w.writeByte(s.metadata.workTime.hours)
	w.writeByte(s.metadata.workTime.minutes)

	w.writeByte(s.tempo)
	w.writeByte(s.transposition)

	w.writeByte(s.metadata.totalTime.days)
	w.writeByte(s.metadata.totalTime.hours)
	w.writeByte(s.metadata.totalTime.minutes)

	w.writeByte(s.reserved3fb9)

	w.writeByte(s.metadata.keyDelay)
	w.writeByte(s.metadata.keyRepeat)
	w.writeByte(s.metadata.font)
	w.writeByte(s.metadata.sync)
	w.writeByte(s.metadata.colorSet)
	w.writeByte(s.reserved3fbf)
	w.writeByte(s.metadata.clone)
	w.writeByte(s.metadata.fileChangedFlag)
	w.writeByte(s.metadata.powerSave)
	w.writeByte(s.metadata.preListen)

	s.synths.writeSynthOverwritten(w)

	w.write(s.reserved3fc6)
	w.writeByte(s.drumMax)
	w.write(s.reserved3fd1)
}

func (s *song) readBank2(r *vio) {
	s.phrases.readCommand(r)

	s.reserved5fe0 = r.read(reserved_5fe0)
}
func (song *song) writeBank2() {

}
func (s *song) readBank3(r *vio) {
	s.waves.readWave(r)
	s.phrases.readInstrument(r)

	// RB
	r.seekCur(2)
	s.reserved7ff2 = r.read(reserved_7ff2)

	// Version number already setInstrument
	r.seekCur(1)
}
func (song *song) writeBank3() {

}
func checkRB(r *vio, i int) {
	r.seek(i)
	fmt.Println(string(r.readByte()), string(r.readByte()))
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
	//Everything is correct, so instrument initialize the s.
	s.Clear()
	r.seek(int(0x7fff))
	s.formatVersion = r.readByte()

	r.seek(int(0x2020))
	tableAllocTable = r.read(lsdj_TABLE_ALLOC_TABLE_SIZE)
	instrAllocTable = r.read(lsdj_INSTR_ALLOC_TABLE_SIZE)

	r.seek(int(0x3E82))
	phraseAllocTable = r.read(lsdj_PHRASE_ALLOC_TABLE_SIZE)
	chainAllocTable = r.read(lsdj_CHAIN_ALLOC_TABLE_SIZE)

	s.tables.initialize(tableAllocTable)
	s.instruments.initialize(instrAllocTable)
	s.chains.initialize(chainAllocTable)
	s.phrases.initialize(phraseAllocTable)

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

/*func (song *song) GetPhrase(index int) *phrase {
	return song.phrases[index]
}*/
/*func (song *song) GetInstrument(index int) *instrument {
	return song.instruments[index]
}*/
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
