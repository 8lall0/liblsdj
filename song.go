package liblsdj

const (
	songDecompressedSize = 0x8000
	rowCnt               = 256
	chainCnt             = 128
	phraseCnt            = 0xFF
	instrCnt             = 64
	synthCnt             = 16
	tableCnt             = 32
	waveCnt              = 256
	grooveCnt            = 32
	wordCnt              = 42
	bookmarkPosCnt       = 16
	noBookmarm           = 0xFF
	cloneDeep            = 0
	cloneSlim            = 1

	instrAllocTableSize  = 64
	tableAllocTableSize  = 32
	chainAllocTableSize  = 16
	phraseAllocTableSize = 32
)

var tableLengthZero = [tableLen]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
var chainLenZero = [tableLen]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
var chainLenFF = [tableLen]byte{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF}
var phraseLenZero = [phraseLen]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
var phraseLenFF = [phraseLen]byte{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF}

var wordNameDefault = [wordCnt][wordNameLen]byte{{'C', ' ', '2', ' '}, {'C', ' ', '2', ' '}, {'D', ' ', '2', ' '}, {'D', ' ', '2', ' '}, {'E', ' ', '2', ' '}, {'F', ' ', '2', ' '}, {'F', ' ', '2', ' '}, {'G', ' ', '2', ' '}, {'G', ' ', '2', ' '}, {'A', ' ', '2', ' '}, {'A', ' ', '2', ' '}, {'B', ' ', '2', ' '}, {'C', ' ', '3', ' '}, {'C', ' ', '3', ' '}, {'D', ' ', '3', ' '}, {'D', ' ', '3', ' '}, {'E', ' ', '3', ' '}, {'F', ' ', '3', ' '}, {'F', ' ', '3', ' '}, {'G', ' ', '3', ' '}, {'G', ' ', '3', ' '}, {'A', ' ', '3', ' '}, {'A', ' ', '3', ' '}, {'B', ' ', '3', ' '}, {'C', ' ', '4', ' '}, {'C', ' ', '4', ' '}, {'D', ' ', '4', ' '}, {'D', ' ', '4', ' '}, {'E', ' ', '4', ' '}, {'F', ' ', '4', ' '}, {'F', ' ', '4', ' '}, {'G', ' ', '4', ' '}, {'G', ' ', '4', ' '}, {'A', ' ', '4', ' '}, {'A', ' ', '4', ' '}, {'B', ' ', '4', ' '}, {'C', ' ', '5', ' '}, {'C', ' ', '5', ' '}, {'D', ' ', '5', ' '}, {'D', ' ', '5', ' '}, {'E', ' ', '5', ' '}, {'F', ' ', '5', ' '}}

type Song struct {
	formatVersion byte
	tempo         byte
	transposition byte
	drumMax       byte

	rows        [rowCnt]*row
	chains      [chainCnt]*chain
	phrases     [phraseCnt]*phrase
	instruments [instrCnt]*instrument
	synths      [synthCnt]*synth
	waves       [waveCnt]*wave
	tables      [tableCnt]*table
	grooves     [grooveCnt]*groove

	words     [wordCnt]*word
	wordNames [wordCnt][wordNameLen]byte

	bookmarks struct {
		pulse1 [bookmarkPosCnt]byte
		pulse2 [bookmarkPosCnt]byte
		wave   [bookmarkPosCnt]byte
		noise  [bookmarkPosCnt]byte
	}

	//TODO: bookmarks
	meta struct {
		keyDelay        byte
		keyRepeat       byte
		font            byte
		sync            byte
		colorSet        byte
		clone           byte
		fileChangedFlag byte
		powerSave       byte
		preListen       byte

		totalTime struct {
			days    byte
			hours   byte
			minutes byte
		}
		workTime struct {
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

func (s *Song) clear() {
	s.formatVersion = 4
	s.tempo = 128
	s.transposition = 0
	s.drumMax = 0x6C

	for i := 0; i < rowCnt; i++ {
		s.rows[i].clear()
	}
	for i := 0; i < waveCnt; i++ {
		s.waves[i].clear()
	}
	for i := 0; i < grooveCnt; i++ {
		s.grooves[i].clear()
	}
	for i := 0; i < chainCnt; i++ {
		s.chains[i] = nil
	}
	for i := 0; i < phraseCnt; i++ {
		s.phrases[i] = nil
	}
	for i := 0; i < instrCnt; i++ {
		s.instruments[i] = nil
	}
	for i := 0; i < synthCnt; i++ {
		s.synths[i] = nil
	}
	for i := 0; i < tableCnt; i++ {
		s.tables[i] = nil
	}
	for i := 0; i < wordCnt; i++ {
		s.words[i].clear()
	}
	s.wordNames = wordNameDefault

	/*
		memset(&Song->bookmarks, LSDJ_NO_BOOKMARK, sizeof(Song->bookmarks));
	*/

	s.meta.keyDelay = 7
	s.meta.keyRepeat = 2
	s.meta.font = 0
	s.meta.sync = 0
	s.meta.colorSet = 0
	s.meta.clone = 0
	s.meta.fileChangedFlag = 0
	s.meta.powerSave = 0
	s.meta.preListen = 1

	s.meta.totalTime.days = 0
	s.meta.totalTime.hours = 0
	s.meta.totalTime.minutes = 0
	s.meta.workTime.hours = 0
	s.meta.workTime.minutes = 0
}
