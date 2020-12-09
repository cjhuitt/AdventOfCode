package bags

import "testing"

func TestParseConstraint(t *testing.T) {
	tests := []struct {
		input string
		want  bag
	}{
		{input: "", want: bag{}},
	}
	for i, tc := range tests {
		got := Parse(tc.input)
		if !got.isEqualTo(tc.want) {
			t.Errorf("Expected Parse(%v) to result in %v, received %v (case %d)", tc.input, tc.want, got, i)
		}
	}
}
