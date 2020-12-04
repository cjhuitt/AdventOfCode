package passports

import "testing"

func TestParseInfoString(t *testing.T) {
	tests := []struct {
		input string
		want  passport
	}{
		// empty
		{input: "", want: passport{}},

		// Individual items
		{input: "iyr:2013", want: passport{iyr: "2013"}},

		// combination
	}
	for i, tc := range tests {
		got := Parse(tc.input)
		if got != tc.want {
			t.Errorf("Expected Parse(%v) to result in %v, received %v (case %d)", tc.input, tc.want, got, i)
		}
	}
}
