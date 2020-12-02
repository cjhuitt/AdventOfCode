package pwparser

import "testing"

func TestParseRange(t *testing.T) {
	tests := []struct {
		input    string
		want_min int
		want_max int
		want_err bool
	}{
		{input: "1-3", want_min: 1, want_max: 3, want_err: false},
	}
	for i, tc := range tests {
		got_min, got_max, err := parseRange(tc.input)
		if !tc.want_err && err != nil {
			t.Errorf("Expected parse(%v) to work, received error %v (case %d)", tc.input, err, i)
		} else if tc.want_min != got_min || tc.want_max != got_max {
			t.Errorf("Expected parse(%v) to result in %d, %d, received %d, %d (case %d)", tc.input, tc.want_min, tc.want_max, got_min, got_max, i)
		}
	}
}
