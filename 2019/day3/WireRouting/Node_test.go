package WireRouting

import "testing"

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
			t.Errorf("Node(%d, %d).ManhattanLength() want %d, got %d (case %d)", tc.x, tc.y, tc.want, got, i)
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

func TestLeftGenerator(t *testing.T) {
	tests := []struct {
		start node
		want  node
	}{
		{start: Node(0, 0), want: Node(-1, 0)},
		{start: Node(1, 0), want: Node(0, 0)},
		{start: Node(-1, 0), want: Node(-2, 0)},
		{start: Node(10, 0), want: Node(9, 0)},
	}
	for i, tc := range tests {
		got := tc.start.Left()
		if !tc.want.EqualTo(got) {
			t.Errorf("%v.Left() want %v, got %v (case %d)", tc.start, tc.want, got, i)
		}
	}
}

func TestUpGenerator(t *testing.T) {
	tests := []struct {
		start node
		want  node
	}{
		{start: Node(0, 0), want: Node(0, 1)},
		{start: Node(0, 1), want: Node(0, 2)},
		{start: Node(0, -1), want: Node(0, 0)},
		{start: Node(0, -10), want: Node(0, -9)},
	}
	for i, tc := range tests {
		got := tc.start.Up()
		if !tc.want.EqualTo(got) {
			t.Errorf("%v.Up() want %v, got %v (case %d)", tc.start, tc.want, got, i)
		}
	}
}

func TestDownGenerator(t *testing.T) {
	tests := []struct {
		start node
		want  node
	}{
		{start: Node(0, 0), want: Node(0, -1)},
		{start: Node(0, 1), want: Node(0, 0)},
		{start: Node(0, -1), want: Node(0, -2)},
		{start: Node(0, 10), want: Node(0, 9)},
	}
	for i, tc := range tests {
		got := tc.start.Down()
		if !tc.want.EqualTo(got) {
			t.Errorf("%v.Down() want %v, got %v (case %d)", tc.start, tc.want, got, i)
		}
	}
}
