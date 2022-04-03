package liblsdj

type Song struct {
	Name              []byte
	Version           byte
	Phrases           []Phrase
	Chains            []Chain
	Tables            []Table
	Instruments       []Instrument
	Bookmarks         []Bookmark
	Waves             []Wave
	Grooves           []Groove
	ChainAssignments  ChainAssignments
	Words             []Word
	AllocationTable   AllocationTable
	SynthParams       SynthParams
	WorkHours         byte
	WorkMinutes       byte
	Tempo             byte
	Transposition     byte
	TotalDays         byte
	TotalHours        byte
	TotalMinutes      byte
	TotalTimeChecksum byte
	KeyDelay          byte
	KeyRepeat         byte
	Font              byte
	SyncMode          byte
	ColorPalette      byte
	CloneMode         byte
	FileChanged       byte
	PowerSave         byte
	PreListen         byte
	SynthOverwrites   []byte
	DrumMax           byte
	FormatVersion     byte
}
