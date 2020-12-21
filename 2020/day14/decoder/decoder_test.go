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
