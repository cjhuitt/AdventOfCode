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

func (n loc) East(steps int) loc {
	return loc{n.x + steps, n.y}
}

func (n loc) West(steps int) loc {
	return loc{n.x - steps, n.y}
}

func (n loc) North(steps int) loc {
	return loc{n.x, n.y + steps}
}

func (n loc) South(steps int) loc {
	return loc{n.x, n.y - steps}
}
