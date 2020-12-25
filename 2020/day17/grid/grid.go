package grid

type cell struct {
	active    bool
	neighbors []*cell
}

func readCell(val rune) *cell {
	c := cell{active: val == '#'}
	return &c
}

type list struct {
	contents []*cell
}

func (l *list) Length() int {
	return len(l.contents)
}

func (l *list) contains(row, col, plane int) bool {
	return false
}

type grid struct {
	list []*cell
}

func Parse(in []string) grid {
	g := grid{}
	for _, line := range in {
		for _, val := range line {
			g.list = append(g.list, readCell(val))
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
