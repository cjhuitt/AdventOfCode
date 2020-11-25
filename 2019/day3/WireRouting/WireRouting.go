package WireRouting

type path struct {
	path []node
}

func Default() path {
	return path{[]node{}}
}

func (p path) Length() int {
	return len(p.path)
}
