package WireRouting

type section struct{ a, b node }

func DefaultSection() section {
	return section{}
}

func (s section) Length() int {
	return 0
}
