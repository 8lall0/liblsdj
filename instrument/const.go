package instrument

type lsdj_plvib_speed byte
type lsdj_vib_shape byte
type lsdj_vib_direction byte

const (
	LSDJ_PLVIB_FAST   lsdj_plvib_speed   = 0
	LSDJ_PLVIB_TICK   lsdj_plvib_speed   = 1
	LSDJ_PLVIB_STEP   lsdj_plvib_speed   = 2
	LSDJ_VIB_TRIANGLE lsdj_vib_shape     = 0
	LSDJ_VIB_SAWTOOTH lsdj_vib_shape     = 1
	LSDJ_VIB_SQUARE   lsdj_vib_shape     = 2
	LSDJ_VIB_UP       lsdj_vib_direction = 0
	LSDJ_VIB_DOWN     lsdj_vib_direction = 1
)
