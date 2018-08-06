package command

type Command struct {
	command byte
	value   byte
}

func (c *Command) Clear() {
	c.command = 0
	c.value = 0
}

func (src *Command) Copy() *Command {
	return &(*src)
}
