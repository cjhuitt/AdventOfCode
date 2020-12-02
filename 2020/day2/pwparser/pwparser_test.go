package pwparser

import "testing"

func equalPair(a, b pair) bool {
	return a.first == b.first && a.second == b.second
}

func TestParseRange(t *testing.T) {
	tests := []struct {
		input    string
		want     pair
		want_err bool
	}{
		{input: "1-3", want: pair{1, 3}, want_err: false},
		{input: "2-9", want: pair{2, 9}, want_err: false},
	}
	for i, tc := range tests {
		got, err := parsePair(tc.input)
		if !tc.want_err && err != nil {
			t.Errorf("Expected parse(%v) to work, received error %v (case %d)", tc.input, err, i)
		} else if !equalPair(tc.want, got) {
			t.Errorf("Expected parse(%v) to result in %v, received %v (case %d)", tc.input, tc.want, got, i)
		}
	}
}

func equalRule(a, b rule) bool {
	return equalPair(a.allowed, b.allowed) && a.char == b.char
}

func TestParseRule(t *testing.T) {
	tests := []struct {
		input    string
		want     rule
		want_err bool
	}{
		{input: "1-3 a", want: rule{pair{1, 3}, "a"}, want_err: false},
		{input: "1-3 b", want: rule{pair{1, 3}, "b"}, want_err: false},
		{input: "2-9 c", want: rule{pair{2, 9}, "c"}, want_err: false},
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

func TestRuleMatchingType1(t *testing.T) {
	tests := []struct {
		filter rule
		pw     string
		want   bool
	}{
		{filter: rule{pair{1, 3}, "a"}, pw: "abcde", want: true},
		{filter: rule{pair{1, 3}, "b"}, pw: "cdefg", want: false},
		{filter: rule{pair{2, 9}, "c"}, pw: "ccccccccc", want: true},
	}
	for i, tc := range tests {
		got := tc.filter.MatchesType1(tc.pw)
		if got != tc.want {
			t.Errorf("Expected rule{%v}.MatchesType1(%v) to result in %v, received %v (case %d)", tc.filter, tc.pw, tc.want, got, i)
		}
	}
}

func TestRuleMatchingType2(t *testing.T) {
	tests := []struct {
		filter rule
		pw     string
		want   bool
	}{
		{filter: rule{pair{1, 3}, "a"}, pw: "abcde", want: true},
		{filter: rule{pair{1, 3}, "b"}, pw: "cdefg", want: false},
		{filter: rule{pair{2, 9}, "c"}, pw: "ccccccccc", want: false},
	}
	for i, tc := range tests {
		got := tc.filter.MatchesType2(tc.pw)
		if got != tc.want {
			t.Errorf("Expected rule{%v}.MatchesType1(%v) to result in %v, received %v (case %d)", tc.filter, tc.pw, tc.want, got, i)
		}
	}
}
