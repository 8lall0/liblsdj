package liblsdj

import "errors"

type Song struct {
	Name              []byte
	Version           byte
	Phrases           []Phrase
	Chains            []Chain
	Tables            []Table
	Instruments       []Instrument
	Bookmarks         []Bookmark
	Waves             []Wave
	Grooves           []Groove
	ChainAssignments  ChainAssignments
	Words             Words
	WordNames         WordNames
	AllocationTable   AllocationTable
	SynthParams       SynthParams
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
	CloneMode         byte
	FileChanged       byte
	PowerSave         byte
	PreListen         byte
	SynthOverwrites   []byte
	DrumMax           byte
	FormatVersion     byte
}

func (s *Song) Init(b []byte) error {
	if len(b) != 0x8000 {
		return errors.New("bad format, not the right lenght")
	}

	if !checkRB(b) {
		return errors.New("rb check has failed")
	}

	if err := s.setPhrases(b); err != nil {
		return err
	}

	if err := s.setChains(b); err != nil {
		return err
	}

	if err := s.setTables(b); err != nil {
		return err
	}

	if err := s.setInstruments(b); err != nil {
		return err
	}

	if err := s.setAllocations(b); err != nil {
		return err
	}

	if err := s.setWaves(b); err != nil {
		return err
	}

	if err := s.setGrooves(b); err != nil {
		return err
	}

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
	s.FormatVersion = b[formatVersionOffset]

	if err := s.ChainAssignments.Set(b[chainAssignmentsOffset:tableEnvelopesOffset]); err != nil {
		return err
	}
	s.Words = b[wordsOffset:wordNamesOffset]
	s.WordNames = b[wordNamesOffset:Rb1Offset]

	return nil
}

func (s *Song) setPhrases(b []byte) error {
	// Set phrases
	phrases := b[phraseNotesOffset:bookmarksOffset]
	commands := b[phraseCommandsOffset:phraseCommandValuesOffset]
	commandValues := b[phraseCommandValuesOffset:emptySpace5]
	instruments := b[phraseInstrumentsOffset:Rb3Offset]

	p, err := setPhrases(phrases, commands, commandValues, instruments)
	if err != nil {
		return err
	}
	s.Phrases = p

	return nil
}

func (s *Song) setChains(b []byte) error {
	// Set chains
	phrases := b[chainPhrasesOffset:chainTranspositionsOffset]
	commands := b[chainTranspositionsOffset:instrumentParamsOffset]

	c, err := setChains(phrases, commands)
	if err != nil {
		return err
	}
	s.Chains = c

	return nil
}

func (s *Song) setTables(b []byte) error {
	// Set tables
	envelopes := b[tableEnvelopesOffset:wordsOffset]
	transpositions := b[tableTranspositionOffset:tableCommand1Offset]
	col1Commands := b[tableCommand1Offset:tableCommand1ValueOffset]
	col1Values := b[tableCommand1ValueOffset:tableCommand2Offset]
	col2Commands := b[tableCommand2Offset:tableCommand2ValueOffset]
	col2Values := b[tableCommand2ValueOffset:Rb2Offset]

	t, err := setTables(envelopes, transpositions, col1Commands, col1Values, col2Commands, col2Values)
	if err != nil {
		return err
	}
	s.Tables = t

	return nil
}

func (s *Song) setAllocations(b []byte) error {
	phrases := b[phraseAllocationsOffset:chainAllocationsOffset]
	chains := b[chainAllocationsOffset:synthParamsOffset]
	instruments := b[instrumentAllocationTableOffset:chainPhrasesOffset]
	tables := b[tableAllocationTableOffset:instrumentAllocationTableOffset]

	if err := s.AllocationTable.Set(phrases, chains, instruments, tables); err != nil {
		return err
	}

	return nil
}

func (s *Song) setBookmarks(b []byte) error {
	bo, err := setBookmarks(b[bookmarksOffset:emptySpace0])

	if err != nil {
		return err
	}
	s.Bookmarks = bo

	return nil
}

func (s *Song) setInstruments(b []byte) error {
	names := b[instrumentNamesOffset:emptySpace1]
	params := b[instrumentParamsOffset:tableTranspositionOffset]

	in, err := setInstruments(names, params)
	if err != nil {
		return err
	}

	s.Instruments = in

	return nil
}

func (s *Song) setWaves(b []byte) error {
	waves := b[wavesOffset:phraseInstrumentsOffset]

	wv, err := setWaves(waves)
	if err != nil {
		return err
	}

	s.Waves = wv

	return nil
}

func (s *Song) setGrooves(b []byte) error {
	grooves := b[groovesOffset:chainAssignmentsOffset]

	gr, err := setGrooves(grooves)
	if err != nil {
		return err
	}

	s.Grooves = gr

	return nil
}
