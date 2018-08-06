package instrument

type lsdj_scommand_type byte

const (
	LSDJ_SCOMMAND_FREE   lsdj_scommand_type = 0
	LSDJ_SCOMMAND_STABLE lsdj_scommand_type = 1
)

type noiseT struct {
	length   byte // 0x40 and above = unlimited
	shape    byte
	sCommand lsdj_scommand_type
}
