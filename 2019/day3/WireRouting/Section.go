package WireRouting

type section struct{ a, b node }

func (s section) ManhattanLength() int {
	return abs(s.a.x-s.b.x) + abs(s.a.y-s.b.y)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// WARNING: Only works for vertical or horizontal line segments
func (s section) Intersect(other section) (bool, node) {
	return false, node{}
}
