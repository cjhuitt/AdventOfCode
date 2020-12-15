package grid

type seat struct {
	state rune
}

func readRow(in string) []seat {
	row := []seat{}
	for _, r := range in {
		switch r {
		case 'L', '.':
			row = append(row, seat{r})
		default:
			return []seat{}
		}
	}
	return row
}

func seatSlicesEqual(a, b []seat) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
