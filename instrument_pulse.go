package liblsdj

const (
	instrumentPulseLengthInfinite = 0x40 //! The value of an infinite pulse length
	instrumentPulseWidth125       = iota
	instrumentPulseWidth25
	instrumentPulseWidth50
	instrumentPulseWidth75
)

// This part requires working at bit level, so it's gonna be a little bit more complicated...
// Remember that you can use AND with masks: fmt.Print(13 & 1) // Output: 1 -> 1

// TODO: temporary structure, probably needing something better
type PulseInstrument struct {
	params           [instrumentByteCount]byte
	instrType        byte
	env1             byte
	pu2Tsp           byte
	lengthUnlimited  byte
	length           byte
	sweep            byte
	pitchStep        byte
	pitchDrum        byte
	transpose        byte
	pitchSpeed       byte
	tableSpeed       byte
	vibratoShape     byte
	vibratoDirection byte
	tableOn          byte
	table            byte
	wave             byte
	output           byte
	cmdRate          byte
	env2             byte
	env3             byte
	fineTune         byte
}

func (p *PulseInstrument) setParams(b []byte) {
	if len(b) != instrumentByteCount {
		// do nothing
	}
	copy(p.params[:], b)
	p.instrType = b[0]
	p.env1 = b[1]
	p.pu2Tsp = b[2]
	// Qui c'è anche la lengthUnlimited
	p.length = b[3]
	p.sweep = b[4]
	// Qui ci sono in realtà molti più parametri
	p.pitchStep = b[5]
	// Qui ci sono in realtà molti più parametri
	p.table = b[6]
	// Qui ci sono in realtà molti più parametri
	p.wave = b[7]
	p.cmdRate = b[8]
	p.env2 = b[9]
	p.env3 = b[10]
	p.fineTune = b[11]
}

func (p *PulseInstrument) getParamsBytes() []byte {
	return p.params[:]
}
