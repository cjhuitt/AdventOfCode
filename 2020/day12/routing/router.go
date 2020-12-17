package routing

import "strconv"

type ship struct {
	pos    loc
	way    loc
	orient byte
}

func Ship() ship {
	return ship{loc{}, loc{10, 0}, 'E'}
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

func rightFrom(orient byte) byte {
	switch orient {
	case 'N':
		return 'E'
	case 'E':
		return 'S'
	case 'S':
		return 'W'
	case 'W':
		return 'N'
	}

	return orient
}

func leftFrom(orient byte) byte {
	switch orient {
	case 'N':
		return 'W'
	case 'E':
		return 'N'
	case 'S':
		return 'E'
	case 'W':
		return 'S'
	}

	return orient
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
		r.pos = s.newPos(a, v)
	case 'F':
		r.pos = s.newPos(r.orient, v)
	case 'R':
		for i := 0; i < v/90; i++ {
			r.orient = rightFrom(r.orient)
		}
	case 'L':
		for i := 0; i < v/90; i++ {
			r.orient = leftFrom(r.orient)
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
