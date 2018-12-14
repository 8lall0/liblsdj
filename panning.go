package lsdj

type panning byte

const (
	panNone panning = iota
	panRight
	panLeft
	panLeftRight
)
