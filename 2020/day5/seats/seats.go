package seats

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
