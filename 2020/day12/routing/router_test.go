package routing

import "testing"

func TestStepped(t *testing.T) {
	tests := []struct {
		move string
		want ship
	}{
		{move: "", want: ship{loc{0, 0}, loc{10, 0}, 'E'}},
		{move: "N1", want: ship{loc{0, 1}, loc{10, 0}, 'E'}},
		{move: "E1", want: ship{loc{1, 0}, loc{10, 0}, 'E'}},
		{move: "W1", want: ship{loc{-1, 0}, loc{10, 0}, 'E'}},
		{move: "S1", want: ship{loc{0, -1}, loc{10, 0}, 'E'}},
		{move: "R90", want: ship{loc{0, 0}, loc{0, -10}, 'S'}},
		{move: "R180", want: ship{loc{0, 0}, loc{-10, 0}, 'W'}},
		{move: "R270", want: ship{loc{0, 0}, loc{0, 10}, 'N'}},
		{move: "L90", want: ship{loc{0, 0}, loc{0, 10}, 'N'}},
		{move: "F1", want: ship{loc{1, 0}, loc{10, 0}, 'E'}},
	}
	for i, tc := range tests {
		got := Ship().Stepped(tc.move)
		if tc.want != got {
			t.Errorf("Ship().Stepped(%v) want %v, got %v (case %d)", tc.move, tc.want, got, i)
		}
	}
}
