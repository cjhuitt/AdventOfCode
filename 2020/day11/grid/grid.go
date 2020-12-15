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
		case 'L', '.', '#':
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

func (d deck) isEqualTo(other deck) bool {
	if len(d.seats) != len(other.seats) {
		return false
	}
	for i := 0; i < len(d.seats); i++ {
		if !seatSlicesEqual(d.seats[i], other.seats[i]) {
			return false
		}
	}
	return true
}

func (d deck) northwestOf(i, j int) *seat {
	if i < 0 || j < 0 {
		return nil
	}
	return &d.seats[i-1][j-1]
}

func (d deck) northOf(i, j int) *seat {
	if j < 0 {
		return nil
	}
	return &d.seats[i][j-1]
}

func (d deck) northeastOf(i, j int) *seat {
	if i >= d.width || j < 0 {
		return nil
	}
	return &d.seats[i+1][j-1]
}

func (d deck) eastOf(i, j int) *seat {
	if i >= d.width {
		return nil
	}
	return &d.seats[i+1][j]
}

func (d deck) southeastOf(i, j int) *seat {
	if i >= d.width || j >= d.height {
		return nil
	}
	return &d.seats[i+1][j+1]
}

func (d deck) southOf(i, j int) *seat {
	if j >= d.height {
		return nil
	}
	return &d.seats[i][j+1]
}

func (d deck) southwestOf(i, j int) *seat {
	if i < 0 || j >= d.height {
		return nil
	}
	return &d.seats[i-1][j+1]
}

func (d deck) westOf(i, j int) *seat {
	if i < 0 {
		return nil
	}
	return &d.seats[i-1][j]
}

func (d deck) neighborsOf(i, j int) []*seat {
	r := []*seat{}

	n := d.northwestOf(i, j)
	if n != nil {
		r = append(r, n)
	}

	n = d.northOf(i, j)
	if n != nil {
		r = append(r, n)
	}

	n = d.northeastOf(i, j)
	if n != nil {
		r = append(r, n)
	}

	n = d.eastOf(i, j)
	if n != nil {
		r = append(r, n)
	}

	n = d.southeastOf(i, j)
	if n != nil {
		r = append(r, n)
	}

	n = d.southOf(i, j)
	if n != nil {
		r = append(r, n)
	}

	n = d.southwestOf(i, j)
	if n != nil {
		r = append(r, n)
	}

	n = d.westOf(i, j)
	if n != nil {
		r = append(r, n)
	}

	return r
}

func Parse(in []string) deck {
	init := readSeating(in)

	for i := 0; i < init.width; i++ {
		for j := 0; j < init.width; j++ {
			init.seats[i][j].neighbors = init.neighborsOf(i, j)
		}
	}

	return init
}

func (d deck) Step() {
}
