package WireRouting

type section struct{ a, b node }
type intersectPoint struct {
	loc   node
	steps int
}

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

func det(a, b node) int {
	return a.x*b.y - a.y*b.x
}

// WARNING: Only works for vertical or horizontal line segments
func (s section) Intersect(other section) (bool, intersectPoint) {
	if max(s.a.x, s.b.x) < min(other.a.x, other.b.x) ||
		max(other.a.x, other.b.x) < min(s.a.x, s.b.x) ||
		max(s.a.y, s.b.y) < min(other.a.y, other.b.y) ||
		max(other.a.y, other.b.y) < min(s.a.y, s.b.y) {
		return false, intersectPoint{}
	}
	if (s.a.EqualTo(other.a) && s.b.EqualTo(other.b)) ||
		(s.a.EqualTo(other.b) && s.b.EqualTo(other.a)) {
		return true, intersectPoint{s.a, 0}
	}
	if s.a.EqualTo(other.a) || s.a.EqualTo(other.b) {
		return true, intersectPoint{s.a, 0}
	}
	if s.b.EqualTo(other.a) || s.b.EqualTo(other.b) {
		return true, intersectPoint{s.b, 0}
	}
	// TODO: overlapping case
	// if s.a.x == other.a.x && s.a.x == other.b.x {
	// if min(other.a.y, other.b.y) < min(s.a.y, s.b.y) &&
	// max(other.a.y, other.b.y) > min(s.a.y, s.b.y) {
	// // Overlapping at least some
	// return true, Node(s.a.x, min(s.a.y, s.b.y))
	// }
	// }
	xdiff := node{s.a.x - s.b.x, other.a.x - other.b.x}
	ydiff := node{s.a.y - s.b.y, other.a.y - other.b.y}
	div := det(xdiff, ydiff)
	if div == 0 {
		return false, intersectPoint{}
	}

	d := node{det(s.a, s.b), det(other.a, other.b)}
	x := det(d, xdiff) / div
	y := det(d, ydiff) / div
	return true, intersectPoint{Node(x, y), 0}
}
