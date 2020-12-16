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
	changed       bool
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

func (s *seat) willChange() bool {
	return s.state != s.next_state
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
	for row := 0; row < d.height; row++ {
		if !seatSlicesEqual(d.seats[row], other.seats[row]) {
			return false
		}
	}
	return true
}

func (d *deck) northwestOf(row, col int) *seat {
	if row <= 0 || col <= 0 {
		return nil
	}
	return d.seats[row-1][col-1]
}

func (d *deck) northOf(row, col int) *seat {
	if row <= 0 {
		return nil
	}
	return d.seats[row-1][col]
}

func (d *deck) northeastOf(row, col int) *seat {
	if row <= 0 || col <= d.width-1 {
		return nil
	}
	return d.seats[row-1][col+1]
}

func (d *deck) eastOf(row, col int) *seat {
	if col >= d.width-1 {
		return nil
	}
	return d.seats[row][col+1]
}

func (d *deck) southeastOf(row, col int) *seat {
	if row >= d.height-1 || col >= d.width-1 {
		return nil
	}
	return d.seats[row+1][col+1]
}

func (d *deck) southOf(row, col int) *seat {
	if row >= d.height-1 {
		return nil
	}
	return d.seats[row+1][col]
}

func (d *deck) southwestOf(row, col int) *seat {
	if row >= d.height-1 || col <= 0 {
		return nil
	}
	return d.seats[row+1][col-1]
}

func (d *deck) westOf(row, col int) *seat {
	if col <= 0 {
		return nil
	}
	return d.seats[row][col-1]
}

func (d *deck) neighborsOf(row, col int) []*seat {
	r := []*seat{}

	n := d.northwestOf(row, col)
	if n != nil {
		r = append(r, n)
	}

	n = d.northOf(row, col)
	if n != nil {
		r = append(r, n)
	}

	n = d.northeastOf(row, col)
	if n != nil {
		r = append(r, n)
	}

	n = d.eastOf(row, col)
	if n != nil {
		r = append(r, n)
	}

	n = d.southeastOf(row, col)
	if n != nil {
		r = append(r, n)
	}

	n = d.southOf(row, col)
	if n != nil {
		r = append(r, n)
	}

	n = d.southwestOf(row, col)
	if n != nil {
		r = append(r, n)
	}

	n = d.westOf(row, col)
	if n != nil {
		r = append(r, n)
	}

	return r
}

func Parse(in []string) deck {
	init := readSeating(in)
	init.changed = true

	for row := 0; row < init.height; row++ {
		for col := 0; col < init.width; col++ {
			init.seats[row][col].neighbors = init.neighborsOf(row, col)
		}
	}

	return init
}

func (d *deck) Step() {
	for row := 0; row < d.height; row++ {
		for col := 0; col < d.width; col++ {
			d.seats[row][col].calculateNext()
		}
	}

	d.changed = false
	for row := 0; row < d.height && !d.changed; row++ {
		for col := 0; col < d.width && !d.changed; col++ {
			d.changed = d.seats[row][col].willChange()
		}
	}

	for row := 0; row < d.height; row++ {
		for col := 0; col < d.width; col++ {
			d.seats[row][col].step()
		}
	}
}

func (d *deck) Printable() string {
	var l strings.Builder
	for row := 0; row < d.height; row++ {
		for col := 0; col < d.width; col++ {
			l.WriteRune(d.seats[row][col].state)
		}
		l.WriteRune('\n')
	}

	return l.String()
}

func (d *deck) Changed() bool {
	return d.changed
}

func (d *deck) Stabilize(max_iters int) int {
	for i := 0; i < max_iters; i++ {
		d.Step()
		if !d.Changed() {
			return i
		}
	}

	return max_iters
}
