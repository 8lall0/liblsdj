package liblsdj

type table struct {
	Command []byte
	Value   []byte
}

const (
	tableCount  = 0x20
	tableLength = 0x10
)
