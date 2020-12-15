package grid

type seat struct {
	state     rune
	neighbors []*seat
}

type deck struct {
	seats         [][]seat
	width, height int
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

func readSeating(in []string) deck {
	s := deck{}
	for _, i := range in {
		r := readRow(i)
		if s.width == 0 {
			s.width = len(r)
		} else if len(r) != s.width {
			return deck{}
		}
		s.seats = append(s.seats, r)
	}
	s.height = len(in)
	return s
}

func (s deck) isEqualTo(other deck) bool {
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

func Parse(in []string) deck {
	init := readSeating(in)
	return init
}
