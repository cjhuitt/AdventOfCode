package decoder

import "testing"

func TestParseMask(t *testing.T) {
	tests := []struct {
		input string
		want  mask
	}{
		{input: "mask = XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X",
			want: mask{allowed: 0b111111111111111111111111111110111101,
				predefined: 0b1000000}},
	}
	for i, tc := range tests {
		got := parseMask(tc.input)
		if got != tc.want {
			t.Errorf("Expected parseMask(%v) to result in %v, received %v (case %d)", tc.input, tc.want, got, i)
		}
	}
}

func TestApplyMask(t *testing.T) {
	tests := []struct {
		m     mask
		input uint64
		want  uint64
	}{
		{m: mask{allowed: 0xFF, predefined: 0x0}, input: 10, want: 10},
		{m: mask{allowed: 0x0, predefined: 0x0}, input: 10, want: 0},
		{m: mask{allowed: 0x0F, predefined: 0xF0}, input: 0xBB, want: 0xFB},
	}
	for i, tc := range tests {
		got := tc.m.processed(tc.input)
		if got != tc.want {
			t.Errorf("Expected %#v.Processed(%d) to result in %d, received %d (case %d)", tc.m, tc.input, tc.want, got, i)
		}
	}
}

func TestStore(t *testing.T) {
	tests := []struct {
		m     mask
		loc   uint64
		input uint64
		want  uint64
	}{
		{m: mask{allowed: 0xFF, predefined: 0x0}, loc: 5, input: 10, want: 10},
		{m: mask{allowed: 0x00, predefined: 0x0}, loc: 5, input: 10, want: 0},
		{m: mask{allowed: 0xF0, predefined: 0xF}, loc: 5, input: 0xBB, want: 0xBF},
	}
	for i, tc := range tests {
		p := Program()
		p.filter = tc.m
		p.store(tc.loc, tc.input)
		got := p.mem[tc.loc]
		if got != tc.want {
			t.Errorf("Expected program{%#v}.store(%d, %d) to result in %d, received %d from %v (case %d)", tc.m, tc.input, tc.loc, tc.want, got, p.mem, i)
		}
	}
}

func TestParseStore(t *testing.T) {
	tests := []struct {
		input    string
		want_loc uint64
		want_val uint64
	}{
		{input: "mem[8] = 10", want_loc: 8, want_val: 10},
	}
	for i, tc := range tests {
		got_loc, got_val := parseStore(tc.input)
		if got_loc != tc.want_loc || got_val != tc.want_val {
			t.Errorf("Expected parseStore(%v) to result in (%d, %d), received (%d, %d) (case %d)", tc.input, tc.want_loc, tc.want_val, got_loc, got_val, i)
		}
	}
}

func TestExecuteMask(t *testing.T) {
	p := Program()
	input := "mask = XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X"
	p.Execute(input)
	want := mask{allowed: 0b111111111111111111111111111110111101,
		predefined: 0b1000000}
	if p.filter != want {
		t.Errorf("Expected Program().execute(%v) to result in mask %v, received %v", input, want, p.filter)
	}
}

func TestExecuteMemStore(t *testing.T) {
	p := Program()
	input := "mem[8] = 10"
	p.filter = mask{0xFF, 0x00}
	p.Execute(input)
	if p.mem[8] != 10 {
		t.Errorf("Expected Program().execute(%v) to result in 10 at location 8, received %v", input, p.mem)
	}
}

func TestParseLocMask(t *testing.T) {
	input := "mask = 000000000000000000000000000000X1001X"
	got := parseLocMask(input)
	if got.base != 0b010010 {
		t.Errorf("Expected parseLocMask(%v) to result in base 0b10010, received %b", input, got.base)
	} else if len(got.varies) != 2 {
		t.Errorf("Expected parseLocMask(%v) to result in 2 varying bit locations, received %d", input, len(got.varies))
	} else if got.varies[0] != 5 || got.varies[1] != 0 {
		t.Errorf("Expected parseLocMask(%v) to result in {5, 0} varying, received %v", input, got.varies)
	}
}

func TestSetMem(t *testing.T) {
	input := "mask = 0X1001X"
	m := parseLocMask(input)
	got := make(map[uint64]uint64)
	m.set(got, 10)
	if len(got) != 4 {
		t.Errorf("Expected parseLocMask(%v).set(got, 10) to set 4 locations, received %d", input, len(got))
	}
}
