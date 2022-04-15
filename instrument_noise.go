package liblsdj

const (
	instrumentNoiseLengthInfinite = 0x40 //! The value of an infinite noise length
	instrumentNoiseFree           = iota
	instrumentNoiseStable
)

type NoiseInstrument struct {
	params           [instrumentByteCount]byte
	instrType        byte
	env1             byte
	pitch            byte
	lengthUnlimited  byte
	length           byte
	lastEnteredNote  byte
	transpose        byte
	tableSpeed       byte
	vibratoShape     byte
	vibratoDirection byte
	tableOffOn       byte
	table            byte
	output           byte
	cmdRate          byte
	env2             byte
	env3             byte
}

func (n *NoiseInstrument) setParams(b []byte) {
	if len(b) != instrumentByteCount {
		// do nothing
	}
	copy(n.params[:], b)
}

func (n *NoiseInstrument) getParamsBytes() []byte {
	return n.params[:]
}
