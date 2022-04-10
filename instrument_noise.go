package liblsdj

const (
	instrumentNoiseLengthInfinite = 0x40 //! The value of an infinite noise length
	instrumentNoiseFree           = iota
	instrumentNoiseStable
)

type NoiseInstrument struct {
	params [instrumentByteCount]byte
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
