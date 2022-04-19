package liblsdj

import "fmt"

// TODO non va bene così
func getBit(b byte, pos uint) byte {
	mask := 1 << pos
	if b != 0 {
		fmt.Println(b & byte(mask))
	}
	return b & byte(mask)
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
