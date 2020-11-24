package Intcode

import "testing"

func TestDefaultProgramIsEmpty(t *testing.T) {
	got := New()
	if !got.IsEmpty() {
		t.Fatal("Expected New().IsEmpty() == false, want true")
	}
}
