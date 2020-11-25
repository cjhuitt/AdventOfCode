package WireRouting

import "strings"

type path struct {
	nodes []node
}

func Default() path {
	return path{[]node{}}
}

func Route(r string) path {
	if r == "" {
		return path{[]node{}}
	}
	nodes := []node{}
	steps := strings.Split(r, ",")
	for i, _ := range steps {
		nodes = append(nodes, node{i + 1, 0})
	}
	return path{nodes}
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
