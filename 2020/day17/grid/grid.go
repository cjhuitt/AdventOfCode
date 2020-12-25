package grid

//==============================================================================
type coord struct {
	row, col, plane int
}

func at(row, col, plane int) coord {
	return coord{row, col, plane}
}

func inPlane(row, col int) coord {
	return coord{row, col, 0}
}

func origin() coord {
	return coord{}
}

func (c *coord) isOrigin() bool {
	return c.row == 0 && c.col == 0 && c.plane == 0
}

func (me *coord) neighbors() []coord {
	l := []coord{}
	for r := me.row - 1; r <= me.row+1; r++ {
		for c := me.col - 1; c <= me.col+1; c++ {
			for p := me.plane - 1; p <= me.plane+1; p++ {
				t := at(r, c, p)
				if t != *me {
					l = append(l, t)
				}
			}
		}
	}

	return l
}

func contains(l []coord, t coord) bool {
	for _, c := range l {
		if c == t {
			return true
		}
	}

	return false
}

//==============================================================================
type grid struct {
	active []coord
}

func Parse(in []string) grid {
	g := grid{}
	for row, line := range in {
		for col, val := range line {
			if val == '#' {
				g.active = append(g.active, inPlane(row, col))
			}
		}
	}
	return g
}

func (g *grid) NumActive() int {
	return len(g.active)
}

func (g *grid) Step() {
	counts := map[coord]int{}
	for _, c := range g.active {
		for _, n := range c.neighbors() {
			counts[n] += 1
		}
	}

	new_active := []coord{}
	for coord, num := range counts {
		if contains(g.active, coord) && (num == 2 || num == 3) {
			new_active = append(new_active, coord)
		} else if num == 3 {
			new_active = append(new_active, coord)
		}
	}

	g.active = new_active
}

func (g *grid) StepTo(count int) {
	for i := 0; i < count; i++ {
		g.Step()
	}
}
