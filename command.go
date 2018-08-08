package liblsdj

const (
	lsdj_COMMAND_NONE          byte = 0x00
	lsdj_COMMAND_A             byte = 0x01
	lsdj_COMMAND_C             byte = 0x02
	lsdj_COMMAND_D             byte = 0x03
	lsdj_COMMAND_E             byte = 0x04
	lsdj_COMMAND_F             byte = 0x05
	lsdj_COMMAND_G             byte = 0x06
	lsdj_COMMAND_H             byte = 0x07
	lsdj_COMMAND_K             byte = 0x08
	lsdj_COMMAND_L             byte = 0x09
	lsdj_COMMAND_M             byte = 0x0a
	lsdj_COMMAND_O             byte = 0x0b
	lsdj_COMMAND_P             byte = 0x0c
	lsdj_COMMAND_R             byte = 0x0d
	lsdj_COMMAND_S             byte = 0x0e
	lsdj_COMMAND_T             byte = 0x0f
	lsdj_COMMAND_V             byte = 0x10
	lsdj_COMMAND_W             byte = 0x11
	lsdj_COMMAND_Z             byte = 0x12
	lsdj_COMMAND_ARDUINO_BOY_N byte = 0x13
	lsdj_COMMAND_ARDUINO_BOY_X byte = 0x14
	lsdj_COMMAND_ARDUINO_BOY_Q byte = 0x15
	lsdj_COMMAND_ARDUINO_BOY_Y byte = 0x16
)

type command struct {
	command byte
	value   byte
}

func (c *command) clear() {
	c.command = 0
	c.value = 0
}

func (c *command) copy() *command {
	return &(*c)
}
