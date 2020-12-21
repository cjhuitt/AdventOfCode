package decoder

import "testing"

func TestReadMask(t *testing.T) {
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
