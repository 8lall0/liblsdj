package liblsdj

import "fmt"

const (
	instrumentWaveVolume0  = 0x00
	instrumentWaveVolume1  = 0x60
	instrumentWaveVolume2  = 0x40
	instrumentWaveVolume3  = 0xA8
	instrumentWavePlayOnce = iota
	instrumentWavePlayLoop
	instrumentWavePlayPingPong
	instrumentWavePlayManual
)

type WaveInstrument struct {
	params                 [instrumentByteCount]byte
	instrType              byte
	volume                 byte // Può avere solo i valori di cui sopra
	speed, length, loopPos byte
	wave                   byte
	channel                byte // 8° byte left right, 1=L, 2=R, 3=LR
	play                   byte
}

func (w *WaveInstrument) setParams(b []byte) {
	if len(b) != instrumentByteCount {
		// do nothing
	}
	copy(w.params[:], b)
	w.instrType = b[0]
	w.volume = b[1]
	w.wave = b[3]
	w.channel = b[7]
	w.play = b[9]

	fmt.Println(w.params)
}

func (w *WaveInstrument) getParamsBytes() []byte {
	return w.params[:]
}

func (w *WaveInstrument) setGetVolume() {
	// set: set_instrument_bits(song, instrument, 1, 0, 8, volume);
	// get: return get_instrument_bits(song, instrument, 1, 0, 8);
}

func (w *WaveInstrument) setGetSynth() {
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

func (w *WaveInstrument) getSetWave() {
	//set:	set_instrument_bits(song, instrument, 3, 0, 8, wave);
	// get: return get_instrument_bits(song, instrument, 3, 0, 8);
}

func (w *WaveInstrument) getSetPlayMode() {
	//set: if (lsdj_song_get_format_version(song) >= 10)
	//set_instrument_bits(song, instrument, 9, 0, 2, ((((uint8_t)mode) + 1) & 0x3));
	//else
	//set_instrument_bits(song, instrument, 9, 0, 2, (uint8_t)mode);

	//get: if (lsdj_song_get_format_version(song) >= 10)
	//return (lsdj_wave_play_mode_t)((get_instrument_bits(song, instrument, 9, 0, 2) - 1) & 0x3);
	//else
	//return (lsdj_wave_play_mode_t)get_instrument_bits(song, instrument, 9, 0, 2);
}

func (w *WaveInstrument) getSetLength() {
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

func (w *WaveInstrument) getSetLoopPos() {
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

func (w *WaveInstrument) setGetRepeat() {
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

func (w *WaveInstrument) getSetSpeed() {
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
