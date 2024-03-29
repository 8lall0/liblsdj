package liblsdj

type instrwave int

func (i *instrwave) setGetVolume() {
	// set: set_instrument_bits(song, instrument, 1, 0, 8, volume);
	// get: return get_instrument_bits(song, instrument, 1, 0, 8);
}

func (i *instrwave) setGetSynth() {
	// set:
	//if (lsdj_song_get_format_version(song) >= 16)
	//	set_instrument_bits(song, instrument, 3, 0, 8, (uint8_t)(synth << 4));
	//else
	//set_instrument_bits(song, instrument, 2, 4, 4, synth);
	// get: if (lsdj_song_get_format_version(song) >= 16)
	//return get_instrument_bits(song, instrument, 3, 4, 4);
	//else
	//return get_instrument_bits(song, instrument, 2, 4, 4);
}

func (i *instrwave) getSetWave() {
	//set:	set_instrument_bits(song, instrument, 3, 0, 8, wave);
	// get: return get_instrument_bits(song, instrument, 3, 0, 8);
}

func (i *instrwave) getSetPlayMode() {
	//set: if (lsdj_song_get_format_version(song) >= 10)
	//set_instrument_bits(song, instrument, 9, 0, 2, ((((uint8_t)mode) + 1) & 0x3));
	//else
	//set_instrument_bits(song, instrument, 9, 0, 2, (uint8_t)mode);

	//get: if (lsdj_song_get_format_version(song) >= 10)
	//return (lsdj_wave_play_mode_t)((get_instrument_bits(song, instrument, 9, 0, 2) - 1) & 0x3);
	//else
	//return (lsdj_wave_play_mode_t)get_instrument_bits(song, instrument, 9, 0, 2);
}

func (i *instrwave) getSetLength() {
	//set:const uint8_t version = lsdj_song_get_format_version(song);
	//if (version >= 7)
	//set_instrument_bits(song, instrument, 10, 0, 4, 0xF - length);
	//else if (version == 6)
	//set_instrument_bits(song, instrument, 10, 0, 4, length);
	//else
	//set_instrument_bits(song, instrument, 14, 4, 4, length);

	//get:const uint8_t version = lsdj_song_get_format_version(song);
	//
	//if (version >= 7)
	//return 0xF - get_instrument_bits(song, instrument, 10, 0, 4);
	//else if (version == 6)
	//return get_instrument_bits(song, instrument, 10, 0, 4);
	//else
	//return get_instrument_bits(song, instrument, 14, 4, 4);
}

func (i *instrwave) getSetLoopPos() {
	// set: if (lsdj_song_get_format_version(song) >= 9)
	//set_instrument_bits(song, instrument, 2, 0, 4, pos & 0xF);
	//else
	//set_instrument_bits(song, instrument, 2, 0, 4, (pos & 0xF) ^ 0x0F);

	//get:const uint8_t byte = get_instrument_bits(song, instrument, 2, 0, 4);
	//
	//if (lsdj_song_get_format_version(song) >= 9)
	//return byte & 0xF;
	//else
	//return (byte & 0xF) ^ 0x0F;
}

func (i *instrwave) setGetRepeat() {
	//set: if (lsdj_song_get_format_version(song) >= 9)
	//	set_instrument_bits(song, instrument, 2, 0, 4, (repeat & 0xF) ^ 0xF);
	//else
	//set_instrument_bits(song, instrument, 2, 0, 4, repeat & 0xF);

	// get: const uint8_t byte = get_instrument_bits(song, instrument, 2, 0, 4);
	//
	//if (lsdj_song_get_format_version(song) >= 9)
	//return (byte & 0xF) ^ 0xF;
	//else
	//return byte & 0xF;
}

func (i *instrwave) getSetSpeed() {
	//set:
	//const uint8_t version = lsdj_song_get_format_version(song);
	//
	//// Speed is stored as starting at 0, but displayed as starting at 1, so subtract 1
	//speed -= 1;
	//
	//if (version >= 7)
	//	set_instrument_bits(song, instrument, 11, 0, 8, speed - 3);
	//else if (version == 6)
	//	set_instrument_bits(song, instrument, 11, 0, 8, speed);
	//else {
	//	if (speed > 0x0F)
	//	return false;
	//
	//	set_instrument_bits(song, instrument, 14, 0, 4, speed);
	// return true

	//get:
	//const uint8_t version = lsdj_song_get_format_version(song);
	//
	//// Read the speed value
	//uint8_t speed = 0;
	//if (version >= 7)
	//speed = get_instrument_bits(song, instrument, 11, 0, 8) + 3;
	//else if (version == 6)
	//speed = get_instrument_bits(song, instrument, 11, 0, 8);
	//else
	//speed = get_instrument_bits(song, instrument, 14, 0, 4);
	//
	//// Speed is stored as starting at 0, but displayed as starting at 1, so add 1
	//return speed + 1;

}
