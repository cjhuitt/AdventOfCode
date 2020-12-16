package routing

type ship struct {
	pos    loc
	orient rune
}

func Ship() ship {
	return ship{loc{}, 'E'}
}

func (s ship) Moved(in string) ship {
	return s
}
