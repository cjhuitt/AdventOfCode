package seats

type seat struct{ row, col, id int }

func invalid() seat {
	return seat{-1, -1, -1}
}

func getSeat(r, c int) seat {
	if r < 0 || c < 0 {
		return invalid()
	}
	return seat{r, c, r*8 + c}
}

func findRow(in string) int {
	if len(in) != 7 {
		return -1
	}

	row := 0
	half := 64
	for _, c := range in {
		if c == 'B' {
			row += half
		}
		half /= 2
	}
	return row
}

func findCol(in string) int {
	if len(in) != 3 {
		return -1
	}

	col := 0
	half := 4
	for _, c := range in {
		if c == 'R' {
			col += half
		}
		half /= 2
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
	return s.row == other.row && s.col == other.col
}
