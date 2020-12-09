package vm

import "testing"

func TestParseOp(t *testing.T) {
	tests := []struct {
		input    string
		want_op  string
		want_val int
	}{
		{input: "", want_op: "", want_val: -1},
	}
	for i, tc := range tests {
		got_op, got_val := parseOp(tc.input)
		if got_op != tc.want_op || got_val != tc.want_val {
			t.Errorf("Expected ParseOp(%v) to result in (%v, %v), received (%v, %v) (case %d)", tc.input, tc.want_op, tc.want_val, got_op, got_val, i)
		}
	}
}
