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
