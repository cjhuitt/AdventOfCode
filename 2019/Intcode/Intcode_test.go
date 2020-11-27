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
