package bags

type bag struct{ row, col, id int }

func Parse(in string) bag {
	return bag{}
}

func (b bag) isEqualTo(other bag) bool {
	return true
}
