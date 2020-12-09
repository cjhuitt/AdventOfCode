package bags

import "testing"

func TestParseConstraint(t *testing.T) {
	tests := []struct {
		input string
		want  bag
	}{
		{input: "", want: bag{}},
		{input: "faded blue bags contain no other bags", want: bag{style: "faded blue", contents: map[string]int{}}},
	}
	for i, tc := range tests {
		got := Parse(tc.input)
		if !got.isEqualTo(tc.want) {
			t.Errorf("Expected Parse(%v) to result in %v, received %v (case %d)", tc.input, tc.want, got, i)
		}
	}
}
