package liblsdj

import "io"

const commandNone = 0x00
const commandA = 0x01
const commandC = 0x02
const commandD = 0x03
const commandE = 0x04
const commandF = 0x05
const commandG = 0x06
const commandH = 0x07
const commandK = 0x08
const commandL = 0x09
const commandM = 0x0a
const commandO = 0x0b
const commandP = 0x0c
const commandR = 0x0d
const commandS = 0x0e
const commandT = 0x0f
const commandV = 0x10
const commandW = 0x11
const commandZ = 0x12
const commandArduinoboyN = 0x13
const commandArduinoboyX = 0x14
const commandArduinoboyQ = 0x15
const commandArduinoboyY = 0x16

type command struct {
	command byte
	value   byte
}

func (c *command) clear() {
	c.command = 0
	c.value = 0
}

func (c *command) write(r io.ReadSeeker) {
	// TODO errori
	c.command, _ = readByte(r)
	c.value, _ = readByte(r)
}

func (c *command) get() []byte {
	var out [2]byte
	out[0] = c.command
	out[1] = c.value

	return out[:]
}
