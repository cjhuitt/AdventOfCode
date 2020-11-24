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
	if got.IsEmpty() {
		t.Fatal("New([]int{1, 2, 3, 4}).IsDone() == true, want false")
	}
}
