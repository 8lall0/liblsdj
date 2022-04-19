package liblsdj

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
	params     [instrumentByteCount]byte
	instrType  byte
	volume     byte
	loopPos    byte
	waveSynth  instrWaveSynth
	byte5param byte5Param
	tableParam instrTableParam
	output     byte
	cmdRate    byte
	playType   byte
	length     byte
	speed      byte
	fineTune   byte
}

func (w *WaveInstrument) setParams(b []byte) {
	if len(b) != instrumentByteCount {
		// do nothing
	}
	copy(w.params[:], b)

	w.instrType = b[0]
	w.volume = b[1]
	w.loopPos = b[2]
	w.waveSynth.set(b[3])
	// 4 is empty
	w.byte5param.set(b[5])
	w.tableParam.set(b[6])
	w.output = b[7]
	w.cmdRate = b[8]
	w.playType = b[9]
	w.length = b[10]
	w.speed = b[11]
	w.fineTune = b[12]
}

func (w *WaveInstrument) getParamsBytes() []byte {
	return w.params[:]
}
