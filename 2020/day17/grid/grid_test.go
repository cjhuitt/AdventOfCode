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

func TestNeighborCoords(t *testing.T) {
	o := origin()
	got := o.neighbors()
	if len(got) != 26 {
		t.Errorf("Expected to get 26 neighbors, received %v", len(got))
	} else {
		for row := -1; row <= 1; row++ {
			for col := -1; col <= 1; col++ {
				for plane := -1; plane <= 1; plane++ {
					loc := at(row, col, plane)
					if !loc.isOrigin() && !contains(got, loc) {
						t.Errorf("Expected neighbors to contain %v, it did not", loc)
					}
				}
			}
		}
	}
}

func TestStepping(t *testing.T) {
	g := Parse([]string{".#.", "..#", "###"})
	tests := []struct {
		steps int
		want  int
	}{
		{steps: 0, want: 5},
		{steps: 1, want: 11},
	}
	for i, tc := range tests {
		for j := 0; j < tc.steps; j++ {
			g.Step()
		}
		got := g.NumActive()
		if got != tc.want {
			t.Errorf("Expected stepping to result in %v active cells, received %v (case %d)", tc.want, got, i)
		}
	}
}
