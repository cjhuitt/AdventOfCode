package grid

//==============================================================================
type coord struct {
	x, y, z int
}

func at(x, y, z int) coord {
	return coord{x, y, z}
}

func inPlane(x, y int) coord {
	return coord{x, y, 0}
}

func origin() coord {
	return coord{}
}

func (c *coord) isOrigin() bool {
	return c.x == 0 && c.y == 0 && c.z == 0
}

func (me *coord) neighbors() []coord {
	l := []coord{}
	for x := me.x - 1; x <= me.x+1; x++ {
		for y := me.y - 1; y <= me.y+1; y++ {
			for z := me.z - 1; z <= me.z+1; z++ {
				t := at(x, y, z)
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
	for x, line := range in {
		for y, val := range line {
			if val == '#' {
				g.active = append(g.active, inPlane(x, y))
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
