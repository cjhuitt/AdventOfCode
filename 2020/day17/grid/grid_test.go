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
	got := g.Neighbors(origin())
	if got.Length() != 26 {
		t.Errorf("Expected to get 26 neighbors, received %v", got.Length())
	} else {
		for row := -1; row <= 1; row++ {
			for col := -1; col <= 1; col++ {
				for plane := -1; plane <= 1; plane++ {
					loc := at(row, col, plane)
					if !loc.isOrigin() && !got.contains(loc) {
						t.Errorf("Expected neighbors to contain %v, it did not", loc)
					}
				}
			}
		}
	}
}

func TestCalculateStep(t *testing.T) {
	tests := []struct {
		active    bool
		neighbors []bool
		want      bool
	}{
		// Turn active with three & only 3 active neighbors
		{active: false, neighbors: []bool{true, false, false, false}, want: false},
		{active: false, neighbors: []bool{true, true, false, false}, want: false},
		{active: false, neighbors: []bool{true, true, true, false}, want: true},
		{active: false, neighbors: []bool{true, true, true, true}, want: false},

		// stay active with only 2 or 3 active neighbors
		{active: true, neighbors: []bool{true, false, false, false}, want: false},
		{active: true, neighbors: []bool{true, true, false, false}, want: true},
		{active: true, neighbors: []bool{true, true, true, false}, want: true},
		{active: true, neighbors: []bool{true, true, true, true}, want: false},
	}
	for i, tc := range tests {
		c := cell{active: tc.active}
		n := list{}
		for i, v := range tc.neighbors {
			tmp := n.findOrAdd(at(i, 0, 0))
			tmp.active = v
		}
		c.calculateStep(n)
		got := c.next_state
		if got != tc.want {
			t.Errorf("Expected calculateNext() to have next state active %v, received %v (case %d)", tc.want, got, i)
		}
	}
}
