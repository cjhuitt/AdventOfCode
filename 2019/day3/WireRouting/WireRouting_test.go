package WireRouting

import "testing"

func TestDefaultPathHasNoLength(t *testing.T) {
	got := Default().Length()
	if got != 0 {
		t.Errorf("Default().Length() == %d, want 0", got)
	}
}

func TestBuildsPathFromInstructions(t *testing.T) {
	tests := []struct {
		path     string
		length   int
		contains []node
	}{
		{path: "", length: 0, contains: []node{}},
	}
	for i, tc := range tests {
		r := Route(tc.path)
		got := r.Length()
		if got != tc.length {
			t.Errorf("Route(%v).Length) want %d, got %d (case %d)", tc.path, tc.length, got, i)
		}
		for _, n := range tc.contains {
			if !r.Contains(n) {
				t.Errorf("Route(%v).Contains(%v) want true, got false (case %d)", tc.path, n, i)
			}
		}
	}
}
