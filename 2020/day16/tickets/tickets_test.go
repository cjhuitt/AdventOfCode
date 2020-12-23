package tickets

import "testing"

func TestParseConstraint(t *testing.T) {
	tests := []struct {
		input string
		want  constraint
	}{
		{input: "", want: constraint{-1, -1}},
		{input: "1-3", want: constraint{min: 1, max: 3}},
	}
	for i, tc := range tests {
		got := parseConstraint(tc.input)
		if got != tc.want {
			t.Errorf("Expected parseConstraint(%v) to result in %v, received %v (case %d)", tc.input, tc.want, got, i)
		}
	}
}
