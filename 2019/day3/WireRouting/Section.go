package WireRouting

type section struct{ a, b node }

func DefaultSection() section {
	return section{}
}

func (s section) ManhattanLength() int {
	return abs(s.a.x-s.b.x) + abs(s.a.y-s.b.y)
}
