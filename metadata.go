package liblsdj

type Metadata struct {
	FileNames    []byte
	FileVersions []byte
	// empty
	SRAM                 []byte
	Active               byte
	BlockAllocationTable []byte
}
