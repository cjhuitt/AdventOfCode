package seats

type seat struct{ row, col, id int }

func invalid() seat {
	return seat{-1, -1, -1}
}

func getSeat(r, c int) seat {
	if r < 0 || c < 0 {
		return invalid()
	}
	return seat{r, c, r<<3 + c}
}

func findRow(in string) int {
	if len(in) != 7 {
		return -1
	}

	row := 0
	for _, c := range in {
		row = row << 1
		if c == 'B' {
			row += 1
		}
	}

	return row
}

func findCol(in string) int {
	if len(in) != 3 {
		return -1
	}

	col := 0
	for _, c := range in {
		col = col << 1
		if c == 'R' {
			col += 1
		}
	}

	return col
}

func Find(in string) seat {
	if len(in) != 10 {
		return invalid()
	}
	return getSeat(findRow(in[0:7]), findCol(in[7:10]))
}

func (s seat) Id() int {
	return s.id
}

func (s seat) IsValid() bool {
	return s.id != -1
}

func (s seat) isEqualTo(other seat) bool {
	return s.id == other.id
}
