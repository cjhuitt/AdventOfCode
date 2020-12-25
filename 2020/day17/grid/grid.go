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

//==============================================================================
type cell struct {
	active bool
	loc    coord
}

func readCell(val rune, loc coord) *cell {
	c := cell{active: val == '#', loc: loc}
	return &c
}

func createCell(loc coord) *cell {
	c := cell{loc: loc}
	return &c
}

func (c *cell) isAt(loc coord) bool {
	return c.loc == loc
}

//==============================================================================
type list struct {
	contents []*cell
}

func (l *list) Length() int {
	return len(l.contents)
}

func (l *list) contains(loc coord) bool {
	return l.find(loc) != nil
}

func (l *list) add(c *cell) {
	if c != nil {
		l.contents = append(l.contents, c)
	}
}

func (l *list) find(loc coord) *cell {
	for _, c := range l.contents {
		if c.isAt(loc) {
			return c
		}
	}

	return nil
}

func (l *list) findOrAdd(loc coord) *cell {
	c := l.find(loc)
	if c == nil {
		c = createCell(loc)
		l.add(c)
	}
	return c
}

func (l *list) numActive() int {
	total := 0
	for _, c := range l.contents {
		if c.active {
			total++
		}
	}
	return total
}

//==============================================================================
type grid struct {
	universe list
}

func Parse(in []string) grid {
	g := grid{}
	for row, line := range in {
		for col, val := range line {
			g.universe.add(readCell(val, inPlane(row, col)))
		}
	}
	return g
}

func (g *grid) NumActive() int {
	return g.universe.numActive()
}

func (g *grid) Neighbors(loc coord) list {
	l := list{}
	for _, n := range loc.neighbors() {
		l.add(g.universe.findOrAdd(n))
	}
	return l
}
