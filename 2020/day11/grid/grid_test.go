package grid

import "testing"

func TestReadRow(t *testing.T) {
	tests := []struct {
		input string
		want  []seat
	}{
		{input: "", want: []seat{}},
		{input: "L", want: []seat{seat{'L'}}},
	}
	for i, tc := range tests {
		got := readRow(tc.input)
		if !seatSlicesEqual(got, tc.want) {
			t.Errorf("Expected readRow(%v) to result in %v, received %v (case %d)", tc.input, tc.want, got, i)
		}
	}
}
