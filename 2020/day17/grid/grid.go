package grid

type cell struct {
	active    bool
	neighbors []*cell
}

func readCell(val byte) cell {
	return cell{active: val == '#'}
}

type grid struct {
	origin cell
}

func Parse(in []string) grid {
	g := grid{}
	if len(in) > 0 && len(in[0]) > 0 {
		g.origin = readCell(in[0][0])
	}
	return g
}

func (g *grid) NumActive() int {
	if g.origin.active {
		return 1
	}
	return 0
}
