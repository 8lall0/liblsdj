package liblsdj

import (
	"errors"
	"fmt"
)

// Fase 1: inizia a salvarti i byte della roba
// Fase 2: trova una struttura migliore per gestire ste cose

type Song struct {
	Name    []byte
	Version byte

	Phrases               []byte
	Bookmarks             []byte
	Grooves               []byte
	ChainAssignments      []byte
	TableEnvelopes        []byte
	WordsOffset           []byte
	WordNamesOffset       []byte
	InstrumentNamesOffset []byte

	TableAllocationTable      []byte
	InstrumentAllocationTable []byte
	ChainPhrases              []byte
	ChainTranspositions       []byte
	InstrumentParams          []byte
	TableTranspositions       []byte
	TableCommand1             []byte
	TableCommand1Value        []byte
	TableCommand2             []byte
	TableCommand2Value        []byte

	// TODO
	Table1 table
	Table2 table

	PhraseAllocations []byte
	ChainAllocations  []byte
	SynthParams       []byte

	WorkHours         byte
	WorkMinutes       byte
	Tempo             byte
	Transposition     byte
	TotalDays         byte
	TotalHours        byte
	TotalMinutes      byte
	TotalTimeChecksum byte
	KeyDelay          byte
	KeyRepeat         byte
	Font              byte
	SyncMode          byte
	ColorPalette      byte

	CloneMode       byte
	FileChanged     byte
	PowerSave       byte
	PreListen       byte
	SynthOverwrites []byte
	DrumMax         byte

	PhraseCommands      []byte
	PhraseCommandValues []byte

	Waves             []byte
	PhraseInstruments []byte
	FormatVersion     byte
}

const (
	bookmarkPerChannelCount = 16
	noBookmarkValue         = 0xFF
)

func checkRB(rb []byte) bool {
	return rb[Rb1Offset] == 'r' && rb[Rb1Offset+1] == 'b' &&
		rb[Rb2Offset] == 'r' && rb[Rb2Offset+1] == 'b' &&
		rb[Rb3Offset] == 'r' && rb[Rb3Offset+1] == 'b'
}

func (s *Song) Init(b []byte) error {
	if len(b) != 0x8000 {
		return errors.New("bad format, not the right lenght")
	}

	if !checkRB(b) {
		return errors.New("rb check has failed")
	}

	// Bank 0
	s.Phrases = b[phraseNotesOffset : bookmarksOffset-1]
	s.Bookmarks = b[bookmarksOffset : emptySpace0-1]
	s.Grooves = b[groovesOffset : chainAssignmentsOffset-1]
	s.ChainAssignments = b[chainAssignmentsOffset : tableEnvelopesOffset-1]
	s.TableEnvelopes = b[tableEnvelopesOffset : wordsOffset-1]
	s.WordsOffset = b[wordsOffset : wordNamesOffset-1]
	s.WordNamesOffset = b[wordNamesOffset : Rb1Offset-1]
	s.InstrumentNamesOffset = b[instrumentNamesOffset : emptySpace1-1]

	// Bank 1
	s.TableAllocationTable = b[tableAllocationTableOffset : instrumentAllocationTableOffset-1]
	s.InstrumentAllocationTable = b[instrumentAllocationTableOffset : chainPhrasesOffset-1]
	s.ChainPhrases = b[chainPhrasesOffset : chainTranspositionsOffset-1]
	s.ChainTranspositions = b[chainTranspositionsOffset : instrumentParamsOffset-1]
	s.InstrumentParams = b[instrumentParamsOffset : tableTranspositionOffset-1]
	s.TableTranspositions = b[tableTranspositionOffset : tableCommand1Offset-1]

	tab1 := &table{
		Command: b[tableCommand1Offset : tableCommand1ValueOffset-1],
		Value:   b[tableCommand1ValueOffset : tableCommand2Offset-1],
	}
	s.Table1 = *tab1
	tab2 := &table{
		Command: b[tableCommand2Offset : tableCommand2ValueOffset-1],
		Value:   b[tableCommand2ValueOffset : Rb2Offset-1],
	}
	s.Table2 = *tab2
	// DELENDA
	s.TableCommand1 = b[tableCommand1Offset : tableCommand1ValueOffset-1]
	s.TableCommand1Value = b[tableCommand1ValueOffset : tableCommand2Offset-1]
	s.TableCommand2 = b[tableCommand2Offset : tableCommand2ValueOffset-1]
	s.TableCommand2Value = b[tableCommand2ValueOffset : Rb2Offset-1]
	// Fine DELENDA

	s.PhraseAllocations = b[phraseAllocationsOffset : chainAllocationsOffset-1]
	s.ChainAllocations = b[chainAllocationsOffset : synthParamsOffset-1]
	s.SynthParams = b[synthParamsOffset : workHoursOffset-1]
	s.WorkHours = b[workHoursOffset]
	s.WorkMinutes = b[workMinutesOffset]
	s.Tempo = b[tempoOffset]
	s.Transposition = b[transpositionOffset]
	s.TotalDays = b[totalDaysOffset]
	s.TotalHours = b[totalHoursOffset]
	s.TotalMinutes = b[totalMinutesOffset]
	s.TotalTimeChecksum = b[totalTimeChecksumOffset]
	s.KeyDelay = b[keyDelayOffset]
	s.KeyRepeat = b[keyRepeatOffset]
	s.Font = b[fontOffset]
	s.SyncMode = b[syncModeOffset]
	s.ColorPalette = b[colorPaletteOffset]
	s.CloneMode = b[cloneModeOffset]
	s.FileChanged = b[fileChangedOffset]
	s.PowerSave = b[powerSaveOffset]
	s.PreListen = b[prelistenOffset]
	s.SynthOverwrites = b[synthOverwritesOffset : emptySpace3-1]
	s.DrumMax = b[drumMaxOffset]

	// Bank 2
	s.PhraseCommands = b[phraseCommandsOffset : phraseCommandValuesOffset-1]
	s.PhraseCommandValues = b[phraseCommandValuesOffset : emptySpace5-1]

	// Bank 3
	s.Waves = b[wavesOffset : phraseInstrumentsOffset-1]
	s.PhraseInstruments = b[phraseInstrumentsOffset : Rb3Offset-1]
	s.FormatVersion = b[formatVersionOffset]

	// indagare, dovrebbero essere almeno 64
	fmt.Println(len(s.Bookmarks))

	return nil
}
