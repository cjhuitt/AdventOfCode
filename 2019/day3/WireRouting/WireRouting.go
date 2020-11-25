package WireRouting

type path struct {
	length int
}

func Default() path {
	return path{0}
}

func (p path) Length() int {
	return p.length
}
