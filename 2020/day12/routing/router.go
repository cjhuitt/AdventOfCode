package routing

import "strconv"

type ship struct {
	pos loc
	way loc
}

func Ship() ship {
	return ship{loc{}, loc{10, 1}}
}

func (s ship) newWayPos(dir byte, steps int) loc {
	switch dir {
	case 'N':
		return s.way.North(steps)
	case 'E':
		return s.way.East(steps)
	case 'W':
		return s.way.West(steps)
	case 'S':
		return s.way.South(steps)
	}
	return s.pos
}

func rotateCw(way loc) loc {
	return loc{way.y, -way.x}
}

func rotateCcw(way loc) loc {
	return loc{-way.y, way.x}
}

func (s ship) Stepped(in string) ship {
	if len(in) < 2 {
		return s
	}

	a := in[0]                     // action
	v, err := strconv.Atoi(in[1:]) // value
	if err != nil {
		return s
	}

	r := s
	switch a {
	case 'N', 'E', 'W', 'S':
		r.way = s.newWayPos(a, v)
	case 'F':
		r.pos = s.pos.AddedTo(s.way.Multiplied(v))
	case 'R':
		for i := 0; i < v/90; i++ {
			r.way = rotateCw(r.way)
		}
	case 'L':
		for i := 0; i < v/90; i++ {
			r.way = rotateCcw(r.way)
		}
	}
	return r
}

func (s ship) Moved(actions []string) ship {
	r := s
	for _, a := range actions {
		r = r.Stepped(a)
	}
	return r
}

func (s ship) Position() loc {
	return s.pos
}
