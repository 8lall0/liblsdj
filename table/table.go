package table

import "github.com/8lall0/liblsdj/command"

type Table struct {
	// The volume column of the table
	volumes [LSDJ_TABLE_LENGTH]byte
	// The transposition column of the table
	transpositions [LSDJ_TABLE_LENGTH]byte
	// The first effect command column of the table
	commands1 [LSDJ_TABLE_LENGTH]*command.Command
	// The second effect command column of the table
	commands2 [LSDJ_TABLE_LENGTH]*command.Command
}

func (table *Table) Clear() {
	for i := 0; i < LSDJ_TABLE_LENGTH; i++ {
		table.volumes[i] = 0
		table.transpositions[i] = 0
		table.commands1[i].Clear()
		table.commands2[i].Clear()
	}
}

func (dest *Table) CopyFrom(source *Table) {
	for i := 0; i < LSDJ_TABLE_LENGTH; i++ {
		dest.volumes[i] = source.volumes[i]
		dest.transpositions[i] = source.transpositions[i]
		dest.commands1[i] = source.commands1[i].Copy()
		dest.commands2[i] = source.commands2[i].Copy()
	}
}

func (dest *Table) SetVolume(volume byte, index int) {
	dest.volumes[index] = volume
}

func (table *Table) SetVolumes(volumes [LSDJ_TABLE_LENGTH]byte) {
	table.volumes = volumes
}

func (table *Table) GetVolume(index int) byte {
	return table.volumes[index]
}

func (table *Table) SetTransposition(transposition byte, index int) {
	table.transpositions[index] = transposition
}

func (table *Table) SetTranspositions(transpositions [LSDJ_TABLE_LENGTH]byte) {
	table.transpositions = transpositions
}

func (table *Table) GetTranspositions(index int) byte {
	return table.transpositions[index]
}

func (table *Table) GetCommand1(index int) *command.Command {
	return table.commands1[index]
}

func (table *Table) GetCommand2(index int) *command.Command {
	return table.commands2[index]
}
