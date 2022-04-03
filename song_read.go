package liblsdj

import (
	"errors"
	"fmt"
)

func ReadSong(b []byte) (*Song, error) {
	if len(b) != 0x8000 {
		return nil, errors.New("bad format, not the right lenght")
	}

	if !checkRB(b) {
		return nil, errors.New("rb check has failed")
	}

	s := new(Song)

	if err := s.setPhrases(b); err != nil {
		return nil, err
	}

	if err := s.setChains(b); err != nil {
		return nil, err
	}

	if err := s.setTables(b); err != nil {
		return nil, err
	}

	if err := s.setInstruments(b); err != nil {
		return nil, err
	}

	if err := s.setWords(b); err != nil {
		return nil, err
	}

	if err := s.setAllocations(b); err != nil {
		return nil, err
	}

	if err := s.setWaves(b); err != nil {
		return nil, err
	}

	if err := s.setGrooves(b); err != nil {
		return nil, err
	}

	if err := s.setBookmarks(b); err != nil {
		return nil, err
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
		return nil, err
	}

	return s, nil
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
	phrases := b[chainPhrasesOffset:chainTranspositionsOffset]
	transpositions := b[chainTranspositionsOffset:instrumentParamsOffset]

	c, err := setChains(phrases, transpositions)
	if err != nil {
		return err
	}
	s.Chains = c

	return nil
}

func (s *Song) setTables(b []byte) error {
	envelopes := b[tableEnvelopesOffset:wordsOffset]
	transpositions := b[tableTranspositionOffset:tableCommand1Offset]
	col1Commands := b[tableCommand1Offset:tableCommand1ValueOffset]
	col1Values := b[tableCommand1ValueOffset:tableCommand2Offset]
	col2Commands := b[tableCommand2Offset:tableCommand2ValueOffset]
	col2Values := b[tableCommand2ValueOffset:Rb2Offset]

	t, err := setTables(envelopes, transpositions, col1Commands, col1Values, col2Commands, col2Values)
	if err != nil {
		return fmt.Errorf("setTables: %w", err)
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
		return fmt.Errorf("setAllocations: %w", err)
	}

	return nil
}

func (s *Song) setBookmarks(b []byte) error {
	bo, err := setBookmarks(b[bookmarksOffset:emptySpace0])

	if err != nil {
		return fmt.Errorf("setBookmarks: %w", err)
	}
	s.Bookmarks = bo

	return nil
}

func (s *Song) setInstruments(b []byte) error {
	names := b[instrumentNamesOffset:emptySpace1]
	params := b[instrumentParamsOffset:tableTranspositionOffset]

	in, err := setInstruments(names, params)
	if err != nil {
		return fmt.Errorf("setInstruments: %w", err)
	}

	s.Instruments = in

	return nil
}

func (s *Song) setWaves(b []byte) error {
	waves := b[wavesOffset:phraseInstrumentsOffset]

	wv, err := setWaves(waves)
	if err != nil {
		return fmt.Errorf("setWaves: %w", err)
	}

	s.Waves = wv

	return nil
}

func (s *Song) setGrooves(b []byte) error {
	grooves := b[groovesOffset:chainAssignmentsOffset]

	gr, err := setGrooves(grooves)
	if err != nil {
		return fmt.Errorf("setGrooves: %w", err)
	}

	s.Grooves = gr

	return nil
}

func (s *Song) setWords(b []byte) error {
	values := b[wordsOffset:wordNamesOffset]
	names := b[wordNamesOffset:Rb1Offset]
	wr, err := setWords(names, values)
	if err != nil {
		return fmt.Errorf("setWords: %w", err)
	}

	s.Words = wr

	return nil
}
