package WireRouting

import "testing"

func TestDefaultPathHasNoLength(t *testing.T) {
	got := Default().Length()
	if got != 0 {
		t.Errorf("Default().Length() == %d, want 0", got)
	}
}
