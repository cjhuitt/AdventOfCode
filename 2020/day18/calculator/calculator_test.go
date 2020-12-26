package calculator

import "testing"

func TestCalculate(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{input: "1", want: 1},
		{input: "1 + 2", want: 3},
		{input: "2 * 2", want: 4},
		{input: "1 + 2 * 3", want: 9},
		{input: "1 + 2 * 3 + 4 * 5 + 6", want: 71},
		{input: "1 + (2 * 3)", want: 7},

		{input: "2 * 3 + (4 * 5)", want: 26},
		{input: "5 + (8 * 3 + 9 + 3 * 4 * 3)", want: 437},
		{input: "5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))", want: 12240},
		{input: "((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2", want: 13632},
	}
	for i, tc := range tests {
		got := Calculate(tc.input)
		if got != tc.want {
			t.Errorf("Expected Calculate(%v) to result in %v, received %v (case %d)", tc.input, tc.want, got, i)
		}
	}
}
