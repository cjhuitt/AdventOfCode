package routing

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
		n := Loc(tc.x, tc.y)
		got := n.ManhattanLength()
		if tc.want != got {
			t.Errorf("Loc(%d, %d).ManhattanLength() want %d, got %d (case %d)", tc.x, tc.y, tc.want, got, i)
		}
	}
}

func TestRightGenerator(t *testing.T) {
	tests := []struct {
		start loc
		want  loc
	}{
		{start: Loc(0, 0), want: Loc(1, 0)},
		{start: Loc(1, 0), want: Loc(2, 0)},
		{start: Loc(-1, 0), want: Loc(0, 0)},
		{start: Loc(-10, 0), want: Loc(-9, 0)},
	}
	for i, tc := range tests {
		got := tc.start.Right()
		if !(tc.want == got) {
			t.Errorf("%v.Right() want %v, got %v (case %d)", tc.start, tc.want, got, i)
		}
	}
}

func TestLeftGenerator(t *testing.T) {
	tests := []struct {
		start loc
		want  loc
	}{
		{start: Loc(0, 0), want: Loc(-1, 0)},
		{start: Loc(1, 0), want: Loc(0, 0)},
		{start: Loc(-1, 0), want: Loc(-2, 0)},
		{start: Loc(10, 0), want: Loc(9, 0)},
	}
	for i, tc := range tests {
		got := tc.start.Left()
		if !(tc.want == got) {
			t.Errorf("%v.Left() want %v, got %v (case %d)", tc.start, tc.want, got, i)
		}
	}
}

func TestUpGenerator(t *testing.T) {
	tests := []struct {
		start loc
		want  loc
	}{
		{start: Loc(0, 0), want: Loc(0, 1)},
		{start: Loc(0, 1), want: Loc(0, 2)},
		{start: Loc(0, -1), want: Loc(0, 0)},
		{start: Loc(0, -10), want: Loc(0, -9)},
	}
	for i, tc := range tests {
		got := tc.start.Up()
		if !(tc.want == got) {
			t.Errorf("%v.Up() want %v, got %v (case %d)", tc.start, tc.want, got, i)
		}
	}
}

func TestDownGenerator(t *testing.T) {
	tests := []struct {
		start loc
		want  loc
	}{
		{start: Loc(0, 0), want: Loc(0, -1)},
		{start: Loc(0, 1), want: Loc(0, 0)},
		{start: Loc(0, -1), want: Loc(0, -2)},
		{start: Loc(0, 10), want: Loc(0, 9)},
	}
	for i, tc := range tests {
		got := tc.start.Down()
		if !(tc.want == got) {
			t.Errorf("%v.Down() want %v, got %v (case %d)", tc.start, tc.want, got, i)
		}
	}
}
