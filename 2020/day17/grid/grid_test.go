package grid

import "testing"

func TestInitGrid(t *testing.T) {
	tests := []struct {
		input string
		want  grid
	}{
		// empty
		{input: "", want: grid{}},
	}
	for i, tc := range tests {
		got := Parse(tc.input)
		if got != tc.want {
			t.Errorf("Expected Parse(%v) to result in %v, received %v (case %d)", tc.input, tc.want, got, i)
		}
	}
}
