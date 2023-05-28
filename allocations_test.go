package liblsdj

import "testing"

func TestAllocationTable_Set(t *testing.T) {
	type fields struct {
		Phrases     [phraseAllocationsLength]byte
		Chains      [16]byte
		Instruments [64]byte
		Tables      [32]byte
	}
	type args struct {
		phrases     []byte
		chains      []byte
		instruments []byte
		tables      []byte
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:   "Wrong size parameters",
			fields: fields{},
			args: args{
				phrases:     []byte{0, 0, 0, 0},
				chains:      []byte{0, 0, 0, 0},
				instruments: []byte{0, 0, 0, 0},
				tables:      []byte{0, 0, 0, 0},
			},
			wantErr: true,
		},
		{
			name:   "Correct size parameters",
			fields: fields{},
			args: args{
				phrases:     make([]byte, 32),
				chains:      make([]byte, 16),
				instruments: make([]byte, 64),
				tables:      make([]byte, 32),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &AllocationTable{
				Phrases:     tt.fields.Phrases,
				Chains:      tt.fields.Chains,
				Instruments: tt.fields.Instruments,
				Tables:      tt.fields.Tables,
			}
			if err := a.Set(tt.args.phrases, tt.args.chains, tt.args.instruments, tt.args.tables); (err != nil) != tt.wantErr {
				t.Errorf("Set() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
