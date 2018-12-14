package lsdj

import "io"

type playbackMode byte

const (
	playOnce playbackMode = iota
	playLoop
	playPingPong
	playManual
)

type instrumentWave struct {
	plVibSpeed   plVibSpeed
	vibShape     vibShape
	vibDirection vibDirection

	transpose byte
	drumMode  byte
	synth     byte
	playback  playbackMode
	length    byte
	repeat    byte
	speed     byte
}

func (i *instrumentWave) clearInstrument() {
	i.plVibSpeed = plVibFast
	i.vibShape = vibTriangle
	i.vibDirection = vipUp

	i.transpose = 1
	i.drumMode = 0
	i.synth = 0
	i.playback = playOnce
	i.length = 0x0F
	i.repeat = 0
	i.speed = 4
}

func (i *instrumentWave) read(r io.ReadSeeker) {
	//TODO: read_wave_instrument
}
