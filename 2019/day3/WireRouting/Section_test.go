package WireRouting

import "testing"

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

func TestIntersect(t *testing.T) {
	tests := []struct {
		a, b, c, d, want node
		want_found       bool
	}{
		{a: Node(1, 1), b: Node(-1, -1), c: Node(1, -1), d: Node(-1, 1), want: Node(0, 0), want_found: true},
	}
	for i, tc := range tests {
		s1 := section{tc.a, tc.b}
		s2 := section{tc.c, tc.d}
		found, got := s1.Intersect(s2)
		if found != tc.want_found {
			t.Errorf("Section(%v, %v).Intersect() want found = %v, got %v (case %d)", tc.a, tc.b, tc.want, got, i)
		} else if found && tc.want != got {
			t.Errorf("Section(%v, %v).Intersect() want %v, got %v (case %d)", tc.a, tc.b, tc.want, got, i)
		}
	}
}
