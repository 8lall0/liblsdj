package liblsdj

import (
	"errors"
)

// Fase 1: inizia a salvarti i byte della roba DONE
// Fase 2: trova una struttura migliore per gestire ste cose (potrei fare degli array anziché un gigatipo così

type Song struct {
	Name                      []byte
	Version                   byte
	Phrases                   Phrases
	Bookmarks                 Bookmarks
	Grooves                   Grooves
	ChainAssignments          ChainAssignments
	TableEnvelopes            TableEnvelopes
	Words                     Words
	WordNames                 WordNames
	InstrumentNames           InstrumentNames
	TableAllocationTable      TableAllocationTable
	InstrumentAllocationTable InstrumentAllocationTable
	ChainPhrases              ChainPhrases
	ChainTranspositions       ChainTranspositions
	InstrumentParams          InstrumentParams
	TableTranspositions       TableTranspositions
	Table1                    Tables
	Table2                    Tables
	PhraseAllocations         PhraseAllocations
	ChainAllocations          ChainAllocations
	SynthParams               SynthParams
	WorkHours                 byte
	WorkMinutes               byte
	Tempo                     byte
	Transposition             byte
	TotalDays                 byte
	TotalHours                byte
	TotalMinutes              byte
	TotalTimeChecksum         byte
	KeyDelay                  byte
	KeyRepeat                 byte
	Font                      byte
	SyncMode                  byte
	ColorPalette              byte
	CloneMode                 byte
	FileChanged               byte
	PowerSave                 byte
	PreListen                 byte
	SynthOverwrites           []byte
	DrumMax                   byte
	Phrase                    Phrase
	Waves                     Waves
	PhraseInstruments         PhraseInstruments
	FormatVersion             byte
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

	if err := s.ChainAssignments.Set(b[chainAssignmentsOffset:tableEnvelopesOffset]); err != nil {
		return err
	}

	s.Words = b[wordsOffset:wordNamesOffset]

	s.WordNames = b[wordNamesOffset:Rb1Offset]

	if err := s.InstrumentNames.Set(b[instrumentNamesOffset:emptySpace1]); err != nil {
		return err
	}

	// Bank 1
	if err := s.InstrumentAllocationTable.Set(b[instrumentAllocationTableOffset:chainPhrasesOffset]); err != nil {
		return err
	}

	s.ChainPhrases = b[chainPhrasesOffset:chainTranspositionsOffset] //0x800, 2048 dovrebbe essere 2032, c'è qualcosa che qualquadra non

	s.ChainTranspositions = b[chainTranspositionsOffset:instrumentParamsOffset]

	if err := s.InstrumentParams.Set(b[instrumentParamsOffset:tableTranspositionOffset]); err != nil {
		return err
	}

	if err := s.Table1.Set(b[tableCommand1Offset:tableCommand1ValueOffset], b[tableCommand1ValueOffset:tableCommand2Offset]); err != nil {
		return err
	}

	if err := s.TableEnvelopes.Set(b[tableEnvelopesOffset:wordsOffset]); err != nil {
		return err
	}

	if err := s.TableAllocationTable.Set(b[tableAllocationTableOffset:instrumentAllocationTableOffset]); err != nil {
		return err
	}

	s.ChainPhrases = b[chainPhrasesOffset:chainTranspositionsOffset]

	s.ChainTranspositions = b[chainTranspositionsOffset:instrumentParamsOffset]

	if err := s.TableTranspositions.Set(b[tableTranspositionOffset:tableCommand1Offset]); err != nil {
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
	if err := s.Phrase.SetCommand(b[phraseCommandsOffset:phraseCommandValuesOffset]); err != nil {
		return err
	}

	if err := s.Phrase.SetValue(b[phraseCommandValuesOffset:emptySpace5]); err != nil {
		return err
	}

	// Bank 3
	if err := s.Waves.Set(b[wavesOffset:phraseInstrumentsOffset]); err != nil {
		return err
	}

	if err := s.PhraseInstruments.Set(b[phraseInstrumentsOffset:Rb3Offset]); err != nil {
		return err
	}
	s.FormatVersion = b[formatVersionOffset]

	return nil
}
