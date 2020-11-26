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
		s := section{tc.a, tc.b, 0, 0}
		got := s.ManhattanLength()
		if tc.want != got {
			t.Errorf("Section(%v, %v).ManhattanLength() want %d, got %d (case %d)", tc.a, tc.b, tc.want, got, i)
		}
	}
}

func TestIntersect(t *testing.T) {
	tests := []struct {
		a, b, c, d node
		want       intersectPoint
		want_found bool
	}{
		// parallel non-intersecting
		{a: Node(1, 0), b: Node(-1, 0), c: Node(1, 1), d: Node(-1, 1), want: intersectPoint{}, want_found: false},
		{a: Node(0, 1), b: Node(0, -1), c: Node(1, 1), d: Node(1, -1), want: intersectPoint{}, want_found: false},
		{a: Node(1, 0), b: Node(-1, 0), c: Node(2, 0), d: Node(3, 0), want: intersectPoint{}, want_found: false},
		{a: Node(0, 1), b: Node(0, -1), c: Node(0, 2), d: Node(0, 3), want: intersectPoint{}, want_found: false},

		// Identical (both directions)
		{a: Node(0, 1), b: Node(0, -1), c: Node(0, 1), d: Node(0, -1), want: intersectPoint{Node(0, 1), 0}, want_found: true},
		{a: Node(0, 1), b: Node(0, -1), c: Node(0, -1), d: Node(0, 1), want: intersectPoint{Node(0, 1), 0}, want_found: true},

		// End-point intersection
		{a: Node(0, 1), b: Node(0, -1), c: Node(0, 1), d: Node(0, 3), want: intersectPoint{Node(0, 1), 0}, want_found: true},
		{a: Node(1, 0), b: Node(-1, 0), c: Node(1, 0), d: Node(3, 0), want: intersectPoint{Node(1, 0), 0}, want_found: true},
		{a: Node(0, 1), b: Node(0, -1), c: Node(0, -1), d: Node(0, -3), want: intersectPoint{Node(0, -1), 0}, want_found: true},
		{a: Node(1, 0), b: Node(-1, 0), c: Node(-1, 0), d: Node(-3, 0), want: intersectPoint{Node(-1, 0), 0}, want_found: true},

		// Parallel complete overlap
		// Ignored for the time being
		// {a: Node(0, 1), b: Node(0, -1), c: Node(0, 3), d: Node(0, -3), want: intersectPoint{Node(0, 1), 0}, want_found: true},
		// {a: Node(0, 3), b: Node(0, -3), c: Node(0, 1), d: Node(0, -1), want: intersectPoint{Node(0, 1), 0}, want_found: true},
		// {a: Node(1, 0), b: Node(-1, 0), c: Node(3, 0), d: Node(-3, 0), want: intersectPoint{Node(1, 0), 0}, want_found: true},
		// {a: Node(3, 0), b: Node(-3, 0), c: Node(1, 0), d: Node(-1, 0), want: intersectPoint{Node(1, 0), 0}, want_found: true},

		// Non-parallel non-intersecting
		{a: Node(1, 0), b: Node(-1, 0), c: Node(3, 1), d: Node(3, -1), want: intersectPoint{}, want_found: false},

		// Intersecting
		{a: Node(1, 0), b: Node(-1, 0), c: Node(0, 1), d: Node(0, -1), want: intersectPoint{Node(0, 0), 0}, want_found: true},
	}
	for i, tc := range tests {
		s1 := section{tc.a, tc.b, 0, 0}
		s2 := section{tc.c, tc.d, 0, 0}
		found, got := s1.Intersect(s2)
		if found != tc.want_found {
			t.Errorf("%v.Intersect(%v) want found = %v, got %v (case %d)", s1, s2, tc.want_found, found, i)
		} else if found && tc.want != got {
			t.Errorf("%v.Intersect(%v) want %v, got %v (case %d)", s1, s2, tc.want, got, i)
		}
	}
}

func TestIntersectRouteLengths(t *testing.T) {
	tests := []struct {
		a, b        section
		want_length int
	}{
		{a: section{Node(1, 0), Node(-1, 0), 0, 3}, b: section{Node(0, 1), Node(0, -1), 0, 3}, want_length: 2},
		{a: section{Node(1, 0), Node(-1, 0), 10, 3}, b: section{Node(0, 1), Node(0, -1), 0, 3}, want_length: 12},
	}
	for i, tc := range tests {
		found, intersect := tc.a.Intersect(tc.b)
		if !found {
			t.Errorf("Expected %v to intersect (%v) (case %d)", tc.a, tc.b, i)
		} else {
			got := intersect.RouteLength()
			if tc.want_length != got {
				t.Errorf("%v.Intersect(%v).RouteLength() want %v, got %v (case %d)", tc.a, tc.b, tc.want_length, got, i)
			}
		}
	}
}
