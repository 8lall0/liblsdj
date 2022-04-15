package liblsdj

const (
	instrumentKitLengthAuto = 0x0 //! The value of a InstrumentKit length set to AUTO
	instrumentKitLoopOff    = iota
	instrumentKitLoopOn
	instrumentKitLoopAttack
	instrumentKitDistortionClip = iota
	instrumentKitDistortionShape
	instrumentKitDistortionShape2
	instrumentKitDistortionWrap
)

type KitInstrument struct {
	params           [instrumentByteCount]byte
	instrType        byte
	volume           byte
	loop1Atk         byte
	loop1Speed       byte
	loop1Kit         byte
	length1          byte
	pitchStep        byte
	loop1            byte
	loop2            byte
	pitchSpeed       byte
	tableSpeed       byte
	vibratoShape     byte
	vibratoDirection byte
	tableOnOff       byte
	table            byte
	output           byte
	loop2Atk         byte
	loop2Kit         byte
	dist             byte
	length2          byte
	offset1          byte
	offset2          byte
}

func (k *KitInstrument) setParams(b []byte) {
	if len(b) != instrumentByteCount {
		// do nothing
	}
	copy(k.params[:], b)
}

func (k *KitInstrument) getParamsBytes() []byte {
	return k.params[:]
}

func (k *KitInstrument) setVolume(volume byte) {

	/*void lsdj_instrument_kit_set_volume(lsdj_song_t* song, uint8_t instrument, uint8_t volume)
	{
	lsdj_instrument_wave_set_volume(song, instrument, volume);
	}*/
}

func (k *KitInstrument) getVolume() {
	//	return lsdj_instrument_wave_get_volume(song, instrument);

}

func (k *KitInstrument) setPitch() {
	//	set_instrument_bits(song, instrument, 8, 0, 8, pitch);
}

func (k *KitInstrument) getPitch() {
	//	return get_instrument_bits(song, instrument, 8, 0, 8);
}

func (k *KitInstrument) setHalfSpeed() {
	//	set_instrument_bits(song, instrument, 2, 6, 1, halfSpeed ? 1 : 0);
}

func (k *KitInstrument) getHalfSpeed() {
	//return get_instrument_bits(song, instrument, 2, 6, 1) == 1;
}

func (k *KitInstrument) setDistortionMode() {
	//	set_instrument_bits(song, instrument, 10, 0, 2, (uint8_t)distortion);
}

func (k *KitInstrument) getDistortionMode() {
	//return (lsdj_kit_distortion_mode_t)get_instrument_bits(song, instrument, 10, 0, 2);
}

func (k *KitInstrument) setKit1() {
	//	set_instrument_bits(song, instrument, 2, 0, 5, InstrumentKit);
}

func (k *KitInstrument) getKit1() {
	//return get_instrument_bits(song, instrument, 2, 0, 5);
}

func (k *KitInstrument) setKit2() {
	//	set_instrument_bits(song, instrument, 9, 0, 5, InstrumentKit);
}

func (k *KitInstrument) getKit2() {
	//return get_instrument_bits(song, instrument, 9, 0, 5);
}

func (k *KitInstrument) setOffset1() {
	//set_instrument_bits(song, instrument, 12, 0, 8, offset);
}

func (k *KitInstrument) getOffset1() {
	//return get_instrument_bits(song, instrument, 12, 0, 8);
}

func (k *KitInstrument) setOffset2() {
	//set_instrument_bits(song, instrument, 13, 0, 8, offset);
}

func (k *KitInstrument) getOffset2() {
	//return get_instrument_bits(song, instrument, 13, 0, 8);
}

func (k *KitInstrument) setGetLength1() {
	// set: set_instrument_bits(song, instrument, 3, 0, 8, length);
	// get: return get_instrument_bits(song, instrument, 3, 0, 8);
}

func (k *KitInstrument) setGetLength2() {
	// set: set_instrument_bits(song, instrument, 13, 0, 8, length);
	// get: return get_instrument_bits(song, instrument, 13, 0, 8);
}

func (k *KitInstrument) setGetLoop1() {
	// set: set_instrument_bits(song, instrument, 2, 7, 1, loop == LSDJ_INSTRUMENT_KIT_LOOP_ATTACK ? 1 : 0);
	//set_instrument_bits(song, instrument, 5, 6, 1, loop == LSDJ_INSTRUMENT_KIT_LOOP_ON ? 1 : 0);

	// get:if (get_instrument_bits(song, instrument, 2, 7, 1) == 1)
	//return LSDJ_INSTRUMENT_KIT_LOOP_ATTACK;
	//else
	//return get_instrument_bits(song, instrument, 5, 6, 1);
}

func (k *KitInstrument) setGetLoop2() {
	// set: 	set_instrument_bits(song, instrument, 9, 7, 1, loop == LSDJ_INSTRUMENT_KIT_LOOP_ATTACK ? 1 : 0);
	//	set_instrument_bits(song, instrument, 5, 5, 1, loop == LSDJ_INSTRUMENT_KIT_LOOP_ON ? 1 : 0);

	// get: if (get_instrument_bits(song, instrument, 9, 7, 1) == 1)
	//return LSDJ_INSTRUMENT_KIT_LOOP_ATTACK;
	//else
	//return get_instrument_bits(song, instrument, 5, 5, 1);

}
