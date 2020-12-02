package pwparser

import "testing"

func TestParseRange(t *testing.T) {
	tests := []struct {
		input    string
		want     span
		want_err bool
	}{
		{input: "1-3", want: span{1, 3}, want_err: false},
	}
	for i, tc := range tests {
		got, err := parseSpan(tc.input)
		if !tc.want_err && err != nil {
			t.Errorf("Expected parse(%v) to work, received error %v (case %d)", tc.input, err, i)
		} else if tc.want.min != got.min || tc.want.max != got.max {
			t.Errorf("Expected parse(%v) to result in %d, %d, received %d, %d (case %d)", tc.input, tc.want.min, tc.want.max, got.min, got.max, i)
		}
	}
}
