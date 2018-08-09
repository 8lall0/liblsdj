package liblsdj

type lsdj_scommand_type byte

const (
	lsdj_SCOMMAND_FREE   lsdj_scommand_type = 0
	lsdj_SCOMMAND_STABLE lsdj_scommand_type = 1
)

type noiseT struct {
	insType  int
	panning  panning
	envelope byte // envelope or byte
	table    byte // 0x20 or higher = lsdj_NO_TABLE
	automate byte
	noise    struct {
		length   byte // 0x40 and above = unlimited
		shape    byte
		sCommand lsdj_scommand_type
	}
}

func (i *noiseT) read(r *vio, ver byte) {
	i.insType = lsdj_INSTR_NOISE

	i.noise.sCommand = parseScommand(r.readByte())
	i.noise.length = parseLength(r.readByte())
	i.noise.shape = r.readByte()
	i.automate = parseAutomate(r.readByte())
	i.table = parseTable(r.readByte())
	i.panning = parsePanning(r.readByte())

	// Bytes 8-15 are empty
	r.seek(r.getCur() + 8)
}

func (i *noiseT) write(w *vio, ver byte) {
	w.writeByte(3)
	w.writeByte(i.envelope)
	w.writeByte(createScommandByte(i.noise.sCommand))
	w.writeByte(createLengthByte(i.noise.length))
	w.writeByte(i.noise.shape)
	w.writeByte(createAutomateByte(i.automate))
	w.writeByte(createTableByte(i.table))
	w.writeByte(createPanningByte(i.panning))
	w.write([]byte{0, 0, 0xD0, 0, 0, 0, 0xF3, 0})
}

func (i *noiseT) clear() {
	i.insType = lsdj_INSTR_NOISE
	i.envelope = 0xA8
	i.panning = lsdj_PAN_LEFT_RIGHT
	i.table = lsdj_NO_TABLE
	i.automate = 0

	i.noise.length = lsdj_INSTRUMENT_UNLIMITED_LENGTH
	i.noise.shape = 0xFF
	i.noise.sCommand = lsdj_SCOMMAND_FREE
}
