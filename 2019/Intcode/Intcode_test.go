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
		{input: []int{2, 3, 3, 0}, want: []int{9, 3, 3, 0}},   // store first
		{input: []int{2, 3, 3, 1}, want: []int{2, 9, 3, 1}},   // store second
		{input: []int{2, 3, 3, 2}, want: []int{2, 3, 9, 2}},   // store third
		{input: []int{2, 3, 3, 3}, want: []int{2, 3, 3, 9}},   // store fourth
		{input: []int{2, 3, 3, 4}, want: []int{2, 3, 3, 4}},   // store out of bounds on end
		{input: []int{2, 3, 3, -1}, want: []int{2, 3, 3, -1}}, // store out of bounds on begin
		{input: []int{2, 3, 3}, want: []int{2, 3, 3}},         // short input test
	}
	for _, tc := range tests {
		p := New(tc.input)
		got := p.Step().Data()
		if !Equal(tc.want, got) {
			t.Errorf("Expected stepping %v to be %v, got %v", tc.input, tc.want, got)
		}
	}
}
