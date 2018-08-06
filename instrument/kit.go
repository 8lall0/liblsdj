package instrument

type lsdj_kit_loop_mode byte
type lsdj_kit_distortion byte
type lsdj_kit_pspeed byte

const (
	LSDJ_KIT_LOOP_OFF    lsdj_kit_loop_mode = 0
	LSDJ_KIT_LOOP_ON     lsdj_kit_loop_mode = 1
	LSDJ_KIT_LOOP_ATTACK lsdj_kit_loop_mode = 2

	LSDJ_KIT_DIST_CLIP   lsdj_kit_distortion = 0xD0
	LSDJ_KIT_DIST_SHAPE  lsdj_kit_distortion = 0xD1
	LSDJ_KIT_DIST_SHAPE2 lsdj_kit_distortion = 0xD2
	LSDJ_KIT_DIST_WRAP   lsdj_kit_distortion = 0xD3

	LSDJ_KIT_PSPEED_FAST lsdj_kit_pspeed = 0
	LSDJ_KIT_PSPEED_SLOW lsdj_kit_pspeed = 1
	LSDJ_KIT_PSPEED_STEP lsdj_kit_pspeed = 2
)

type kitT struct {
	kit1    byte
	offset1 byte
	length1 byte
	loop1   lsdj_kit_loop_mode

	kit2    byte
	offset2 byte
	length2 byte
	loop2   lsdj_kit_loop_mode

	pitch            byte
	halfSpeed        byte
	distortion       lsdj_kit_distortion
	plvibSpeed       lsdj_plvib_speed
	vibShape         lsdj_vib_shape
	vibratoDirection lsdj_vib_direction
}
