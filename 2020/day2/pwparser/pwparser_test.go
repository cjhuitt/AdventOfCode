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
		{input: "1-3: a", want: rule{span{1, 3}, "a"}, want_err: false},
		{input: "1-3: b", want: rule{span{1, 3}, "b"}, want_err: false},
		{input: "2-9: c", want: rule{span{2, 9}, "c"}, want_err: false},
	}
	for i, tc := range tests {
		got, err := parseRule(tc.input)
		if !tc.want_err && err != nil {
			t.Errorf("Expected parse(%v) to work, received error %v (case %d)", tc.input, err, i)
		} else if !equalRule(tc.want, got) {
			t.Errorf("Expected parse(%v) to result in %#v, received %#v (case %d)", tc.input, tc.want, got, i)
		}
	}
}
