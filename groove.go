package liblsdj

type groove []byte

const (
	grooveCount   = 0x1f // The amount of grooves in a song
	grooveLength  = 16   // The number of steps in a groove
	grooveNoValue = 0    // The value of an empty (unused) step
)

func (g groove) SetStep() {

}

func (g groove) GetStep() {
	
}
