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

func TestParse(t *testing.T) {
	tests := []struct {
		input []string
		want  program
	}{
		{input: []string{""}, want: program{}},
		{input: []string{"nop +10"}, want: program{[]opcode{opcode{"nop", 10}}, 0, 0}},
		{input: []string{"acc -3"}, want: program{[]opcode{opcode{"acc", -3}}, 0, 0}},
		{input: []string{"jmp +5"}, want: program{[]opcode{opcode{"jmp", 5}}, 0, 0}},

		{input: []string{"nop +1", "acc +2", "jmp +3"},
			want: program{[]opcode{opcode{"nop", 1}, opcode{"acc", 2}, opcode{"jmp", 3}}, 0, 0}},
	}
	for i, tc := range tests {
		got := Parse(tc.input)
		if len(got.code) != len(tc.want.code) {
			t.Errorf("Expected Parse(%v) to result in (%v), received (%v) (case %d)", tc.input, tc.want, got, i)
		}
	}
}

func TestStep(t *testing.T) {
	tests := []struct {
		input    []string
		steps    int
		want_pos int
	}{
		{input: []string{""}, steps: 1, want_pos: -1},
		{input: []string{"nop +1", "acc +2", "jmp -2"}, steps: 1, want_pos: 1},
		{input: []string{"nop +1", "acc +2", "jmp -2"}, steps: 2, want_pos: 2},
		{input: []string{"nop +1", "acc +2", "jmp -2"}, steps: 3, want_pos: 0},
	}
	for i, tc := range tests {
		got := Parse(tc.input)
		for i := 0; i < tc.steps; i++ {
			got = got.Step()
		}
		if got.pos != tc.want_pos {
			t.Errorf("Expected Parse(%v).Step() x %d to end at %d, received %d (case %d)", tc.input, tc.steps, tc.want_pos, got.pos, i)
		}
	}
}

func TestAccumulator(t *testing.T) {
	tests := []struct {
		input    []string
		steps    int
		want_acc int
	}{
		{input: []string{""}, steps: 1, want_acc: 0},
		{input: []string{"nop +1", "acc +2", "jmp -2"}, steps: 1, want_acc: 0},
		{input: []string{"nop +1", "acc +2", "jmp -2"}, steps: 2, want_acc: 2},
		{input: []string{"nop +1", "acc +2", "jmp -2"}, steps: 3, want_acc: 2},
		{input: []string{"nop +1", "acc +2", "jmp -2"}, steps: 5, want_acc: 4},
	}
	for i, tc := range tests {
		got := Parse(tc.input)
		for i := 0; i < tc.steps; i++ {
			got = got.Step()
		}
		if got.acc != tc.want_acc {
			t.Errorf("Expected Parse(%v).Step() x %d to accumulate %d, received %d (case %d)", tc.input, tc.steps, tc.want_acc, got.acc, i)
		}
	}
}

func TestExecute(t *testing.T) {
	tests := []struct {
		input   []string
		success bool
	}{
		{input: []string{""}, success: true},
		{input: []string{"nop +1", "acc +2", "jmp -2"}, success: false},
		{input: []string{"nop +1", "acc +2", "jmp +2"}, success: true},
	}
	for i, tc := range tests {
		good, _ := Parse(tc.input).Execute()
		if good != tc.success {
			t.Errorf("Expected Parse(%v).Execute() to return %v, received %v (case %d)", tc.input, tc.success, good, i)
		}
	}
}
