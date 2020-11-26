package WireRouting

type node struct{ x, y int }

func Node(x int, y int) node {
	return node{x, y}
}

func abs(val int) int {
	if val < 0 {
		return -val
	}
	return val
}

func (n node) ManhattanLength() int {
	return abs(n.x) + abs(n.y)
}

func (n node) EqualTo(other node) bool {
	return n.x == other.x && n.y == other.y
}

func (n node) Right() node {
	return node{n.x + 1, n.y}
}

func (n node) Left() node {
	return node{n.x - 1, n.y}
}

func (n node) Up() node {
	return node{n.x, n.y + 1}
}

func (n node) Down() node {
	return node{n.x, n.y - 1}
}
