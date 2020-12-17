package routing

import "strconv"

type ship struct {
	pos    loc
	orient rune
}

func Ship() ship {
	return ship{loc{}, 'E'}
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
	case 'N':
		r.pos = r.pos.North(v)
	case 'E':
		r.pos = r.pos.East(v)
	case 'W':
		r.pos = r.pos.West(v)
	case 'S':
		r.pos = r.pos.South(v)
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
	case 'F':
		switch r.orient {
		case 'N':
			r.pos = r.pos.North(v)
		case 'E':
			r.pos = r.pos.East(v)
		case 'W':
			r.pos = r.pos.West(v)
		case 'S':
			r.pos = r.pos.South(v)
		}
	}
	return r
}
