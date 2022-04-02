package liblsdj

// Metadata Offsets
const (
	fileNamesOffset      = 0x0000
	fileversions         = 0x100
	emptyMeta            = 0x8120
	sram                 = 0x813E
	active               = 0x8140
	blockAllocationTable = 0x8141
)

// --- Bank 0 --- //
const (
	phraseNotesOffset      = 0x0000
	bookmarksOffset        = 0x0FF0
	emptySpace0            = 0x1030
	groovesOffset          = 0x1090
	chainAssignmentsOffset = 0x1290
	tableEnvelopesOffset   = 0x1690
	wordsOffset            = 0x1890
	wordNamesOffset        = 0x1DD0
	Rb1Offset              = 0x1E78
	instrumentNamesOffset  = 0x1e7A
	emptySpace1            = 0x1FBA
)

// --- Bank 1 ---
// empty 0x2000
const (
	tableAllocationTableOffset      = 0x2020
	instrumentAllocationTableOffset = 0x2040
	chainPhrasesOffset              = 0x2080
	chainTranspositionsOffset       = 0x2880
	instrumentParamsOffset          = 0x3080
	tableTranspositionOffset        = 0x3480
	tableCommand1Offset             = 0x3680
	tableCommand1ValueOffset        = 0x3880
	tableCommand2Offset             = 0x3A80
	tableCommand2ValueOffset        = 0x3C80
	Rb2Offset                       = 0x3E80
	phraseAllocationsOffset         = 0x3E82
	chainAllocationsOffset          = 0x3EA2
	synthParamsOffset               = 0x3EB2
	workHoursOffset                 = 0x3FB2
	workMinutesOffset               = 0x3FB3
	tempoOffset                     = 0x3FB4
	transpositionOffset             = 0x3FB5
	totalDaysOffset                 = 0x3FB6
	totalHoursOffset                = 0x3FB7
	totalMinutesOffset              = 0x3FB8
	totalTimeChecksumOffset         = 0x3FB9
	keyDelayOffset                  = 0x3FBA
	keyRepeatOffset                 = 0x3FBB
	fontOffset                      = 0x3FBC
	syncModeOffset                  = 0x3FBD
	colorPaletteOffset              = 0x3FBE
	emptySpace2                     = 0x3FBF
	cloneModeOffset                 = 0x3FC0
	fileChangedOffset               = 0x3FC1
	powerSaveOffset                 = 0x3FC2
	prelistenOffset                 = 0x3FC3
	synthOverwritesOffset           = 0x3FC4
	emptySpace3                     = 0x3FC6
	drumMaxOffset                   = 0x3FD0
	emptySpace4                     = 0x3FD1
)

// --- Bank 2 --- //
const (
	phraseCommandsOffset      = 0x4000
	phraseCommandValuesOffset = 0x4FF0
	emptySpace5               = 0x5FE0
)

// --- Bank 3 --- //
const (
	wavesOffset             = 0x6000
	phraseInstrumentsOffset = 0x7000
	Rb3Offset               = 0x7FF0 // Empty 0x7FF2
	formatVersionOffset     = 0x7FFF
)
