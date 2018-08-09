package liblsdj

type lsdj_playback_mode byte

const (
	lsdj_PLAY_ONCE      lsdj_playback_mode = 0
	lsdj_PLAY_LOOP      lsdj_playback_mode = 1
	lsdj_PLAY_PING_PONG lsdj_playback_mode = 2
	lsdj_PLAY_MANUAL    lsdj_playback_mode = 3
)

type waveT struct {
	insType  int
	panning  panning
	volume   byte
	table    byte // 0x20 or higher = lsdj_NO_TABLE
	automate byte
	wave     struct {
		plvibSpeed       lsdj_plvib_speed
		vibShape         lsdj_vib_shape
		vibratoDirection lsdj_vib_direction
		transpose        byte
		drumMode         byte
		synth            byte
		playback         lsdj_playback_mode
		length           byte
		repeat           byte
		speed            byte
	}
}

func (i *waveT) read(r *vio, ver byte) {
	var b byte

	i.insType = lsdj_INSTR_WAVE
	i.volume = r.readByte()

	b = r.readByte()
	i.wave.synth = (b >> 4) & 0xF
	i.wave.repeat = b & 0xF

	// Bytes 3 and 4 are empty
	r.seek(r.getCur() + 2)

	b = r.readByte()
	i.wave.drumMode = parseDrumMode(b, ver)
	i.wave.transpose = parseTranspose(b, ver)
	i.automate = parseAutomate(b)
	i.wave.vibratoDirection = lsdj_vib_direction(b & 1)

	if int(ver) < 4 {
		switch int(b>>1) & 3 {
		case 0:
			i.wave.plvibSpeed = lsdj_PLVIB_FAST
			i.wave.vibShape = lsdj_VIB_TRIANGLE
		case 1:
			i.wave.plvibSpeed = lsdj_PLVIB_TICK
			i.wave.vibShape = lsdj_VIB_SAWTOOTH
		case 2:
			i.wave.plvibSpeed = lsdj_PLVIB_TICK
			i.wave.vibShape = lsdj_VIB_TRIANGLE
		case 3:
			i.wave.plvibSpeed = lsdj_PLVIB_FAST
			i.wave.vibShape = lsdj_VIB_SQUARE
		}
	} else {
		if b&0x80 == 1 {
			i.wave.plvibSpeed = lsdj_PLVIB_STEP
		} else if b&0x10 == 1 {
			i.wave.plvibSpeed = lsdj_PLVIB_TICK
		} else {
			i.wave.plvibSpeed = lsdj_PLVIB_FAST
		}

		switch int((b >> 1) & 3) {
		case 0:
			i.wave.vibShape = lsdj_VIB_TRIANGLE
		case 1:
			i.wave.vibShape = lsdj_VIB_SAWTOOTH
		case 2:
			i.wave.vibShape = lsdj_VIB_SQUARE
		}
	}

	// TODO: chiedi del kit nell'audio
	//instrument->kit.vibratoDirection = (byte & 1) == 1 ? LSDJ_VIB_UP : LSDJ_VIB_DOWN;

	i.table = parseTable(r.readByte())
	i.panning = parsePanning(r.readByte())
	// Byte 8 is empty
	r.seek(r.getCur() + 1)
	i.wave.playback = parsePlaybackMode(r.readByte())
	// Bytes 10-13 are empty
	r.seek(r.getCur() + 4)
	b = r.readByte()
	i.wave.length = (b >> 4) & 0xF
	i.wave.speed = b & 0xF

	r.seek(r.getCur() + 1)
}

func (i *waveT) write(w *vio, ver byte) {
	var b byte
	w.writeByte(1)
	w.writeByte(createWaveVolumeByte(i.volume))
	w.writeByte(((i.wave.synth & 0xF) << 4) | (i.wave.repeat & 0xF))
	w.writeByte(0)
	w.writeByte(0xFF)

	b = createDrumModeByte(i.wave.drumMode, ver)
	b |= createTransposeByte(i.wave.transpose, ver)
	b |= createAutomateByte(i.automate)
	b |= createVibrationDirectionByte(i.wave.vibratoDirection)
	if ver < 4 {
		if i.wave.vibShape == lsdj_VIB_SAWTOOTH {
			b |= 2
		} else if i.wave.vibShape == lsdj_VIB_SQUARE {
			b |= 6
		} else if i.wave.vibShape == lsdj_VIB_TRIANGLE {
			if i.wave.plvibSpeed != lsdj_PLVIB_FAST {
				b |= 4
			}
		}
	} else {
		b |= (byte(i.wave.vibShape) & 3) << 1
		if i.wave.plvibSpeed != lsdj_PLVIB_TICK {
			b |= 0x10
		} else if i.wave.plvibSpeed != lsdj_PLVIB_STEP {
			b |= 0x80
		}
	}
	w.writeByte(b)
	w.writeByte(createTableByte(i.table))
	w.writeByte(createPanningByte(i.panning))
	w.writeByte(0)
	w.writeByte(createPlaybackModeByte(i.wave.playback))
	w.writeByte(0xD0)
	w.writeByte(0)
	w.writeByte(0)
	w.writeByte(0)
	w.writeByte(((i.wave.length & 0xF) << 4) | (i.wave.speed & 0xF))
	w.writeByte(0)
}

func (i *waveT) clear() {
	i.insType = lsdj_INSTR_WAVE
	i.volume = 3
	i.panning = lsdj_PAN_LEFT_RIGHT
	i.table = lsdj_NO_TABLE
	i.automate = 0

	i.wave.plvibSpeed = lsdj_PLVIB_FAST
	i.wave.vibShape = lsdj_VIB_TRIANGLE
	i.wave.vibratoDirection = lsdj_VIB_UP
	i.wave.transpose = 1
	i.wave.drumMode = 0
	i.wave.synth = 0
	i.wave.playback = lsdj_PLAY_ONCE
	i.wave.length = 0x0F
	i.wave.repeat = 0
	i.wave.speed = 4
}
