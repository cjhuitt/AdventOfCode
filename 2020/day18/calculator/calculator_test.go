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

func TestFindParens(t *testing.T) {
	tests := []struct {
		input      []string
		want_start int
		want_end   int
	}{
		{input: []string{}, want_start: 0, want_end: 0},
		{input: []string{"1"}, want_start: 0, want_end: 0},
		{input: []string{"(", "1", ")"}, want_start: 0, want_end: 2},
	}
	for i, tc := range tests {
		got_start, got_end := findParens(tc.input)
		if got_start != tc.want_start || got_end != tc.want_end {
			t.Errorf("Expected findParens(%v) to result in (%v, %v), received (%v, %v) (case %d)", tc.input, tc.want_start, tc.want_end, got_start, got_end, i)
		}
	}
}

func TestProcessParens(t *testing.T) {
	tests := []struct {
		input []string
		want  []string
	}{
		{input: []string{}, want: []string{}},
		{input: []string{"1"}, want: []string{"1"}},
		{input: []string{"(", "1", ")"}, want: []string{"1"}},
		{input: []string{"(", "1", ")", "+", "(", "2", ")"}, want: []string{"1", "+", "(", "2", ")"}},
		{input: []string{"(", "(", "1", ")", ")"}, want: []string{"(", "1", ")"}},
	}
	for i, tc := range tests {
		got := processParens(tc.input)
		if len(got) != len(tc.want) {
			t.Errorf("Expected processParens(%v) to result in (%v), received (%v) (case %d)", tc.input, tc.want, got, i)
		}
	}
}

func TestCalculateWithTree(t *testing.T) {
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
		got := CalculateWithTree(tc.input)
		if got != tc.want {
			t.Errorf("Expected CalculateWithTree(%v) to result in %v, received %v (case %d)", tc.input, tc.want, got, i)
		}
	}
}
