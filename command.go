package liblsdj

import "io"

const (
	commandNone        = 0x00
	commandA           = 0x01
	commandC           = 0x02
	commandD           = 0x03
	commandE           = 0x04
	commandF           = 0x05
	commandG           = 0x06
	commandH           = 0x07
	commandK           = 0x08
	commandL           = 0x09
	commandM           = 0x0a
	commandO           = 0x0b
	commandP           = 0x0c
	commandR           = 0x0d
	commandS           = 0x0e
	commandT           = 0x0f
	commandV           = 0x10
	commandW           = 0x11
	commandZ           = 0x12
	commandArduinoboyN = 0x13
	commandArduinoboyX = 0x14
	commandArduinoboyQ = 0x15
	commandArduinoboyY = 0x16
)

type command struct {
	command byte
	value   byte
}

func (c *command) clear() {
	c.command = 0
	c.value = 0
}

func (c *command) write(r io.ReadSeeker) {
	c.command, _ = readByte(r)
	c.value, _ = readByte(r)
}

func (c *command) get() []byte {
	var out [2]byte
	out[0] = c.command
	out[1] = c.value

	return out[:]
}
