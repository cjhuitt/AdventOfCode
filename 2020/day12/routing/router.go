package routing

import "strconv"

type ship struct {
	pos    loc
	orient byte
}

func Ship() ship {
	return ship{loc{}, 'E'}
}

func (s ship) newPos(dir byte, steps int) loc {
	switch dir {
	case 'N':
		return s.pos.North(steps)
	case 'E':
		return s.pos.East(steps)
	case 'W':
		return s.pos.West(steps)
	case 'S':
		return s.pos.South(steps)
	}
	return s.pos
}

func (s ship) Moved(in string) ship {
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
		r.pos = s.newPos(a, v)
	case 'F':
		r.pos = s.newPos(r.orient, v)
	case 'R':
		for i := 0; i < v/90; i++ {
			switch r.orient {
			case 'N':
				r.orient = 'E'
			case 'E':
				r.orient = 'S'
			case 'S':
				r.orient = 'W'
			case 'W':
				r.orient = 'N'
			}
		}
	case 'L':
		for i := 0; i < v/90; i++ {
			switch r.orient {
			case 'N':
				r.orient = 'W'
			case 'E':
				r.orient = 'N'
			case 'S':
				r.orient = 'E'
			case 'W':
				r.orient = 'S'
			}
		}
	}
	return r
}
