package WireRouting

type path struct {
	path []node
}

func Default() path {
	return path{[]node{}}
}

func Route(r string) path {
	return path{[]node{}}
}

func (p path) Length() int {
	return len(p.path)
}

func (p path) Contains(n node) bool {
	for _, test := range p.path {
		if n.EqualTo(test) {
			return true
		}
	}
	return false
}
