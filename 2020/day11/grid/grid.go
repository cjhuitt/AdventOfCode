package grid

import "strings"

type seat struct {
	state      rune
	neighbors  []*seat
	next_state rune
}

type deck struct {
	seats         [][]*seat
	width, height int
}

func newSeat(state rune) *seat {
	s := new(seat)
	s.state = state
	return s
}

func readRow(in string) []*seat {
	row := []*seat{}
	for _, r := range in {
		switch r {
		case 'L', '.', '#':
			row = append(row, newSeat(r))
		default:
			return []*seat{}
		}
	}
	return row
}

func (s *seat) isOccupied() bool {
	return s.state == '#'
}

func (s *seat) calculateNext() {
	occupied := 0
	if s.state == '.' {
		s.next_state = '.'
		return
	}

	for _, n := range s.neighbors {
		if n.isOccupied() {
			occupied++
		}
	}

	if s.state == 'L' && occupied == 0 {
		s.next_state = '#'
	} else if s.state == '#' && occupied >= 4 {
		s.next_state = 'L'
	} else {
		s.next_state = s.state
	}
}

func (s *seat) step() {
	s.state = s.next_state
}

func (s seat) isEqualTo(other seat) bool {
	return s.state == other.state
}

func seatSlicesEqual(a, b []*seat) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if !a[i].isEqualTo(*b[i]) {
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

func (d *deck) isEqualTo(other deck) bool {
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

func (d *deck) northwestOf(i, j int) *seat {
	if i <= 0 || j <= 0 {
		return nil
	}
	return d.seats[i-1][j-1]
}

func (d *deck) northOf(i, j int) *seat {
	if j <= 0 {
		return nil
	}
	return d.seats[i][j-1]
}

func (d *deck) northeastOf(i, j int) *seat {
	if i >= d.width-1 || j <= 0 {
		return nil
	}
	return d.seats[i+1][j-1]
}

func (d *deck) eastOf(i, j int) *seat {
	if i >= d.width-1 {
		return nil
	}
	return d.seats[i+1][j]
}

func (d *deck) southeastOf(i, j int) *seat {
	if i >= d.width-1 || j >= d.height-1 {
		return nil
	}
	return d.seats[i+1][j+1]
}

func (d *deck) southOf(i, j int) *seat {
	if j >= d.height-1 {
		return nil
	}
	return d.seats[i][j+1]
}

func (d *deck) southwestOf(i, j int) *seat {
	if i <= 0 || j >= d.height-1 {
		return nil
	}
	return d.seats[i-1][j+1]
}

func (d *deck) westOf(i, j int) *seat {
	if i <= 0 {
		return nil
	}
	return d.seats[i-1][j]
}

func (d *deck) neighborsOf(i, j int) []*seat {
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
		for j := 0; j < init.height; j++ {
			init.seats[i][j].neighbors = init.neighborsOf(i, j)
		}
	}

	return init
}

func (d *deck) Step() {
	for i := 0; i < d.width; i++ {
		for j := 0; j < d.height; j++ {
			d.seats[i][j].calculateNext()
		}
	}

	for i := 0; i < d.width; i++ {
		for j := 0; j < d.height; j++ {
			d.seats[i][j].step()
		}
	}
}

func (d *deck) Printable() string {
	var l strings.Builder
	for i := 0; i < d.width; i++ {
		for j := 0; j < d.height; j++ {
			l.WriteRune(d.seats[i][j].state)
		}
		l.WriteRune('\n')
	}

	return l.String()
}
