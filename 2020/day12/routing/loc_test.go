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

func TestEastGenerator(t *testing.T) {
	tests := []struct {
		start loc
		steps int
		want  loc
	}{
		{start: Loc(0, 0), steps: 1, want: Loc(1, 0)},
		{start: Loc(1, 0), steps: 1, want: Loc(2, 0)},
		{start: Loc(-1, 0), steps: 1, want: Loc(0, 0)},
		{start: Loc(-10, 0), steps: 1, want: Loc(-9, 0)},
		{start: Loc(0, 0), steps: 5, want: Loc(5, 0)},
	}
	for i, tc := range tests {
		got := tc.start.East(tc.steps)
		if !(tc.want == got) {
			t.Errorf("%v.East() want %v, got %v (case %d)", tc.start, tc.want, got, i)
		}
	}
}

func TestWestGenerator(t *testing.T) {
	tests := []struct {
		start loc
		steps int
		want  loc
	}{
		{start: Loc(0, 0), steps: 1, want: Loc(-1, 0)},
		{start: Loc(1, 0), steps: 1, want: Loc(0, 0)},
		{start: Loc(-1, 0), steps: 1, want: Loc(-2, 0)},
		{start: Loc(10, 0), steps: 1, want: Loc(9, 0)},
		{start: Loc(0, 0), steps: 5, want: Loc(-5, 0)},
	}
	for i, tc := range tests {
		got := tc.start.West(tc.steps)
		if !(tc.want == got) {
			t.Errorf("%v.West() want %v, got %v (case %d)", tc.start, tc.want, got, i)
		}
	}
}

func TestNorthGenerator(t *testing.T) {
	tests := []struct {
		start loc
		steps int
		want  loc
	}{
		{start: Loc(0, 0), steps: 1, want: Loc(0, 1)},
		{start: Loc(0, 1), steps: 1, want: Loc(0, 2)},
		{start: Loc(0, -1), steps: 1, want: Loc(0, 0)},
		{start: Loc(0, -10), steps: 1, want: Loc(0, -9)},
		{start: Loc(0, 0), steps: 5, want: Loc(0, 5)},
	}
	for i, tc := range tests {
		got := tc.start.North(tc.steps)
		if !(tc.want == got) {
			t.Errorf("%v.North() want %v, got %v (case %d)", tc.start, tc.want, got, i)
		}
	}
}

func TestSouthGenerator(t *testing.T) {
	tests := []struct {
		start loc
		steps int
		want  loc
	}{
		{start: Loc(0, 0), steps: 1, want: Loc(0, -1)},
		{start: Loc(0, 1), steps: 1, want: Loc(0, 0)},
		{start: Loc(0, -1), steps: 1, want: Loc(0, -2)},
		{start: Loc(0, 10), steps: 1, want: Loc(0, 9)},
		{start: Loc(0, 0), steps: 5, want: Loc(0, -5)},
	}
	for i, tc := range tests {
		got := tc.start.South(tc.steps)
		if !(tc.want == got) {
			t.Errorf("%v.South() want %v, got %v (case %d)", tc.start, tc.want, got, i)
		}
	}
}

func TestMultiplied(t *testing.T) {
	tests := []struct {
		start loc
		mult  int
		want  loc
	}{
		{start: loc{0, 0}, mult: 10, want: loc{0, 0}},
		{start: loc{5, 6}, mult: 0, want: loc{0, 0}},
		{start: loc{5, 6}, mult: 10, want: loc{50, 60}},
		{start: loc{5, 6}, mult: -10, want: loc{-50, -60}},
	}
	for i, tc := range tests {
		got := tc.start.Multiplied(tc.mult)
		if tc.want != got {
			t.Errorf("%v.Multiplied(%v) want %v, got %v (case %d)", tc.start, tc.mult, tc.want, got, i)
		}
	}
}
