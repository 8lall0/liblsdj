package liblsdj

const LSDJ_TABLE_LENGTH int = 16

type Lsdj_table_t struct {
	// The volume column of the table
	volumes [LSDJ_TABLE_LENGTH]byte
	// The transposition column of the table
	transpositions [LSDJ_TABLE_LENGTH]byte
	// The first effect command column of the table
	commands1 [LSDJ_TABLE_LENGTH]Lsdj_command_t
	// The second effect command column of the table
	commands2 [LSDJ_TABLE_LENGTH]Lsdj_command_t
}

/*
	TODO: Total refactoring in a more idiomatic way
*/

func lsdj_table_new() *Lsdj_table_t {
	var newTable Lsdj_table_t
	newTable.Clear()
	return &newTable
}

func (table Lsdj_table_t) Clear() {
	for i := 0; i < LSDJ_TABLE_LENGTH; i++ {
		table.volumes[i] = 0
		table.transpositions[i] = 0
		table.commands1[i].Clear()
		table.commands2[i].Clear()
	}
}

/*
	TODO: manage command arrays better
*/
func (dest Lsdj_table_t) Copy(source Lsdj_table_t) {
	for i := 0; i < LSDJ_TABLE_LENGTH; i++ {
		dest.volumes[i] = source.volumes[i]
		dest.transpositions[i] = source.transpositions[i]
		dest.commands1[i].CopyFrom(source.commands1[i])
		dest.commands2[i].CopyFrom(source.commands2[i])
	}
}

func (dest Lsdj_table_t) SetVolume(volume byte, index int) {
	dest.volumes[index] = volume
}

func (table Lsdj_table_t) SetVolumes(volumes [LSDJ_TABLE_LENGTH]byte) {
	dest.volumes = volumes
}

func (table Lsdj_table_t) GetVolume(index int) byte {
	return table.volumes[index]
}

func (table Lsdj_table_t) SetTransposition(transposition byte, index int) {
	table.transpositions[index] = transposition
}

func (table Lsdj_table_t) SetTranspositions(transpositions [LSDJ_TABLE_LENGTH]byte) {
	table.transpositions = transpositions
}

func (table Lsdj_table_t) GetTranspositions(index int) byte {
	return table.transpositions[index]
}

func (table Lsdj_table_t) GetCommand1(index int) *Lsdj_command_t {
	return &table.commands1[index]
}

func (table Lsdj_table_t) GetCommand2(index int) *Lsdj_command_t {
	return &table.commands2[index]
}
