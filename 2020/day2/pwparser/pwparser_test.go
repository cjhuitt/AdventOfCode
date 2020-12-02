package pwparser

import "testing"

func equalSpan(a, b span) bool {
	return a.min == b.min && a.max == b.max
}

func TestParseRange(t *testing.T) {
	tests := []struct {
		input    string
		want     span
		want_err bool
	}{
		{input: "1-3", want: span{1, 3}, want_err: false},
		{input: "2-9", want: span{2, 9}, want_err: false},
	}
	for i, tc := range tests {
		got, err := parseSpan(tc.input)
		if !tc.want_err && err != nil {
			t.Errorf("Expected parse(%v) to work, received error %v (case %d)", tc.input, err, i)
		} else if !equalSpan(tc.want, got) {
			t.Errorf("Expected parse(%v) to result in %v, received %v (case %d)", tc.input, tc.want, got, i)
		}
	}
}

func equalRule(a, b rule) bool {
	return equalSpan(a.allowed, b.allowed) && a.char == b.char
}

func TestParseRule(t *testing.T) {
	tests := []struct {
		input    string
		want     rule
		want_err bool
	}{
		{input: "1-3 a", want: rule{span{1, 3}, "a"}, want_err: false},
		{input: "1-3 b", want: rule{span{1, 3}, "b"}, want_err: false},
		{input: "2-9 c", want: rule{span{2, 9}, "c"}, want_err: false},
	}
	for i, tc := range tests {
		got, err := ParseRule(tc.input)
		if !tc.want_err && err != nil {
			t.Errorf("Expected parse(%v) to work, received error %v (case %d)", tc.input, err, i)
		} else if !equalRule(tc.want, got) {
			t.Errorf("Expected parse(%v) to result in %#v, received %#v (case %d)", tc.input, tc.want, got, i)
		}
	}
}

func TestRuleMatching(t *testing.T) {
	tests := []struct {
		filter rule
		pw     string
		want   bool
	}{
		{filter: rule{span{1, 3}, "a"}, pw: "abcde", want: true},
		{filter: rule{span{1, 3}, "b"}, pw: "cdefg", want: false},
		{filter: rule{span{2, 9}, "c"}, pw: "ccccccccc", want: true},
	}
	for i, tc := range tests {
		got := tc.filter.Matches(tc.pw)
		if got != tc.want {
			t.Errorf("Expected rule{%v}.Matches(%v) to result in %v, received %v (case %d)", tc.filter, tc.pw, tc.want, got, i)
		}
	}
}
