package liblsdj

const (
	instrumentPulseLengthInfinite = 0x40 //! The value of an infinite pulse length
	instrumentPulseWidth125       = iota
	instrumentPulseWidth25
	instrumentPulseWidth50
	instrumentPulseWidth75
)

type PulseInstrument struct {
	params [instrumentByteCount]byte
}

func (p *PulseInstrument) setParams(b []byte) {
	if len(b) != instrumentByteCount {
		// do nothing
	}
	copy(p.params[:], b)
}

func (p *PulseInstrument) getParamsBytes() []byte {
	return p.params[:]
}
