package routing

import "testing"

func TestMove(t *testing.T) {
	tests := []struct {
		move string
		want ship
	}{
		{move: "", want: ship{loc{0, 0}, 'E'}},
	}
	for i, tc := range tests {
		got := Ship().Moved(tc.move)
		if tc.want != got {
			t.Errorf("Ship().Move(%v) want %v, got %v (case %d)", tc.move, tc.want, got, i)
		}
	}
}
