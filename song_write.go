package liblsdj

import "io"

func (s *Song) writeBank0(w io.WriteSeeker) {
	for i := 0; i < phraseCnt; i++ {
		if s.phrases[i] != nil {
			_, _ = w.Write(s.phrases[i].notes[:])
		} else {
			_, _ = w.Write(phraseLenZero[:])
		}
	}

	_, _ = w.Write(s.bookmarks.pulse1[:])
	_, _ = w.Write(s.bookmarks.pulse2[:])
	_, _ = w.Write(s.bookmarks.wave[:])
	_, _ = w.Write(s.bookmarks.noise[:])

	_, _ = w.Write(s.reserved1030[:])

	for i := 0; i < len(s.grooves); i++ {
		_, _ = w.Write(s.grooves[i][:])
	}

	for i := 0; i < len(s.rows); i++ {
		s.rows[i].write(w)
	}

	for i := 0; i < tableCnt; i++ {
		if s.tables[i] != nil {
			_, _ = w.Write(s.tables[i].volumes[:])
		} else {
			_, _ = w.Write(tableLengthZero[:])
		}
	}

	for i := 0; i < wordCnt; i++ {
		s.words[i].write(w)
	}

	for i := 0; i < len(s.wordNames); i++ {
		_, _ = w.Write(s.wordNames[i][:])
	}

	_, _ = w.Write([]byte("rb"))

	for i := 0; i < instrumentNameLen; i++ {
		if s.instruments[i] != nil {
			_, _ = w.Write(s.instruments[i].name[:])
		} else {
			_, _ = w.Write(instrumentNameEmpty[:])
		}
	}

	_, _ = w.Write(s.reserved1fba[:])
}

func (s *Song) writeBank1(w io.WriteSeeker, version byte) {
	var instrAllocTable [instrAllocTableSize]byte
	for i := 0; i < instrCnt; i++ {
		if s.instruments[i] != nil {
			instrAllocTable[i] = 1
		}
	}

	var tableAllocTable [tableAllocTableSize]byte
	for i := 0; i < tableCnt; i++ {
		if s.tables[i] != nil {
			tableAllocTable[i] = 1
		}
	}

	var chainAllocTable [chainAllocTableSize]byte
	for i := 0; i < chainCnt; i++ {
		if s.chains[i] != nil {
			chainAllocTable[i/8] |= 1 << uint(i%8)
		}
	}

	var phraseAllocTable [phraseAllocTableSize]byte
	for i := 0; i < phraseCnt; i++ {
		if s.phrases[i] != nil {
			phraseAllocTable[i/8] |= 1 << uint(i%8)
		}
	}

	_, _ = w.Write(s.reserved2000[:])
	_, _ = w.Write(tableAllocTable[:])
	_, _ = w.Write(instrAllocTable[:])

	for i := 0; i < chainCnt; i++ {
		if s.chains[i] != nil {
			_, _ = w.Write(s.chains[i].phrases[:])
		} else {
			_, _ = w.Write(chainLenFF[:])
		}
	}

	for i := 0; i < chainCnt; i++ {
		if s.chains[i] != nil {
			_, _ = w.Write(s.chains[i].transpositions[:])
		} else {
			_, _ = w.Write(chainLenZero[:])
		}
	}

	for i := 0; i < instrCnt; i++ {
		if s.instruments[i] != nil {
			s.instruments[i].instrument.write(s.instruments[i], w, version)
		} else {
			_, _ = w.Write(instrumentDefault[:])
		}
	}

	for i := 0; i < tableCnt; i++ {
		if s.tables[i] != nil {
			_, _ = w.Write(s.tables[i].transpositions[:])
		} else {
			_, _ = w.Write(tableLengthZero[:])
		}
	}

	// Command1-2 times 2
	for i := 0; i < tableCnt; i++ {
		if s.tables[i] != nil {
			_, _ = w.Write(s.tables[i].getCommand1())
		} else {
			_, _ = w.Write(tableLengthZero[:])
		}
	}

	for i := 0; i < tableCnt; i++ {
		if s.tables[i] != nil {
			_, _ = w.Write(s.tables[i].getCommand1())
		} else {
			_, _ = w.Write(tableLengthZero[:])
		}
	}

	for i := 0; i < tableCnt; i++ {
		if s.tables[i] != nil {
			_, _ = w.Write(s.tables[i].getCommand2())
		} else {
			_, _ = w.Write(tableLengthZero[:])
		}
	}

	for i := 0; i < tableCnt; i++ {
		if s.tables[i] != nil {
			_, _ = w.Write(s.tables[i].getCommand2())
		} else {
			_, _ = w.Write(tableLengthZero[:])
		}
	}

	_, _ = w.Write([]byte("rb"))
	_, _ = w.Write(phraseAllocTable[:])
	_, _ = w.Write(chainAllocTable[:])

	for i := 0; i < synthCnt; i++ {
		s.synths[i].writeSoftSynthParams(w)
	}

	_ = writeByte(s.meta.workTime.hours, w)
	_ = writeByte(s.meta.workTime.minutes, w)

	_ = writeByte(s.tempo, w)
	_ = writeByte(s.transposition, w)

	_ = writeByte(s.meta.totalTime.days, w)
	_ = writeByte(s.meta.totalTime.hours, w)
	_ = writeByte(s.meta.totalTime.minutes, w)

	_ = writeByte(s.reserved3fb9, w)
	_ = writeByte(s.meta.keyDelay, w)
	_ = writeByte(s.meta.keyRepeat, w)
	_ = writeByte(s.meta.font, w)
	_ = writeByte(s.meta.sync, w)
	_ = writeByte(s.meta.colorSet, w)
	_ = writeByte(s.reserved3fbf, w)
	_ = writeByte(s.meta.clone, w)
	_ = writeByte(s.meta.fileChangedFlag, w)
	_ = writeByte(s.meta.powerSave, w)
	_ = writeByte(s.meta.preListen, w)

	var waveSynthOverwriteLocks [2]byte
	for i := 0; i < synthCnt; i++ {
		if s.synths[i].overwritten != 0 {
			waveSynthOverwriteLocks[1-(i/8)] |= 1 << uint(i%8)
		}
	}
	_, _ = w.Write(waveSynthOverwriteLocks[:])
	_, _ = w.Write(s.reserved3fc6[:])
	_ = writeByte(s.drumMax, w)
	_, _ = w.Write(s.reserved3fd1[:])
}

func (s *Song) writeBank2(w io.WriteSeeker) {
	for i := 0; i < phraseCnt; i++ {
		if s.phrases[i] != nil {
			for j := 0; j < phraseLen; j++ {
				_ = writeByte(s.phrases[i].commands[j].command, w)
			}
		} else {
			_, _ = w.Write(phraseLenZero[:])
		}
	}

	for i := 0; i < phraseCnt; i++ {
		if s.phrases[i] != nil {
			for j := 0; j < phraseLen; j++ {
				_ = writeByte(s.phrases[i].commands[j].value, w)
			}
		} else {
			_, _ = w.Write(phraseLenZero[:])
		}
	}

	_, _ = w.Write(s.reserved5fe0[:])
}

func (s *Song) writeBank3(w io.WriteSeeker) {
	for i := 0; i < waveCnt; i++ {
		_, _ = w.Write(s.waves[i][:])
	}

	for i := 0; i < phraseCnt; i++ {
		if s.phrases[i] != nil {
			_, _ = w.Write(s.phrases[i].instruments[:])
		} else {
			_, _ = w.Write(phraseLenFF[:])
		}
	}

	_, _ = w.Write([]byte("rb"))
	_, _ = w.Write(s.reserved7ff2[:])
	_ = writeByte(s.formatVersion, w)
}

func WriteSong(w io.WriteSeeker, s *Song, version byte) {
	s.writeBank0(w)
	s.writeBank1(w, version)
	s.writeBank2(w)
	s.writeBank3(w)
}
