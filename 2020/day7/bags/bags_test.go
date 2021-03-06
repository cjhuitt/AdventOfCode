package bags

import "testing"

func TestParseConstraint(t *testing.T) {
	tests := []struct {
		input string
		want  parsedBag
	}{
		{input: "", want: parsedBag{}},
		{input: "faded blue bags contain no other bags.",
			want: parsedBag{style: "faded blue", contents: map[string]int{}}},
		{input: "bright white bags contain 1 shiny gold bag.",
			want: parsedBag{style: "bright white", contents: map[string]int{"shiny gold": 1}}},
		{input: "shiny gold bags contain 1 dark olive bag, 2 vibrant plum bags.",
			want: parsedBag{style: "shiny gold", contents: map[string]int{"dark olive": 1, "vibrant plum": 2}}},
	}
	for i, tc := range tests {
		got := ParseConstraint(tc.input)
		if !got.isEqualTo(tc.want) {
			t.Errorf("Expected ParseConstraint(%v) to result in %v, received %v (case %d)", tc.input, tc.want, got, i)
		}
	}
}
