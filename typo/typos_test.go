package typo

import (
	"testing"
)

type fields struct {
	Original string
	Typos    []string
	Excluded []string
}

type typotest struct {
	name         string
	fields       fields
	resultLength int
}

func Test_typos_AddTypo(t1 *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name         string
		fields       fields
		args         args
		resultLength int
	}{
		{"DontAddSameAsOriginal", fields{Original: "ddog"}, args{s: "ddog"}, 0},
		{"DontAddExisting", fields{Original: "cat", Typos: []string{"ccat"}}, args{s: "ccat"}, 1},
		{"DontAddExcluded", fields{Original: "fastapi", Typos: []string{"ffastapi"}, Excluded: []string{"faastapi"}}, args{s: "faastapi"}, 1},
		{"DoAddNew", fields{Original: "dog"}, args{s: "ddog"}, 1},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &typos{
				Original: tt.fields.Original,
				Typos:    tt.fields.Typos,
				Excluded: tt.fields.Excluded,
			}
			t.AddTypo(tt.args.s)
			if tt.resultLength != len(t.Typos) {
				t1.Fatalf("Expected Typos length '%d' got length '%d'", len(t.Typos), tt.resultLength)
			}
		})
	}
}

func Test_typos_AllTypos(t1 *testing.T) {
	tests := []typotest{
		{"DoAll", fields{Original: "sly", Typos: []string{}}, 215},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &typos{
				Original: tt.fields.Original,
				Typos:    tt.fields.Typos,
			}
			t.AllTypos()
			if tt.resultLength != len(t.Typos) {
				t1.Fatalf("Expected Typos length '%d' got length '%d'", len(t.Typos), tt.resultLength)
			}
		})
	}
}

func Test_typos_DoubleLetter(t1 *testing.T) {
	tests := []typotest{
		{"DontDoubleForwardSlash", fields{Original: "cat/jam", Typos: []string{}}, 6},
		{"DoDouble", fields{Original: "groundhog", Typos: []string{}}, 9},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &typos{
				Original: tt.fields.Original,
				Typos:    tt.fields.Typos,
			}
			t.DoubleLetter()
			if tt.resultLength != len(t.Typos) {
				t1.Fatalf("Expected Typos length '%d' got length '%d'", len(t.Typos), tt.resultLength)
			}
		})
	}
}

func Test_typos_InsertedKey(t1 *testing.T) {
	tests := []typotest{
		{"DoInsert", fields{Original: "dog", Typos: []string{}}, 106},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &typos{
				Original: tt.fields.Original,
				Typos:    tt.fields.Typos,
			}
			t.InsertedKey()
			if tt.resultLength != len(t.Typos) {
				t1.Fatalf("Expected Typos length '%d' got length '%d'", len(t.Typos), tt.resultLength)
			}
		})
	}
}

func Test_typos_ReverseLetter(t1 *testing.T) {
	tests := []typotest{
		{"DontReverseLessThanTwo", fields{Original: "b", Typos: []string{}}, 0},
		{"DoReverseShortest", fields{Original: "bo", Typos: []string{}}, 1},
		{"DoReverseShort", fields{Original: "bol", Typos: []string{}}, 1},
		{"DoReverse", fields{Original: "bold", Typos: []string{}}, 2},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &typos{
				Original: tt.fields.Original,
				Typos:    tt.fields.Typos,
			}
			t.ReverseLetter()
			if tt.resultLength != len(t.Typos) {
				t1.Fatalf("Expected Typos length '%d' got length '%d'", len(t.Typos), tt.resultLength)
			}
		})
	}
}

func Test_typos_SkipLetter(t1 *testing.T) {
	tests := []typotest{
		{"DoSkip", fields{Original: "camp", Typos: []string{}}, 4},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &typos{
				Original: tt.fields.Original,
				Typos:    tt.fields.Typos,
			}
			t.SkipLetter()
			if tt.resultLength != len(t.Typos) {
				t1.Fatalf("Expected Typos length '%d' got length '%d'", len(t.Typos), tt.resultLength)
			}
		})
	}
}

func Test_typos_WrongKey(t1 *testing.T) {
	tests := []typotest{
		{"DoWrongKey", fields{Original: "log", Typos: []string{}}, 105},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &typos{
				Original: tt.fields.Original,
				Typos:    tt.fields.Typos,
			}
			t.WrongKey()
			if tt.resultLength != len(t.Typos) {
				t1.Fatalf("Expected Typos length '%d' got length '%d'", len(t.Typos), tt.resultLength)
			}
		})
	}
}

func Test_typos_WrongVowel(t1 *testing.T) {
	tests := []typotest{
		{"DoWrongVowelsAY", fields{Original: "scary", Typos: []string{}}, 10},
		{"DoWrongVowelsEI", fields{Original: "dire", Typos: []string{}}, 10},
		{"DoWrongVowelsYOU", fields{Original: "your", Typos: []string{}}, 15},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &typos{
				Original: tt.fields.Original,
				Typos:    tt.fields.Typos,
			}
			t.WrongVowel()
			if tt.resultLength != len(t.Typos) {
				t1.Fatalf("Expected Typos length '%d' got length '%d'", len(t.Typos), tt.resultLength)
			}
		})
	}
}
