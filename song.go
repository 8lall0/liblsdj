package liblsdj

import (
	"errors"
)

// Fase 1: inizia a salvarti i byte della roba
// Fase 2: trova una struttura migliore per gestire ste cose

type Song struct {
	Name    []byte
	Version byte

	Phrases               Phrases
	Bookmarks             Bookmarks
	Grooves               Grooves
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

	Table1 Tables
	Table2 Tables

	PhraseAllocations PhraseAllocations
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

	PhraseCommands      PhraseCommands
	PhraseCommandValues PhraseCommandValues

	Waves             Waves
	PhraseInstruments []byte
	FormatVersion     byte
}

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
	if err := s.Phrases.Set(b[phraseNotesOffset:bookmarksOffset]); err != nil {
		return err
	}
	if err := s.Bookmarks.Set(b[bookmarksOffset:emptySpace0]); err != nil {
		return err
	}
	if err := s.Grooves.Set(b[groovesOffset:chainAssignmentsOffset]); err != nil {
		return err
	}

	s.ChainAssignments = b[chainAssignmentsOffset:tableEnvelopesOffset]
	s.TableEnvelopes = b[tableEnvelopesOffset:wordsOffset]
	s.WordsOffset = b[wordsOffset:wordNamesOffset]
	s.WordNamesOffset = b[wordNamesOffset:Rb1Offset]
	s.InstrumentNamesOffset = b[instrumentNamesOffset:emptySpace1]

	// Bank 1
	s.TableAllocationTable = b[tableAllocationTableOffset:instrumentAllocationTableOffset]
	s.InstrumentAllocationTable = b[instrumentAllocationTableOffset:chainPhrasesOffset]
	s.ChainPhrases = b[chainPhrasesOffset:chainTranspositionsOffset]
	s.ChainTranspositions = b[chainTranspositionsOffset:instrumentParamsOffset]
	s.InstrumentParams = b[instrumentParamsOffset:tableTranspositionOffset]
	s.TableTranspositions = b[tableTranspositionOffset:tableCommand1Offset]

	if err := s.Table1.Set(b[tableCommand1Offset:tableCommand1ValueOffset], b[tableCommand1ValueOffset:tableCommand2Offset]); err != nil {
		return err
	}
	if err := s.Table2.Set(b[tableCommand2Offset:tableCommand2ValueOffset], b[tableCommand2ValueOffset:Rb2Offset]); err != nil {
		return err
	}

	if err := s.PhraseAllocations.Set(b[phraseAllocationsOffset:chainAllocationsOffset]); err != nil {
		return err
	}

	s.ChainAllocations = b[chainAllocationsOffset:synthParamsOffset]
	s.SynthParams = b[synthParamsOffset:workHoursOffset]
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
	s.SynthOverwrites = b[synthOverwritesOffset:emptySpace3]
	s.DrumMax = b[drumMaxOffset]

	// Bank 2
	if err := s.PhraseCommands.Set(b[phraseCommandsOffset:phraseCommandValuesOffset]); err != nil {
		return err
	}
	if err := s.PhraseCommandValues.Set(b[phraseCommandValuesOffset:emptySpace5]); err != nil {
		return err
	}

	// Bank 3
	if err := s.Waves.Set(b[wavesOffset:phraseInstrumentsOffset]); err != nil {
		return err
	}

	s.PhraseInstruments = b[phraseInstrumentsOffset:Rb3Offset]
	s.FormatVersion = b[formatVersionOffset]

	return nil
}
