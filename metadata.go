package liblsdj

type Metadata struct {
	FileNames    [0x20][8]byte
	FileVersions [0x20]byte
	// empty
	SRAM                 []byte
	ActiveFile           byte
	BlockAllocationTable []byte
}

func Init() {

}

func (m *Metadata) setFileNames(b []byte) {
	for i, v := range m.FileNames {
		copy(v[:], b[i*8:(i+1)*8])
	}
}

func (m *Metadata) setFileVersions(b []byte) {
	for i, v := range b {
		m.FileVersions[i] = v
	}
}
