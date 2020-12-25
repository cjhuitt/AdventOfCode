package grid

import "testing"

func TestInitGridForActiveCells(t *testing.T) {
	tests := []struct {
		input []string
		want  int
	}{
		// empty
		{input: []string{}, want: 0},

		{input: []string{"."}, want: 0},
		{input: []string{"#"}, want: 1},
		{input: []string{".#"}, want: 1},
		{input: []string{".#", "#."}, want: 2},
	}
	for i, tc := range tests {
		g := Parse(tc.input)
		got := g.NumActive()
		if got != tc.want {
			t.Errorf("Expected Parse(%v) to result in %v active cells, received %v (case %d)", tc.input, tc.want, got, i)
		}
	}
}

func TestGetNeighbors(t *testing.T) {
	g := Parse([]string{"."})
	got := g.Neighbors(coord{})
	if got.Length() != 26 {
		t.Errorf("Expected to get 26 neighbors, received %v", got.Length())
	} else {
		for row := -1; row <= 1; row++ {
			for col := -1; col <= 1; col++ {
				for plane := -1; plane <= 1; plane++ {
					if (row != 0 || col != 0 || plane != 0) && !got.contains(at(row, col, plane)) {
						t.Errorf("Expected neighbors to contain (%d, %d, %d), it did not", row, col, plane)
					}
				}
			}
		}
	}
}
