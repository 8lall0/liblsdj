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

	// TODO i bookmarks sono tutti i pulse+wav+no oppure (pul-wa-no)*i???
	_, _ = w.Write(s.bookmarks.pulse1[:])
	_, _ = w.Write(s.bookmarks.pulse2[:])
	_, _ = w.Write(s.bookmarks.wave[:])
	_, _ = w.Write(s.bookmarks.noise[:])
	_, _ = w.Write(s.reserved1030[:])

	for i := 0; i < len(s.grooves); i++ {
		_, _ = w.Write(s.grooves[i][:])
	}

	for i := 0; i < len(s.rows); i++ {
		_ = writeByte(s.rows[i].pulse1, w)
		_ = writeByte(s.rows[i].pulse2, w)
		_ = writeByte(s.rows[i].wave, w)
		_ = writeByte(s.rows[i].noise, w)
	}

	for i := 0; i < tableCnt; i++ {
		if s.tables[i] != nil {
			_, _ = w.Write(s.tables[i].volumes[:])
		} else {
			_, _ = w.Write(tableLengthZero[:])
		}
	}

	for i := 0; i < wordCnt; i++ {
		_, _ = w.Write(s.words[i].allophones[:])
		_, _ = w.Write(s.words[i].lenghts[:])
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

func (s *Song) writeBank1(w io.WriteSeeker) {
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
			chainAllocTable[i/8] |= (1 << uint(i%8))
		}
	}

	var phraseAllocTable [phraseAllocTableSize]byte
	for i := 0; i < phraseCnt; i++ {
		if s.phrases[i] != nil {
			phraseAllocTable[i/8] |= (1 << uint(i%8))
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
			// TODO: instrument write - aggiungi ad interfaccia
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

	// Command1-2 è ripetuto due volte

	for i := 0; i < tableCnt; i++ {
		if s.tables[i] != nil {
			// TODO get command1
		} else {
			_, _ = w.Write(tableLengthZero[:])
		}
	}

	for i := 0; i < tableCnt; i++ {
		if s.tables[i] != nil {
			// TODO get command1
		} else {
			_, _ = w.Write(tableLengthZero[:])
		}
	}

	for i := 0; i < tableCnt; i++ {
		if s.tables[i] != nil {
			// TODO get command2
		} else {
			_, _ = w.Write(tableLengthZero[:])
		}
	}

	for i := 0; i < tableCnt; i++ {
		if s.tables[i] != nil {
			// TODO get command2
		} else {
			_, _ = w.Write(tableLengthZero[:])
		}
	}

	_, _ = w.Write([]byte("rb"))
	_, _ = w.Write(phraseAllocTable[:])
	_, _ = w.Write(chainAllocTable[:])

	for i := 0; i < synthCnt; i++ {
		// TODO write_soft_synth_parameters
	}

	// TODO controlla se viene prima minuti o ore
	_ := writeByte(s.meta.workTime.minutes, w)
	_ := writeByte(s.meta.workTime.hours, w)
	_ := writeByte(s.tempo, w)
	_ := writeByte(s.transposition, w)
	// TODO controlla se viene prima minuti o ore
	_ := writeByte(s.meta.totalTime.minutes, w)
	_ := writeByte(s.meta.totalTime.hours, w)
	_ := writeByte(s.meta.totalTime.days, w)
	_ := writeByte(s.reserved3fb9, w)
	_ := writeByte(s.meta.keyDelay, w)
	_ := writeByte(s.meta.keyRepeat, w)
	_ := writeByte(s.meta.font, w)
	_ := writeByte(s.meta.sync, w)
	_ := writeByte(s.meta.colorSet, w)
	_ := writeByte(s.reserved3fbf, w)
	_ := writeByte(s.meta.clone, w)
	_ := writeByte(s.meta.fileChangedFlag, w)
	_ := writeByte(s.meta.powerSave, w)
	_ := writeByte(s.meta.preListen, w)

	var waveSynthOverwriteLocks [2]byte
	for i := 0; i < synthCnt; i++ {
		if s.synths[i].overwritten != 0 {
			waveSynthOverwriteLocks[1-(i/8)] |= (1 << uint(i%8))
		}
	}
	_, _ = w.Write(waveSynthOverwriteLocks[:])
	_, _ = w.Write(s.reserved3fc6[:])
	_ := writeByte(s.drumMax, w)
	_, _ = w.Write(s.reserved3fd1[:])
}

func (s *Song) writeBank2(w io.WriteSeeker) {
	for i := 0; i < phraseCnt; i++ {
		if s.phrases[i] != nil {
			for j := 0; j < phraseLen; j++ {
				_ := writeByte(s.phrases[i].commands[j].command, w)
			}
		} else {
			_, _ = w.Write(phraseLenZero[:])
		}
	}

	for i := 0; i < phraseCnt; i++ {
		if s.phrases[i] != nil {
			for j := 0; j < phraseLen; j++ {
				_ := writeByte(s.phrases[i].commands[j].value, w)
			}
		} else {
			_, _ = w.Write(phraseLenZero[:])
		}
	}

	_, _ = w.Write(s.reserved5fe0[:])
}

func (s *Song) writeBank3(w io.WriteSeeker) {
	for i := 0; i < waveCnt; i++ {
		_, _ = w.Write(s.waves[i].data[:])
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
	_ := writeByte(s.formatVersion, w)
}
