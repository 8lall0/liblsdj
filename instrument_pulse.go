package liblsdj

type lsdj_pulse_wave byte

const (
	lsdj_PULSE_WAVE_PW_125 lsdj_pulse_wave = 0
	lsdj_PULSE_WAVE_PW_25  lsdj_pulse_wave = 1
	lsdj_PULSE_WAVE_PW_50  lsdj_pulse_wave = 2
	lsdj_PULSE_WAVE_PW_75  lsdj_pulse_wave = 3
)

type pulseT struct {
	pulseWidth       lsdj_pulse_wave
	length           byte // 0x40 and above = unlimited
	sweep            byte
	plvibSpeed       lsdj_plvib_speed
	vibShape         lsdj_vib_shape
	vibratoDirection lsdj_vib_direction
	transpose        byte
	drumMode         byte
	pulse2tune       byte
	fineTune         byte
}
