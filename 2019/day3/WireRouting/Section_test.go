package WireRouting

import "testing"

func TestDefaultLength(t *testing.T) {
	got := DefaultSection().Length()
	if got != 0 {
		t.Errorf("DefaultSection.Length() want 0, got %d (case %d)", got, 0)
	}
}
