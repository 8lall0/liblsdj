package liblsdj

import (
	"fmt"
	"io"
)

func (s *Song) readBank0(r io.ReadSeeker) {
	for i := 0; i < phraseCnt; i++ {
		if s.phrases[i] != nil {
			if _, err := io.ReadFull(r, s.phrases[i].notes[:]); err != nil {
				panic(err)
			}
		} else {
			if _, err := r.Seek(phraseLen, io.SeekCurrent); err != nil {
				panic(err)
			}
		}
	}

	if _, err := io.ReadFull(r, s.bookmarks.pulse1[:]); err != nil {
		panic(err)
	}
	if _, err := io.ReadFull(r, s.bookmarks.pulse2[:]); err != nil {
		panic(err)
	}
	if _, err := io.ReadFull(r, s.bookmarks.wave[:]); err != nil {
		panic(err)
	}
	if _, err := io.ReadFull(r, s.bookmarks.noise[:]); err != nil {
		panic(err)
	}
	if _, err := io.ReadFull(r, s.reserved1030[:]); err != nil {
		panic(err)
	}
	for i := 0; i < grooveCnt; i++ {
		if _, err := io.ReadFull(r, s.grooves[i][:]); err != nil {
			panic(err)
		}
	}
	for i := 0; i < rowCnt; i++ {
		s.rows[i].read(r)
	}

	for i := 0; i < tableCnt; i++ {
		if s.tables[i] != nil {
			if _, err := io.ReadFull(r, s.tables[i].volumes[:]); err != nil {
				panic(err)
			}
		} else {
			if _, err := r.Seek(tableLen, io.SeekCurrent); err != nil {
				panic(err)
			}
		}
	}

	for i := 0; i < wordCnt; i++ {
		if _, err := io.ReadFull(r, s.words[i].allophones[:]); err != nil {
			panic(err)
		}
		if _, err := io.ReadFull(r, s.words[i].lenghts[:]); err != nil {
			panic(err)
		}
	}
	for i := 0; i < wordCnt; i++ {
		if _, err := io.ReadFull(r, s.wordNames[i][:]); err != nil {
			panic(err)
		}
	}

	//RB
	if _, err := r.Seek(2, io.SeekCurrent); err != nil {
		panic(err)
	}

	for i := 0; i < instrCnt; i++ {
		if s.instruments[i] != nil {
			if _, err := io.ReadFull(r, s.instruments[i].name[:]); err != nil {
				panic(err)
			}
		} else {
			if _, err := r.Seek(instrumentNameLen, io.SeekCurrent); err != nil {
				panic(err)
			}
		}
	}

	if _, err := io.ReadFull(r, s.reserved1fba[:]); err != nil {
		panic(err)
	}
}

func (s *Song) readBank1(r io.ReadSeeker, version byte) {
	if _, err := io.ReadFull(r, s.reserved2000[:]); err != nil {
		panic(err)
	}

	if _, err := r.Seek(tableAllocTableSize+instrAllocTableSize, io.SeekCurrent); err != nil {
		panic(err)
	}

	for i := 0; i < chainCnt; i++ {
		if s.chains[i] != nil {
			if _, err := io.ReadFull(r, s.chains[i].phrases[:]); err != nil {
				panic(err)
			}
		} else {
			if _, err := r.Seek(chainLen, io.SeekCurrent); err != nil {
				panic(err)
			}
		}
	}

	for i := 0; i < chainCnt; i++ {
		if s.chains[i] != nil {
			if _, err := io.ReadFull(r, s.chains[i].transpositions[:]); err != nil {
				panic(err)
			}
		} else {
			if _, err := r.Seek(chainLen, io.SeekCurrent); err != nil {
				panic(err)
			}
		}
	}

	for i := 0; i < instrCnt; i++ {
		if s.instruments[i] != nil {
			s.instruments[i].read(r, version)
		} else {
			if _, err := r.Seek(16, io.SeekCurrent); err != nil {
				panic(err)
			}
		}
	}

	for i := 0; i < tableCnt; i++ {
		if s.tables[i] != nil {
			if _, err := io.ReadFull(r, s.tables[i].transpositions[:]); err != nil {
				panic(err)
			}
		} else {
			if _, err := r.Seek(tableLen, io.SeekCurrent); err != nil {
				panic(err)
			}
		}
	}

	// Twice command1
	for i := 0; i < tableCnt; i++ {
		if s.tables[i] != nil {
			for j := 0; j < tableLen; j++ {
				s.tables[i].commands1[j].command, _ = readByte(r)
			}
		} else {
			if _, err := r.Seek(tableLen, io.SeekCurrent); err != nil {
				panic(err)
			}
		}
	}

	for i := 0; i < tableCnt; i++ {
		if s.tables[i] != nil {
			for j := 0; j < tableLen; j++ {
				s.tables[i].commands1[j].command, _ = readByte(r)
			}
		} else {
			if _, err := r.Seek(tableLen, io.SeekCurrent); err != nil {
				panic(err)
			}
		}
	}

	// Twice command2
	for i := 0; i < tableCnt; i++ {
		if s.tables[i] != nil {
			for j := 0; j < tableLen; j++ {
				s.tables[i].commands2[j].command, _ = readByte(r)
			}
		} else {
			if _, err := r.Seek(tableLen, io.SeekCurrent); err != nil {
				panic(err)
			}
		}
	}

	for i := 0; i < tableCnt; i++ {
		if s.tables[i] != nil {
			for j := 0; j < tableLen; j++ {
				s.tables[i].commands2[j].command, _ = readByte(r)
			}
		} else {
			if _, err := r.Seek(tableLen, io.SeekCurrent); err != nil {
				panic(err)
			}
		}
	}

	//RB
	if _, err := r.Seek(2, io.SeekCurrent); err != nil {
		panic(err)
	}
	if _, err := r.Seek(phraseAllocTableSize+chainAllocTableSize, io.SeekCurrent); err != nil {
		panic(err)
	}

	for i := 0; i < synthCnt; i++ {
		s.synths[i].readSoftSynthParams(r)
	}

	s.meta.workTime.hours, _ = readByte(r)
	s.meta.workTime.minutes, _ = readByte(r)
	s.tempo, _ = readByte(r)
	s.transposition, _ = readByte(r)
	s.meta.totalTime.days, _ = readByte(r)
	s.meta.totalTime.hours, _ = readByte(r)
	s.meta.totalTime.minutes, _ = readByte(r)
	s.reserved3fb9, _ = readByte(r)
	s.meta.keyDelay, _ = readByte(r)
	s.meta.keyRepeat, _ = readByte(r)
	s.meta.font, _ = readByte(r)
	s.meta.sync, _ = readByte(r)
	s.meta.colorSet, _ = readByte(r)
	s.reserved3fbf, _ = readByte(r)
	s.meta.clone, _ = readByte(r)
	s.meta.fileChangedFlag, _ = readByte(r)
	s.meta.powerSave, _ = readByte(r)
	s.meta.preListen, _ = readByte(r)

	var waveSynthOverwriteLocks [2]byte
	if _, err := io.ReadFull(r, waveSynthOverwriteLocks[:]); err != nil {
		panic(err)
	}

	for i := 0; i < synthCnt; i++ {
		s.synths[i].overwritten = ((waveSynthOverwriteLocks[1-(i/8)]) >> uint(i%8)) & 1
	}

	if _, err := io.ReadFull(r, s.reserved3fc6[:]); err != nil {
		panic(err)
	}

	s.drumMax, _ = readByte(r)

	if _, err := io.ReadFull(r, s.reserved3fd1[:]); err != nil {
		panic(err)
	}
}

func (s *Song) readBank2(r io.ReadSeeker) {
	for i := 0; i < phraseCnt; i++ {
		if s.phrases[i] != nil {
			for j := 0; j < phraseLen; j++ {
				s.phrases[i].commands[j].value, _ = readByte(r)
			}
		} else {
			if _, err := r.Seek(phraseLen, io.SeekCurrent); err != nil {
				panic(err)
			}
		}
	}
	for i := 0; i < phraseCnt; i++ {
		if s.phrases[i] != nil {
			for j := 0; j < phraseLen; j++ {
				s.phrases[i].commands[j].value, _ = readByte(r)
			}
		} else {
			if _, err := r.Seek(phraseLen, io.SeekCurrent); err != nil {
				panic(err)
			}
		}
	}

	if _, err := io.ReadFull(r, s.reserved5fe0[:]); err != nil {
		panic(err)
	}
}

func (s *Song) readBank3(r io.ReadSeeker) {
	for i := 0; i < waveCnt; i++ {
		if _, err := io.ReadFull(r, s.waves[i][:]); err != nil {
			panic(err)
		}
	}

	for i := 0; i < phraseCnt; i++ {
		if s.phrases[i] != nil {
			if _, err := io.ReadFull(r, s.phrases[i].instruments[:]); err != nil {
				panic(err)
			}
		} else {
			if _, err := r.Seek(phraseLen, io.SeekCurrent); err != nil {
				panic(err)
			}
		}
	}

	//rb
	if _, err := r.Seek(2, io.SeekCurrent); err != nil {
		panic(err)
	}
	if _, err := io.ReadFull(r, s.reserved7ff2[:]); err != nil {
		panic(err)
	}
	// Already read the version number
	if _, err := r.Seek(1, io.SeekCurrent); err != nil {
		panic(err)
	}
}

func checkRb(r io.ReadSeeker, position int64) bool {
	var rb [2]byte
	if _, err := r.Seek(position, io.SeekStart); err != nil {
		panic(err)
	}

	if _, err := io.ReadFull(r, rb[:]); err != nil {
		panic(err)
	}

	return rb[0] == 'r' && rb[1] == 'b'
}

func ReadSong(r io.ReadSeeker, version byte) (s *Song, err error) {
	var tableAllocTable [tableAllocTableSize]byte
	var instrAllocTable [instrAllocTableSize]byte
	var phraseAllocTable [phraseAllocTableSize]byte
	var chainAllocTable [chainAllocTableSize]byte

	s = new(Song)
	s.clear()

	pos, _ := r.Seek(0, io.SeekCurrent)

	if !checkRb(r, pos+0x1E78) {
		return s, fmt.Errorf("memory flag 'rb' not found at 0x1E78")
	}
	if !checkRb(r, pos+0x3E80) {
		return s, fmt.Errorf("memory flag 'rb' not found at 0x3E80")
	}
	if !checkRb(r, pos+0x7FF0) {
		return s, fmt.Errorf("memory flag 'rb' not found at 0x7FF0")
	}

	if _, err := r.Seek(pos+int64(0x7FFF), io.SeekStart); err != nil {
		panic(err)
	}

	s.formatVersion, _ = readByte(r)

	if _, err := r.Seek(pos+int64(0x2020), io.SeekStart); err != nil {
		panic(err)
	}
	if _, err := io.ReadFull(r, tableAllocTable[:]); err != nil {
		panic(err)
	}
	if _, err := io.ReadFull(r, instrAllocTable[:]); err != nil {
		panic(err)
	}
	if _, err := r.Seek(pos+int64(0x3E82), io.SeekStart); err != nil {
		panic(err)
	}
	if _, err := io.ReadFull(r, phraseAllocTable[:]); err != nil {
		panic(err)
	}
	if _, err := io.ReadFull(r, chainAllocTable[:]); err != nil {
		panic(err)
	}
	/*
		Qui ragiono in logica opposta: se non c'è tabella pongo quell'indice a nil
	*/
	for i := 0; i < tableAllocTableSize; i++ {
		if tableAllocTable[i] == 0 {
			s.tables[i] = nil
		} else {
			s.tables[i] = new(table)
			s.tables[i].clear()
		}
	}

	for i := 0; i < instrAllocTableSize; i++ {
		if instrAllocTable[i] == 0 {
			s.instruments[i] = nil
		} else {
			s.instruments[i] = new(instrument)
			s.instruments[i].clearAsPulse()
		}
	}

	for i := 0; i < chainAllocTableSize; i++ {
		if (chainAllocTable[i/8]>>uint(i%8))&1 == 0 {
			s.chains[i] = nil
		} else {
			s.chains[i] = new(chain)
			s.chains[i].clear()
		}
	}

	for i := 0; i < phraseAllocTableSize; i++ {
		if (phraseAllocTable[i/8]>>uint(i%8))&1 == 0 {
			s.phrases[i] = nil
		} else {
			s.phrases[i] = new(phrase)
			s.phrases[i].clear()
		}
	}

	if _, err := r.Seek(pos, io.SeekStart); err != nil {
		panic(err)
	}

	s.readBank0(r)
	s.readBank1(r, version)
	s.readBank2(r)
	s.readBank3(r)

	return
}
