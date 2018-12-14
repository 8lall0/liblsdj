package lsdj

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

	for i := 0; i < len(s.wordNames); i++ {
		_, _ = w.Write(s.wordNames[i][:])
	}

	_, _ = w.Write([]byte("rb"))



}
