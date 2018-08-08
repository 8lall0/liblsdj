package liblsdj

type lsdj_playback_mode byte

const (
	lsdj_PLAY_ONCE      lsdj_playback_mode = 0
	lsdj_PLAY_LOOP      lsdj_playback_mode = 1
	lsdj_PLAY_PING_PONG lsdj_playback_mode = 2
	lsdj_PLAY_MANUAL    lsdj_playback_mode = 3
)

type waveT struct {
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
