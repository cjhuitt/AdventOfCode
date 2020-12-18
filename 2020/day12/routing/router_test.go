package routing

import "testing"

func TestStepped(t *testing.T) {
	tests := []struct {
		start ship
		move  string
		want  ship
	}{
		{start: Ship(), move: "", want: ship{loc{0, 0}, loc{10, 0}}},
		{start: Ship(), move: "N1", want: ship{loc{0, 1}, loc{10, 0}}},
		{start: Ship(), move: "E1", want: ship{loc{1, 0}, loc{10, 0}}},
		{start: Ship(), move: "W1", want: ship{loc{-1, 0}, loc{10, 0}}},
		{start: Ship(), move: "S1", want: ship{loc{0, -1}, loc{10, 0}}},
		{start: Ship(), move: "R90", want: ship{loc{0, 0}, loc{0, -10}}},
		{start: Ship(), move: "R180", want: ship{loc{0, 0}, loc{-10, 0}}},
		{start: Ship(), move: "R270", want: ship{loc{0, 0}, loc{0, 10}}},
		{start: Ship(), move: "L90", want: ship{loc{0, 0}, loc{0, 10}}},
		{start: Ship(), move: "F1", want: ship{loc{10, 0}, loc{10, 0}}},
	}
	for i, tc := range tests {
		got := tc.start.Stepped(tc.move)
		if tc.want != got {
			t.Errorf("%v.Stepped(%v) want %v, got %v (case %d)", tc.start, tc.move, tc.want, got, i)
		}
	}
}
