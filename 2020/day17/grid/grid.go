package grid

type cell struct {
	active    bool
	neighbors []*cell
}
type grid struct {
	origin cell
}

func Parse(in []string) grid {
	g := grid{}
	return g
}

func (g *grid) NumActive() int {
	return 0
}
