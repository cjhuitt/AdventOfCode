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
	if max(s.a.x, s.b.x) < min(other.a.x, other.b.x) ||
		max(other.a.x, other.b.x) < min(s.a.x, s.b.x) ||
		max(s.a.y, s.b.y) < min(other.a.y, other.b.y) ||
		max(other.a.y, other.b.y) < min(s.a.y, s.b.y) {
		return false, node{}
	}
	if (s.a.EqualTo(other.a) && s.b.EqualTo(other.b)) ||
		(s.a.EqualTo(other.b) && s.b.EqualTo(other.a)) {
		return true, s.a
	}
	if s.a.EqualTo(other.a) || s.a.EqualTo(other.b) {
		return true, s.a
	}
	if s.b.EqualTo(other.a) || s.b.EqualTo(other.b) {
		return true, s.b
	}

	return true, node{}
}
