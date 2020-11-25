package WireRouting

type path struct {
	nodes []node
}

func Default() path {
	return path{[]node{}}
}

func Route(r string) path {
	return path{[]node{}}
}

func (p path) Length() int {
	return len(p.nodes)
}

func (p path) Contains(n node) bool {
	for _, test := range p.nodes {
		if n.EqualTo(test) {
			return true
		}
	}
	return false
}
