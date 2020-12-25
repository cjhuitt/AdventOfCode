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
	}
	for i, tc := range tests {
		got := Calculate(tc.input)
		if got != tc.want {
			t.Errorf("Expected Calculate(%v) to result in %v, received %v (case %d)", tc.input, tc.want, got, i)
		}
	}
}
