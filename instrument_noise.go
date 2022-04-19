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
	pitch            bool
	lengthParams     lengthParam
	lastEnteredNote  byte
	transpose        byte
	tableSpeed       byte
	vibratoShape     byte
	vibratoDirection byte
	tableParams      instrTableParam
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
	n.instrType = b[0]
	n.env1 = b[1]
	// TODO definisci una sintassi per i booleani che non sono acceso/spento
	n.pitch = b[2] == 1
	n.lengthParams.set(b[3])
	n.lastEnteredNote = b[4]
	n.tableParams.set(b[6])
	n.output = b[7]
	n.cmdRate = b[8]
	n.env2 = b[9]
	n.env3 = b[10]
}

func (n *NoiseInstrument) getParamsBytes() []byte {
	return n.params[:]
}
