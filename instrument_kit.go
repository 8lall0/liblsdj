package liblsdj

type InstrumentKit struct {
	pippo int
}

func (i *InstrumentKit) setVolume(volume byte) {

	/*void lsdj_instrument_kit_set_volume(lsdj_song_t* song, uint8_t instrument, uint8_t volume)
	{
	lsdj_instrument_wave_set_volume(song, instrument, volume);
	}*/
}

func (i *InstrumentKit) getVolume() {
	//	return lsdj_instrument_wave_get_volume(song, instrument);

}

func (i *InstrumentKit) setPitch() {
	//	set_instrument_bits(song, instrument, 8, 0, 8, pitch);
}

func (i *InstrumentKit) getPitch() {
	//	return get_instrument_bits(song, instrument, 8, 0, 8);
}

func (i *InstrumentKit) setHalfSpeed() {
	//	set_instrument_bits(song, instrument, 2, 6, 1, halfSpeed ? 1 : 0);
}

func (i *InstrumentKit) getHalfSpeed() {
	//return get_instrument_bits(song, instrument, 2, 6, 1) == 1;
}

func (i *InstrumentKit) setDistortionMode() {
	//	set_instrument_bits(song, instrument, 10, 0, 2, (uint8_t)distortion);
}

func (i *InstrumentKit) getDistortionMode() {
	//return (lsdj_kit_distortion_mode_t)get_instrument_bits(song, instrument, 10, 0, 2);
}

func (i *InstrumentKit) setKit1() {
	//	set_instrument_bits(song, instrument, 2, 0, 5, InstrumentKit);
}

func (i *InstrumentKit) getKit1() {
	//return get_instrument_bits(song, instrument, 2, 0, 5);
}

func (i *InstrumentKit) setKit2() {
	//	set_instrument_bits(song, instrument, 9, 0, 5, InstrumentKit);
}

func (i *InstrumentKit) getKit2() {
	//return get_instrument_bits(song, instrument, 9, 0, 5);
}

func (i *InstrumentKit) setOffset1() {
	//set_instrument_bits(song, instrument, 12, 0, 8, offset);
}

func (i *InstrumentKit) getOffset1() {
	//return get_instrument_bits(song, instrument, 12, 0, 8);
}

func (i *InstrumentKit) setOffset2() {
	//set_instrument_bits(song, instrument, 13, 0, 8, offset);
}

func (i *InstrumentKit) getOffset2() {
	//return get_instrument_bits(song, instrument, 13, 0, 8);
}

func (i *InstrumentKit) setGetLength1() {
	// set: set_instrument_bits(song, instrument, 3, 0, 8, length);
	// get: return get_instrument_bits(song, instrument, 3, 0, 8);
}

func (i *InstrumentKit) setGetLength2() {
	// set: set_instrument_bits(song, instrument, 13, 0, 8, length);
	// get: return get_instrument_bits(song, instrument, 13, 0, 8);
}

func (i *InstrumentKit) setGetLoop1() {
	// set: set_instrument_bits(song, instrument, 2, 7, 1, loop == LSDJ_INSTRUMENT_KIT_LOOP_ATTACK ? 1 : 0);
	//set_instrument_bits(song, instrument, 5, 6, 1, loop == LSDJ_INSTRUMENT_KIT_LOOP_ON ? 1 : 0);

	// get:if (get_instrument_bits(song, instrument, 2, 7, 1) == 1)
	//return LSDJ_INSTRUMENT_KIT_LOOP_ATTACK;
	//else
	//return get_instrument_bits(song, instrument, 5, 6, 1);
}

func (i *InstrumentKit) setGetLoop2() {
	// set: 	set_instrument_bits(song, instrument, 9, 7, 1, loop == LSDJ_INSTRUMENT_KIT_LOOP_ATTACK ? 1 : 0);
	//	set_instrument_bits(song, instrument, 5, 5, 1, loop == LSDJ_INSTRUMENT_KIT_LOOP_ON ? 1 : 0);

	// get: if (get_instrument_bits(song, instrument, 9, 7, 1) == 1)
	//return LSDJ_INSTRUMENT_KIT_LOOP_ATTACK;
	//else
	//return get_instrument_bits(song, instrument, 5, 5, 1);

}
