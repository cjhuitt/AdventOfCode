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
		for i := 0; i < v; i++ {
			r.pos = r.pos.North(1)
		}
	}
	return r
}
