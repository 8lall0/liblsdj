package liblsdj

import "fmt"

var b [0x8000]byte

// WriteSong TODO - DESIDERATA: Ristruttura, così è orribile. Però fa il suo dovere.
func WriteSong(s *Song) ([]byte, error) {
	s.writeRb()
	s.writePhrases()
	s.writeChains()
	s.writeTables()
	s.writeInstruments()
	s.writeWords()
	s.writeAllocations()
	s.writeWaves()
	s.writeGrooves()
	s.writeBookmarks()

	appendTo(s.SynthParams, synthParamsOffset)
	appendTo(s.ChainAssignments.Get(), chainAssignmentsOffset)

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
	appendTo(s.SynthOverwrites, synthOverwritesOffset)
	b[drumMaxOffset] = s.DrumMax
	b[formatVersionOffset] = s.FormatVersion

	// TODO capire come funziona emptySpace3
	appendTo([]byte{255, 255, 255, 255}, 0x3FC6)

	if !checkRB(b[:]) {
		fmt.Println("Errore rb")
	}

	return b[:], nil
}

func appendTo(input []byte, index int) {
	for i, v := range input {
		b[index+i] = v
	}
}

func (s *Song) writePhrases() {
	phrases := make([]byte, 0)
	commands := make([]byte, 0)
	values := make([]byte, 0)
	instruments := make([]byte, 0)

	for _, v := range s.Phrases {
		phrases = append(phrases, v.Phrase[:]...)
		commands = append(commands, v.Command[:]...)
		values = append(values, v.Value[:]...)
		instruments = append(instruments, v.Instruments[:]...)
	}

	appendTo(phrases, phraseNotesOffset)
	appendTo(commands, phraseCommandsOffset)
	appendTo(values, phraseCommandValuesOffset)
	appendTo(instruments, phraseInstrumentsOffset)
}

func (s *Song) writeBookmarks() {
	bookm := make([]byte, 0)

	for _, v := range s.Bookmarks {
		bookm = append(bookm, v[:]...)
	}

	appendTo(bookm, bookmarksOffset)
}

func (s *Song) writeChains() {
	phrases := make([]byte, 0)
	transpositions := make([]byte, 0)

	for _, v := range s.Chains {
		phrases = append(phrases, v.phrase[:]...)
		transpositions = append(transpositions, v.transposition[:]...)
	}

	appendTo(phrases, chainPhrasesOffset)
	appendTo(transpositions, chainTranspositionsOffset)
}

func (s *Song) writeAllocations() {
	appendTo(s.AllocationTable.Phrases[:], phraseAllocationsOffset)
	appendTo(s.AllocationTable.Chains[:], chainAllocationsOffset)
	appendTo(s.AllocationTable.Instruments[:], instrumentAllocationTableOffset)
	appendTo(s.AllocationTable.Tables[:], tableAllocationTableOffset)
}

func (s *Song) writeWords() {
	values := make([]byte, 0)
	names := make([]byte, 0)

	for _, v := range s.Words {
		values = append(values, v.value[:]...)
		names = append(names, v.name[:]...)
	}

	appendTo(values, wordsOffset)
	appendTo(names, wordNamesOffset)
}

func (s *Song) writeGrooves() {
	grooves := make([]byte, 0)

	for _, v := range s.Grooves {
		grooves = append(grooves, v[:]...)
	}

	appendTo(grooves, groovesOffset)
}

func (s *Song) writeWaves() {
	waves := make([]byte, 0)

	for _, v := range s.Waves {
		waves = append(waves, v[:]...)
	}

	appendTo(waves, wavesOffset)
}

func (s *Song) writeInstruments() {
	names := make([]byte, 0)
	params := make([]byte, 0)

	for _, v := range s.Instruments {
		names = append(names, v.Name[:]...)
		params = append(params, v.Params[:]...)
	}

	appendTo(params, instrumentParamsOffset)
	appendTo(names, instrumentNamesOffset)
}

func (s *Song) writeTables() {
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

	appendTo(envelopes, tableEnvelopesOffset)
	appendTo(transpositions, tableTranspositionOffset)
	appendTo(col1com, tableCommand1Offset)
	appendTo(col1val, tableCommand1ValueOffset)
	appendTo(col2com, tableCommand2Offset)
	appendTo(col2val, tableCommand2ValueOffset)
}

func (s *Song) writeRb() {
	offsets := []int{Rb1Offset, Rb2Offset, Rb3Offset}

	for _, i := range offsets {
		b[i] = 'r'
		b[i+1] = 'b'
	}
}
