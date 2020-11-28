package Intcode

import "testing"

func TestDefaultProgramIsEmpty(t *testing.T) {
	got := Default()
	if !got.IsEmpty() {
		t.Fatal("Default().IsEmpty() == false, want true")
	}
}

func TestDefaultProgramIsDone(t *testing.T) {
	got := Default()
	if !got.IsDone() {
		t.Fatal("Default().IsDone() == false, want true")
	}
}

func TestProgramWithDataIsNotEmpty(t *testing.T) {
	got := New([]int{1, 2, 3, 4})
	if got.IsEmpty() {
		t.Fatal("New([]int{1, 2, 3, 4}).IsEmpty() == true, want false")
	}
}

func TestNewProgramWithDataIsNotDone(t *testing.T) {
	got := New([]int{1, 2, 3, 4})
	if got.IsDone() {
		t.Fatal("New([]int{1, 2, 3, 4}).IsDone() == true, want false")
	}
}

func TestNewProgramWithTerminateOpcode(t *testing.T) {
	got := New([]int{99, 2, 3, 4})
	if !got.IsDone() {
		t.Fatal("New([]int{99, 2, 3, 4}).IsDone() == false, want true")
	}
}

func Equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func TestAddOpCode(t *testing.T) {
	tests := []struct {
		input []int
		want  []int
	}{
		{input: []int{1, 4, 5, 0, 2, 3}, want: []int{5, 4, 5, 0, 2, 3}}, // store first
		{input: []int{1, 4, 5, 1, 2, 3}, want: []int{1, 5, 5, 1, 2, 3}}, // store second
		{input: []int{1, 4, 5, 2, 3, 3}, want: []int{1, 4, 6, 2, 3, 3}}, // store third
		{input: []int{1, 4, 5, 3, 2, 3}, want: []int{1, 4, 5, 5, 2, 3}}, // store fourth

		{input: []int{1, 6, 3, 0, 2, 3}, want: []int{1, 6, 3, 0, 2, 3}},   // read out of bounds on end
		{input: []int{1, -1, 3, 0, 2, 3}, want: []int{1, -1, 3, 0, 2, 3}}, // read out of bounds on begin

		{input: []int{1, 3, 6, 0, 2, 3}, want: []int{1, 3, 6, 0, 2, 3}},   // read out of bounds on end
		{input: []int{1, 3, -1, 0, 2, 3}, want: []int{1, 3, -1, 0, 2, 3}}, // read out of bounds on begin

		{input: []int{1, 3, 3, 6, 2, 3}, want: []int{1, 3, 3, 6, 2, 3}},   // store out of bounds on end
		{input: []int{1, 3, 3, -1, 2, 3}, want: []int{1, 3, 3, -1, 2, 3}}, // store out of bounds on begin

		{input: []int{1, 3, 3}, want: []int{1, 3, 3}}, // short input test

		// Parameter modes
		{input: []int{1101, 100, 3, 1, 99}, want: []int{1101, 103, 3, 1, 99}},
		{input: []int{1001, 1, 3, 1, 99}, want: []int{1001, 4, 3, 1, 99}},
		{input: []int{101, 100, 3, 1, 99}, want: []int{101, 101, 3, 1, 99}},
		{input: []int{1101, 100, -5, 1, 99}, want: []int{1101, 95, -5, 1, 99}},
		{input: []int{11101, 100, -5, 1, 99}, want: []int{11101, 100, -5, 1, 99}}, //invalid parameterized store
	}
	for i, tc := range tests {
		p := New(tc.input)
		got := p.Step().Data()
		if !Equal(tc.want, got) {
			t.Errorf("Expected stepping %v to be %v, got %v (case %d)", tc.input, tc.want, got, i)
		}
	}
}

func TestMultOpCode(t *testing.T) {
	tests := []struct {
		input []int
		want  []int
	}{
		{input: []int{2, 4, 5, 0, 3, 3}, want: []int{9, 4, 5, 0, 3, 3}}, // store first
		{input: []int{2, 4, 5, 1, 3, 3}, want: []int{2, 9, 5, 1, 3, 3}}, // store second
		{input: []int{2, 4, 5, 2, 3, 3}, want: []int{2, 4, 9, 2, 3, 3}}, // store third
		{input: []int{2, 4, 5, 3, 3, 3}, want: []int{2, 4, 5, 9, 3, 3}}, // store fourth

		{input: []int{2, 6, 3, 0, 3, 3}, want: []int{2, 6, 3, 0, 3, 3}},   // read out of bounds on end
		{input: []int{2, -1, 3, 0, 3, 3}, want: []int{2, -1, 3, 0, 3, 3}}, // read out of bounds on begin

		{input: []int{2, 3, 6, 0, 3, 3}, want: []int{2, 3, 6, 0, 3, 3}},   // read out of bounds on end
		{input: []int{2, 3, -1, 0, 3, 3}, want: []int{2, 3, -1, 0, 3, 3}}, // read out of bounds on begin

		{input: []int{2, 3, 3, 6, 3, 3}, want: []int{2, 3, 3, 6, 3, 3}},   // store out of bounds on end
		{input: []int{2, 3, 3, -1, 3, 3}, want: []int{2, 3, 3, -1, 3, 3}}, // store out of bounds on begin

		{input: []int{2, 3, 3}, want: []int{2, 3, 3}}, // short input test

		// Parameter modes
		{input: []int{1102, 100, 3, 1, 99}, want: []int{1102, 300, 3, 1, 99}},
		{input: []int{1002, 1, 3, 1, 99}, want: []int{1002, 3, 3, 1, 99}},
		{input: []int{102, 100, 1, 1, 99}, want: []int{102, 10000, 1, 1, 99}},
		{input: []int{1102, 100, -5, 1, 99}, want: []int{1102, -500, -5, 1, 99}},
		{input: []int{11102, 100, -5, 1, 99}, want: []int{11102, 100, -5, 1, 99}}, //invalid parameterized store
	}
	for i, tc := range tests {
		p := New(tc.input)
		got := p.Step().Data()
		if !Equal(tc.want, got) {
			t.Errorf("Expected stepping %v to be %v, got %v (case %d)", tc.input, tc.want, got, i)
		}
	}
}

func TestSamplePrograms(t *testing.T) {
	tests := []struct {
		input []int
		want  []int
	}{
		{input: []int{1, 0, 0, 0, 99}, want: []int{2, 0, 0, 0, 99}},
		{input: []int{2, 3, 0, 3, 99}, want: []int{2, 3, 0, 6, 99}},
		{input: []int{2, 4, 4, 5, 99, 0}, want: []int{2, 4, 4, 5, 99, 9801}},
		{input: []int{1, 1, 1, 4, 99, 5, 6, 0, 99}, want: []int{30, 1, 1, 4, 2, 5, 6, 0, 99}},
		{input: []int{1, 9, 10, 3, 2, 3, 11, 0, 99, 30, 40, 50}, want: []int{3500, 9, 10, 70, 2, 3, 11, 0, 99, 30, 40, 50}},
		{input: []int{0, 0, 0, 0, 99}, want: []int{0, 0, 0, 0, 99}},
	}
	for i, tc := range tests {
		p := New(tc.input)
		got := p.Execute().Data()
		if !Equal(tc.want, got) {
			t.Errorf("Expected stepping %v to be %v, got %v (case %d)", tc.input, tc.want, got, i)
		}
	}
}

func TestOutputOpCode(t *testing.T) {
	tests := []struct {
		input    []int
		expected bool
		want     int
	}{
		{input: []int{4, 1, 99}, expected: true, want: 1},
		{input: []int{4, 0, 99}, expected: true, want: 4},
		{input: []int{4, 2, 99}, expected: true, want: 99},

		{input: []int{4, -1, 99}, expected: false, want: 0}, // out of bounds
		{input: []int{4, 3, 99}, expected: false, want: 0},  // out of bounds

		// immediate mode
		{input: []int{104, 25, 99}, expected: true, want: 25},
	}
	for i, tc := range tests {
		p := New(tc.input)
		got := p.Step().Output()
		if got == nil && tc.expected {
			t.Errorf("Expected stepping %v to have output, it doesn't (case %d)", tc.input, i)
		} else if got != nil && !tc.expected {
			t.Errorf("Expected stepping %v to not have output, it does (case %d)", tc.input, i)
		} else if got != nil && tc.want != *got {
			t.Errorf("Expected stepping %v to output %v, got %v (case %d)", tc.input, tc.want, *got, i)
		}
	}
}

func TestPausesAtOutput(t *testing.T) {
	tests := []struct {
		input  []int
		paused bool
		ended  bool
	}{
		{input: []int{4, 1, 2, 0, 1, 2, 99}, paused: true, ended: false},
		{input: []int{2, 0, 1, 2, 4, 1, 2, 0, 1, 2, 99}, paused: true, ended: false},
		{input: []int{2, 0, 1, 2, 2, 0, 1, 2, 99}, paused: false, ended: true},

		// if terminating execution, don't count as "paused"
		{input: []int{4, 1, 99}, paused: false, ended: true},
	}
	for i, tc := range tests {
		p := New(tc.input)
		got := p.Execute()
		if got.IsPaused() != tc.paused {
			t.Errorf("Expected executing %v to not end, it does (case %d)", tc.input, i)
		} else if got.IsDone() != tc.ended {
			t.Errorf("Expected executing %v to not end, it does (case %d)", tc.input, i)
		}
	}
}

func TestInputOpCode(t *testing.T) {
	tests := []struct {
		program []int
		input   int
		want    []int
	}{
		{program: []int{3, 1, 99}, input: 5, want: []int{3, 5, 99}},
		{program: []int{3, 0, 99}, input: 5, want: []int{5, 0, 99}},
		{program: []int{3, 2, 99}, input: 5, want: []int{3, 2, 5}},

		{program: []int{3, -1, 99}, input: 5, want: []int{3, -1, 99}},   // out of bounds
		{program: []int{3, 4, 99}, input: 5, want: []int{3, 4, 99}},     // out of bounds
		{program: []int{103, 2, 99}, input: 5, want: []int{103, 2, 99}}, // Invalid parameterized store
	}
	for i, tc := range tests {
		p := New(tc.program)
		p = p.Step().WithInput(&tc.input)
		got := p.Step().Data()
		if !Equal(tc.want, got) {
			t.Errorf("Expected stepping %v to be %v, got %v (case %d)", tc.program, tc.want, got, i)
		}
	}
}

func TestPausesAtInput(t *testing.T) {
	tests := []struct {
		program []int
		paused  bool
		ended   bool
	}{
		{program: []int{3, 1, 2, 0, 1, 2, 99}, paused: true, ended: false},
		{program: []int{2, 0, 1, 2, 3, 1, 2, 0, 1, 2, 99}, paused: true, ended: false},
		{program: []int{2, 0, 1, 2, 2, 0, 1, 2, 99}, paused: false, ended: true},

		// if terminating execution but input hasn't been received yet, be paused
		{program: []int{3, 1, 99}, paused: true, ended: false},
	}
	for i, tc := range tests {
		p := New(tc.program)
		got := p.Execute()
		if got.IsPaused() != tc.paused {
			t.Errorf("Expected executing %v to not end, it does (case %d)", tc.program, i)
		} else if got.IsDone() != tc.ended {
			t.Errorf("Expected executing %v to not end, it does (case %d)", tc.program, i)
		}
	}
}

func TestJumps(t *testing.T) {
	tests := []struct {
		program []int
		want    []int
		finish  bool
	}{
		{program: []int{5, 1, 8, 1101, 10, 20, 0, 99, 7}, want: []int{5, 1, 8, 1101, 10, 20, 0, 99, 7}, finish: true},
		{program: []int{5, 6, 8, 1101, 10, 20, 0, 99, 7}, want: []int{30, 6, 8, 1101, 10, 20, 0, 99, 7}, finish: true},
		{program: []int{6, 1, 8, 1101, 10, 20, 0, 99, 7}, want: []int{30, 1, 8, 1101, 10, 20, 0, 99, 7}, finish: true},
		{program: []int{6, 6, 8, 1101, 10, 20, 0, 99, 7}, want: []int{6, 6, 8, 1101, 10, 20, 0, 99, 7}, finish: true},

		// Out of range locations
		{program: []int{5, -1, 8, 1101, 10, 20, 0, 99, 7}, want: []int{5, -1, 8, 1101, 10, 20, 0, 99, 7}, finish: false},
		{program: []int{5, 9, 8, 1101, 10, 20, 0, 99, 7}, want: []int{5, 9, 8, 1101, 10, 20, 0, 99, 7}, finish: false},
		{program: []int{5, 1, -1, 1101, 10, 20, 0, 99, 7}, want: []int{5, 1, -1, 1101, 10, 20, 0, 99, 7}, finish: false},
		{program: []int{5, 1, 9, 1101, 10, 20, 0, 99, 7}, want: []int{5, 1, 9, 1101, 10, 20, 0, 99, 7}, finish: false},
		{program: []int{6, -1, 8, 1101, 10, 20, 0, 99, 7}, want: []int{6, -1, 8, 1101, 10, 20, 0, 99, 7}, finish: false},
		{program: []int{6, 9, 8, 1101, 10, 20, 0, 99, 7}, want: []int{6, 9, 8, 1101, 10, 20, 0, 99, 7}, finish: false},
		{program: []int{6, 1, -1, 1101, 10, 20, 0, 99, 7}, want: []int{6, 1, -1, 1101, 10, 20, 0, 99, 7}, finish: false},
		{program: []int{6, 1, 9, 1101, 10, 20, 0, 99, 7}, want: []int{6, 1, 9, 1101, 10, 20, 0, 99, 7}, finish: false},

		// immediate mode
		{program: []int{105, 0, 8, 1101, 10, 20, 0, 99, 7}, want: []int{30, 0, 8, 1101, 10, 20, 0, 99, 7}, finish: true},
		{program: []int{1105, 1, 7, 1101, 10, 20, 0, 99, 7}, want: []int{1105, 1, 7, 1101, 10, 20, 0, 99, 7}, finish: true},
		{program: []int{1005, 6, 7, 1101, 10, 20, 0, 99, 7}, want: []int{30, 6, 7, 1101, 10, 20, 0, 99, 7}, finish: true},
		{program: []int{106, 0, 8, 1101, 10, 20, 0, 99, 7}, want: []int{106, 0, 8, 1101, 10, 20, 0, 99, 7}, finish: true},
		{program: []int{1106, 1, 7, 1101, 10, 20, 0, 99, 7}, want: []int{30, 1, 7, 1101, 10, 20, 0, 99, 7}, finish: true},
		{program: []int{1006, 6, 7, 1101, 10, 20, 0, 99, 7}, want: []int{1006, 6, 7, 1101, 10, 20, 0, 99, 7}, finish: true},
	}
	for i, tc := range tests {
		p := New(tc.program)
		p = p.Execute()
		got := p.Data()
		if tc.finish && p.IsErrored() {
			t.Errorf("Expected executing %v to succeed, got %v (case %d)", tc.program, p, i)
		} else if !tc.finish && !p.IsErrored() {
			t.Errorf("Expected executing %v to error, got %v (case %d)", tc.program, p, i)
		} else if !Equal(tc.want, got) {
			t.Errorf("Expected executing %v to be %v, got %v (case %d)", tc.program, tc.want, got, i)
		}
	}
}

func TestComparisons(t *testing.T) {
	tests := []struct {
		program []int
		want    []int
		finish  bool
	}{
		{program: []int{7, 4, 5, 0, 99, 98}, want: []int{0, 4, 5, 0, 99, 98}, finish: true},
		{program: []int{7, 5, 4, 0, 99, 98}, want: []int{1, 5, 4, 0, 99, 98}, finish: true},
		{program: []int{8, 4, 5, 0, 99, 99}, want: []int{1, 4, 5, 0, 99, 99}, finish: true},
		{program: []int{8, 4, 5, 0, 99, 98}, want: []int{0, 4, 5, 0, 99, 98}, finish: true},

		// Out of range locations
		{program: []int{7, -1, 5, 0, 99, 98}, want: []int{7, -1, 5, 0, 99, 98}, finish: false},
		{program: []int{7, 6, 5, 0, 99, 98}, want: []int{7, 6, 5, 0, 99, 98}, finish: false},
		{program: []int{7, 4, -1, 0, 99, 98}, want: []int{7, 4, -1, 0, 99, 98}, finish: false},
		{program: []int{7, 4, 6, 0, 99, 98}, want: []int{7, 4, 6, 0, 99, 98}, finish: false},
		{program: []int{8, -1, 5, 0, 99, 98}, want: []int{8, -1, 5, 0, 99, 98}, finish: false},
		{program: []int{8, 6, 5, 0, 99, 98}, want: []int{8, 6, 5, 0, 99, 98}, finish: false},
		{program: []int{8, 4, -1, 0, 99, 98}, want: []int{8, 4, -1, 0, 99, 98}, finish: false},
		{program: []int{8, 4, 6, 0, 99, 98}, want: []int{8, 4, 6, 0, 99, 98}, finish: false},

		// immediate mode
		{program: []int{1107, 4, 5, 0, 99, 98}, want: []int{1, 4, 5, 0, 99, 98}, finish: true},
		{program: []int{107, 4, 5, 0, 99, 98}, want: []int{1, 4, 5, 0, 99, 98}, finish: true},
		{program: []int{1007, 4, 5, 0, 99, 98}, want: []int{0, 4, 5, 0, 99, 98}, finish: true},
		{program: []int{1108, 4, 5, 0, 99, 98}, want: []int{0, 4, 5, 0, 99, 98}, finish: true},
		{program: []int{108, 4, 5, 0, 99, 98}, want: []int{0, 4, 5, 0, 99, 98}, finish: true},
		{program: []int{1008, 4, 5, 0, 99, 98}, want: []int{0, 4, 5, 0, 99, 98}, finish: true},
	}
	for i, tc := range tests {
		p := New(tc.program)
		p = p.Execute()
		got := p.Data()
		if tc.finish && p.IsErrored() {
			t.Errorf("Expected executing %v to succeed, got %v (case %d)", tc.program, p, i)
		} else if !tc.finish && !p.IsErrored() {
			t.Errorf("Expected executing %v to error, got %v (case %d)", tc.program, p, i)
		} else if !Equal(tc.want, got) {
			t.Errorf("Expected executing %v to be %v, got %v (case %d)", tc.program, tc.want, p, i)
		}
	}
}
