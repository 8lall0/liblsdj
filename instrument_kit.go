package liblsdj

const (
	instrumentKitLengthAuto = 0x0 //! The value of a InstrumentKit length set to AUTO
	instrumentKitLoopOff    = iota
	instrumentKitLoopOn
	instrumentKitLoopAttack
	instrumentKitDistortionClip = iota
	instrumentKitDistortionShape
	instrumentKitDistortionShape2
	instrumentKitDistortionWrap
)

type KitInstrument struct {
	params           [instrumentByteCount]byte
	instrType        byte
	volume           byte
	loop1Atk         byte
	loop1Speed       byte
	loop1Kit         byte
	length1          byte
	pitchStep        byte
	loop1            byte
	loop2            byte
	pitchSpeed       byte
	tableSpeed       byte
	vibratoShape     byte
	vibratoDirection byte
	tableOnOff       byte
	table            byte
	output           byte
	loop2Atk         byte
	loop2Kit         byte
	dist             byte
	length2          byte
	offset1          byte
	offset2          byte
}

func (k *KitInstrument) setParams(b []byte) {
	if len(b) != instrumentByteCount {
		// do nothing
	}
	copy(k.params[:], b)
}

func (k *KitInstrument) getParamsBytes() []byte {
	return k.params[:]
}
