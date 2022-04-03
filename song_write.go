package liblsdj

import "fmt"

// WriteSong TODO - DESIDERATA: Ristruttura in modo che siano i singoli tipi a restituire il giusto slice?
func WriteSong(s *Song) ([]byte, error) {
	b := make([]byte, 0x8000)

	s.writeRb(b)
	s.writePhrases(b)
	s.writeChains(b)
	s.writeTables(b)
	s.writeInstruments(b)
	s.writeWords(b)
	s.writeAllocations(b)
	s.writeWaves(b)
	s.writeGrooves(b)
	s.writeBookmarks(b)

	appendTo(b, s.SynthParams, synthParamsOffset)
	appendTo(b, s.ChainAssignments.Get(), chainAssignmentsOffset)

	b[workHoursOffset] = s.WorkHours
	b[workMinutesOffset] = s.WorkMinutes
	b[tempoOffset] = s.Tempo
	b[transpositionOffset] = s.Transposition
	b[totalDaysOffset] = s.TotalDays
	b[totalHoursOffset] = s.TotalHours
	b[totalMinutesOffset] = s.TotalMinutes
	b[totalTimeChecksumOffset] = s.TotalTimeChecksum
	b[keyDelayOffset] = s.KeyDelay
	b[keyRepeatOffset] = s.KeyRepeat
	b[fontOffset] = s.Font
	b[syncModeOffset] = s.SyncMode
	b[colorPaletteOffset] = s.ColorPalette
	b[cloneModeOffset] = s.CloneMode
	b[fileChangedOffset] = s.FileChanged
	b[powerSaveOffset] = s.PowerSave
	b[prelistenOffset] = s.PreListen
	appendTo(b, s.SynthOverwrites, synthOverwritesOffset)
	b[drumMaxOffset] = s.DrumMax
	b[formatVersionOffset] = s.FormatVersion

	if !checkRB(b) {
		fmt.Println("Errore rb")
	}

	return b, nil
}

func appendTo(buffer, input []byte, index int) {
	buffer = append(buffer[:index], input...)
}

func (s *Song) writePhrases(b []byte) {
	phrases := make([]byte, 0)
	commands := make([]byte, 0)
	values := make([]byte, 0)
	instruments := make([]byte, 0)

	for _, v := range s.Phrases {
		phrases = append(phrases, v.Phrase[:]...)
		commands = append(commands, v.Command[:]...)
		values = append(phrases, v.Value[:]...)
		instruments = append(phrases, v.Instruments[:]...)
	}

	appendTo(b, phrases, phraseNotesOffset)
	appendTo(b, commands, phraseCommandsOffset)
	appendTo(b, values, phraseCommandValuesOffset)
	appendTo(b, instruments, phraseInstrumentsOffset)
}

func (s *Song) writeBookmarks(b []byte) {
	bookm := make([]byte, 0)

	for _, v := range s.Bookmarks {
		bookm = append(bookm, v...)
	}

	appendTo(b, bookm, bookmarksOffset)
}

func (s *Song) writeChains(b []byte) {
	phrases := make([]byte, 0)
	transpositions := make([]byte, 0)

	for _, v := range s.Chains {
		phrases = append(phrases, v.phrase[:]...)
		transpositions = append(transpositions, v.transposition[:]...)
	}

	appendTo(b, phrases, chainPhrasesOffset)
	appendTo(b, transpositions, chainTranspositionsOffset)
}

func (s *Song) writeAllocations(b []byte) {
	appendTo(b, s.AllocationTable.Phrases, phraseAllocationsOffset)
	appendTo(b, s.AllocationTable.Chains, chainAllocationsOffset)
	appendTo(b, s.AllocationTable.Instruments, instrumentAllocationTableOffset)
	appendTo(b, s.AllocationTable.Tables, tableAllocationTableOffset)
}

func (s *Song) writeWords(b []byte) {
	values := make([]byte, 0)
	names := make([]byte, 0)

	for _, v := range s.Words {
		values = append(values, v.value[:]...)
		names = append(names, v.name[:]...)
	}

	appendTo(b, values, wordsOffset)
	appendTo(b, names, wordNamesOffset)
}

func (s *Song) writeGrooves(b []byte) {
	grooves := make([]byte, 0)

	for _, v := range s.Grooves {
		grooves = append(grooves, v[:]...)
	}

	appendTo(b, grooves, groovesOffset)
}

func (s *Song) writeWaves(b []byte) {
	waves := make([]byte, 0)

	for _, v := range s.Waves {
		waves = append(waves, v[:]...)
	}

	appendTo(b, waves, wavesOffset)
}

func (s *Song) writeInstruments(b []byte) {
	names := make([]byte, 0)
	params := make([]byte, 0)

	for _, v := range s.Instruments {
		names = append(names, v.Name[:]...)
		params = append(params, v.Params[:]...)
	}

	appendTo(b, params, instrumentParamsOffset)
	appendTo(b, names, instrumentNamesOffset)
}

func (s *Song) writeTables(b []byte) {
	envelopes := make([]byte, 0)
	transpositions := make([]byte, 0)
	col1com := make([]byte, 0)
	col1val := make([]byte, 0)
	col2com := make([]byte, 0)
	col2val := make([]byte, 0)

	for _, v := range s.Tables {
		envelopes = append(envelopes, v.Envelopes[:]...)
		transpositions = append(transpositions, v.Transposition[:]...)
		col1com = append(col1com, v.Col1.Command[:]...)
		col1val = append(col1val, v.Col1.Value[:]...)
		col2com = append(col2com, v.Col2.Command[:]...)
		col2val = append(col2val, v.Col2.Value[:]...)
	}

	appendTo(b, envelopes, tableEnvelopesOffset)
	appendTo(b, transpositions, tableTranspositionOffset)
	appendTo(b, col1com, tableCommand1Offset)
	appendTo(b, col1val, tableCommand1ValueOffset)
	appendTo(b, col2com, tableCommand2Offset)
	appendTo(b, col2val, tableCommand2ValueOffset)
}

func (s *Song) writeRb(b []byte) {
	offsets := []int{Rb1Offset, Rb2Offset, Rb3Offset}

	for _, i := range offsets {
		b[i] = 'r'
		b[i+1] = 'b'
	}
}
