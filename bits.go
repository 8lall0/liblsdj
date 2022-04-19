package liblsdj

import "math"

func getBit(b byte, pos uint) byte {
	pow := math.Pow(2, float64(pos))
	return b & byte(pow)
}

func setBit(n byte, pos uint) byte {
	n |= 1 << pos
	return n
}

func clearBit(n byte, pos uint) byte {
	mask := ^(1 << pos)
	n &= byte(mask)
	return n
}

func hasBit(n byte, pos uint) bool {
	val := n & (1 << pos)
	return val > 0
}
