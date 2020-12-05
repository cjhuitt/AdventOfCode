package seats

import "testing"

func TestFindRow(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{input: "", want: -1},
		{input: "FBFBBFF", want: 44},
		{input: "BFFFBBF", want: 70},
		{input: "FFFBBBF", want: 14},
		{input: "BBFFBBF", want: 102},
	}
	for i, tc := range tests {
		got := findRow(tc.input)
		if got != tc.want {
			t.Errorf("Expected findRow(%v) to result in %v, received %v (case %d)", tc.input, tc.want, got, i)
		}
	}
}

func TestFindCol(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{input: "", want: -1},
		{input: "RLR", want: 5},
		{input: "RRR", want: 7},
		{input: "RLL", want: 4},
	}
	for i, tc := range tests {
		got := findCol(tc.input)
		if got != tc.want {
			t.Errorf("Expected findCol(%v) to result in %v, received %v (case %d)", tc.input, tc.want, got, i)
		}
	}
}
