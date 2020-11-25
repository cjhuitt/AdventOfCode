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
