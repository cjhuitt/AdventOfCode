package grid

//==============================================================================
type cell struct {
	active          bool
	row, col, plane int
}

func readCell(val rune, row, col int) *cell {
	c := cell{active: val == '#', row: row, col: col}
	return &c
}

func (c *cell) isAt(row, col, plane int) bool {
	return c.row == row && c.col == col && c.plane == plane
}

//==============================================================================
type list struct {
	contents []*cell
}

func (l *list) Length() int {
	return len(l.contents)
}

func (l *list) contains(row, col, plane int) bool {
	return false
}

func (l *list) add(c *cell) {
	l.contents = append(l.contents, c)
}

func (l *list) findOrAdd(row, col, plane int) *cell {
	for _, c := range l.contents {
		if c.isAt(row, col, plane) {
			return c
		}
	}

	return nil
}

//==============================================================================
type grid struct {
	universe []*cell
}

func Parse(in []string) grid {
	g := grid{}
	for row, line := range in {
		for col, val := range line {
			g.universe = append(g.universe, readCell(val, row, col))
		}
	}
	return g
}

func (g *grid) NumActive() int {
	total := 0
	for _, c := range g.universe {
		if c.active {
			total++
		}
	}
	return total
}

func (g *grid) Neighbors(row, col, plane int) list {
	l := list{make([]*cell, 26)}
	for r := row - 1; r <= row+1; r++ {
		for c := col - 1; c <= col+1; c++ {
			for p := plane - 1; p <= plane+1; p++ {
				if r != row || c != col || p != plane {
					//l.add(g.list.findOrAdd(r, c, p))
				}
			}
		}
	}
	return l
}
