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

func TestFind(t *testing.T) {
	tests := []struct {
		input      string
		want_valid bool
		want_id    int
	}{
		{input: "", want_valid: false, want_id: -1},
		{input: "FBFBBFFRLR", want_valid: true, want_id: 357},
		{input: "BFFFBBFRRR", want_valid: true, want_id: 567},
		{input: "FFFBBBFRRR", want_valid: true, want_id: 119},
		{input: "BBFFBBFRLL", want_valid: true, want_id: 820},
	}
	for i, tc := range tests {
		got := Find(tc.input)
		if got.IsValid() != tc.want_valid {
			t.Errorf("Expected Find(%v).IsValid() to result in %v, received %v (case %d)", tc.input, tc.want_valid, got.IsValid(), i)
		} else if got.Id() != tc.want_id {
			t.Errorf("Expected Find(%v).Id() to result in %v, received %v (case %d)", tc.input, tc.want_id, got.Id(), i)
		}
	}
}
