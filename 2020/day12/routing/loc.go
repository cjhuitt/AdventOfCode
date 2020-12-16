package routing

type loc struct{ x, y int }

func Loc(x int, y int) loc {
	return loc{x, y}
}

func abs(val int) int {
	if val < 0 {
		return -val
	}
	return val
}

func (n loc) ManhattanLength() int {
	return abs(n.x) + abs(n.y)
}

func (n loc) Right() loc {
	return loc{n.x + 1, n.y}
}

func (n loc) Left() loc {
	return loc{n.x - 1, n.y}
}

func (n loc) Up() loc {
	return loc{n.x, n.y + 1}
}

func (n loc) Down() loc {
	return loc{n.x, n.y - 1}
}
