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

//==============================================================================
type grid struct {
	list []*cell
}

func Parse(in []string) grid {
	g := grid{}
	for row, line := range in {
		for col, val := range line {
			g.list = append(g.list, readCell(val, row, col))
		}
	}
	return g
}

func (g *grid) NumActive() int {
	total := 0
	for _, c := range g.list {
		if c.active {
			total++
		}
	}
	return total
}

func (g *grid) Neighbors(row, col, plane int) list {
	r := list{make([]*cell, 26)}
	return r
}
