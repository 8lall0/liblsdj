package liblsdj

type lsdj_scommand_type byte

const (
	lsdj_SCOMMAND_FREE   lsdj_scommand_type = 0
	lsdj_SCOMMAND_STABLE lsdj_scommand_type = 1
)

type noiseT struct {
	name     []byte /*[lsdj_INSTRUMENT_NAME_LENGTH]*/
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

	i.noise.sCommand = parseScommand(r.readSingle())
	i.noise.length = parseLength(r.readSingle())
	i.noise.shape = r.readSingle()
	i.automate = parseAutomate(r.readSingle())
	i.table = parseTable(r.readSingle())
	i.panning = parsePanning(r.readSingle())

	// Bytes 8-15 are empty
	r.seek(r.getCur() + 8)
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
