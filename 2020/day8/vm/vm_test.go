package vm

import "testing"

func TestParseOp(t *testing.T) {
	tests := []struct {
		input string
		want  opcode
	}{
		{input: "", want: opcode{"", 0}},
		{input: "nop +0", want: opcode{"nop", 0}},
		{input: "acc +1", want: opcode{"acc", 1}},
		{input: "jmp -4", want: opcode{"jmp", -4}},
	}
	for i, tc := range tests {
		got := parseOp(tc.input)
		if got.op != tc.want.op || got.val != tc.want.val {
			t.Errorf("Expected ParseOp(%v) to result in (%v), received (%v) (case %d)", tc.input, tc.want, got, i)
		}
	}
}
