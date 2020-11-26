package WireRouting

type section struct{ a, b node }

func (s section) ManhattanLength() int {
	return abs(s.a.x-s.b.x) + abs(s.a.y-s.b.y)
}

func (s section) Intersect(other section) (bool, node) {
	return true, node{}
}
