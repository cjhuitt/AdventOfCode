package grid

type seat struct {
	state     rune
	neighbors []*seat
}

type state struct {
	seats [][]seat
}

func newSeat(state rune) seat {
	return seat{state, []*seat{}}
}

func readRow(in string) []seat {
	row := []seat{}
	for _, r := range in {
		switch r {
		case 'L', '.':
			row = append(row, newSeat(r))
		default:
			return []seat{}
		}
	}
	return row
}

func (s seat) isEqualTo(other seat) bool {
	return s.state == other.state
}

func seatSlicesEqual(a, b []seat) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if !a[i].isEqualTo(b[i]) {
			return false
		}
	}
	return true
}

func readSeating(in []string) state {
	s := state{}
	l := 0
	for _, i := range in {
		r := readRow(i)
		if l == 0 {
			l = len(r)
		} else if len(r) != l {
			return state{}
		}
		s.seats = append(s.seats, r)
	}
	return s
}

func (s state) isEqualTo(other state) bool {
	if len(s.seats) != len(other.seats) {
		return false
	}
	for i := 0; i < len(s.seats); i++ {
		if !seatSlicesEqual(s.seats[i], other.seats[i]) {
			return false
		}
	}
	return true
}

func Parse(in []string) state {
	return readSeating(in)
}
