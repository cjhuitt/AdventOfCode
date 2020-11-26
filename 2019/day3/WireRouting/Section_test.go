package WireRouting

import "testing"

func TestDefaultLength(t *testing.T) {
	got := DefaultSection().ManhattanLength()
	if got != 0 {
		t.Errorf("DefaultSection.ManhattanLength() want 0, got %d (case %d)", got, 0)
	}
}

func TestLengths(t *testing.T) {
	tests := []struct {
		a, b node
		want int
	}{
		{a: Node(0, 0), b: Node(1, 0), want: 1},
		{a: Node(0, 0), b: Node(0, 1), want: 1},
		{a: Node(1, 0), b: Node(0, 0), want: 1},
		{a: Node(0, 1), b: Node(0, 0), want: 1},
		{a: Node(1, 1), b: Node(-1, -1), want: 4},
	}
	for i, tc := range tests {
		s := section{tc.a, tc.b}
		got := s.ManhattanLength()
		if tc.want != got {
			t.Errorf("Section(%v, %v).ManhattanLength() want %d, got %d (case %d)", tc.a, tc.b, tc.want, got, i)
		}
	}
}
