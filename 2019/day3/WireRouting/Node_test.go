package WireRouting

import "testing"

func TestDefaultNodeHasZeroManhattanLength(t *testing.T) {
	got := DefaultNode().ManhattanLength()
	if got != 0 {
		t.Errorf("Default().ManhattanLength() == %d, want 0", got)
	}
}

func TestManhattanLengths(t *testing.T) {
	tests := []struct{ x, y, want int }{
		{x: 1, y: 0, want: 1},
		{x: 0, y: 1, want: 1},
		{x: 1, y: 1, want: 2},
		{x: -1, y: 0, want: 1},
		{x: 0, y: -1, want: 1},
		{x: -1, y: -1, want: 2},
		{x: -1, y: -1, want: 2},
		{x: -1, y: 1, want: 2},
		{x: 10, y: 15, want: 25},
		{x: -10, y: -15, want: 25},
	}
	for i, tc := range tests {
		n := Node(tc.x, tc.y)
		got := n.ManhattanLength()
		if tc.want != got {
			t.Errorf("Node(%d, %d).MahattenLength() want %d, got %d (case %d)", tc.x, tc.y, tc.want, got, i)
		}
	}
}

func TestEquality(t *testing.T) {
	tests := []struct {
		a, b node
		want bool
	}{
		{a: Node(0, 0), b: Node(0, 0), want: true},
		{a: Node(1, 0), b: Node(1, 0), want: true},
		{a: Node(0, 1), b: Node(0, 1), want: true},
		{a: Node(1, 1), b: Node(1, 1), want: true},
		{a: Node(1, 0), b: Node(0, 0), want: false},
		{a: Node(0, 0), b: Node(1, 0), want: false},
		{a: Node(0, 0), b: Node(0, 1), want: false},
		{a: Node(0, 0), b: Node(1, 1), want: false},
		{a: Node(1, 0), b: Node(0, 0), want: false},
		{a: Node(0, 1), b: Node(0, 0), want: false},
		{a: Node(1, 1), b: Node(0, 0), want: false},
	}
	for i, tc := range tests {
		got := tc.a.EqualTo(tc.b)
		if tc.want != got {
			t.Errorf("%v.EqualTo(%v) want %v, got %v (case %d)", tc.a, tc.b, tc.want, got, i)
		}
	}
}

func TestRightGenerator(t *testing.T) {
	tests := []struct {
		start node
		want  node
	}{
		{start: Node(0, 0), want: Node(1, 0)},
		{start: Node(1, 0), want: Node(2, 0)},
		{start: Node(-1, 0), want: Node(0, 0)},
		{start: Node(-10, 0), want: Node(-9, 0)},
	}
	for i, tc := range tests {
		got := tc.start.Right()
		if !tc.want.EqualTo(got) {
			t.Errorf("%v.Right() want %v, got %v (case %d)", tc.start, tc.want, got, i)
		}
	}
}
