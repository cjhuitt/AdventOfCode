package passports

import "testing"

func TestParseInfoString(t *testing.T) {
	tests := []struct {
		input string
		want  passport
	}{
		{input: "", want: passport{}},
	}
	for i, tc := range tests {
		got := Parse(tc.input)
		if got != tc.want {
			t.Errorf("Expected Parse(%v) to result in %v, received %v (case %d)", tc.input, tc.want, got, i)
		}
	}
}
