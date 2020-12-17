package routing

import "fmt"

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

func (n loc) String() string {
	lat := "E"
	if n.x < 0 {
		lat = "W"
	}
	lon := "N"
	if n.y < 0 {
		lon = "S"
	}
	return fmt.Sprintf("(%s%d, %s%d)", lat, abs(n.x), lon, abs(n.y))
}

func (n loc) Multiplied(mult int) loc {
	return n
}
