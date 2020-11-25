package WireRouting

type node struct{ x, y int }

func DefaultNode() node {
	return node{0, 0}
}

func Node(x int, y int) node {
	return node{x, y}
}

func (n node) ManhattanLength() int {
	return n.x + n.y
}
