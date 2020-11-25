package WireRouting

import (
	"strconv"
	"strings"
)

type path struct {
	nodes []node
}

func Default() path {
	return path{[]node{}}
}

func Route(r string) path {
	if r == "" {
		return path{[]node{{0, 0}}}
	}
	lastnode := Node(0, 0)
	nodes := []node{lastnode}
	steps := strings.Split(r, ",")
	for _, step := range steps {
		dir := step[0]
		count, err := strconv.Atoi(step[1:])
		if err != nil {
			return Default()
		}
		i := 0
		switch dir {
		case 'R':
			for i < count {
				n := lastnode.Right()
				nodes = append(nodes, n)
				lastnode = n
				i++
			}
		case 'L':
			for i < count {
				n := lastnode.Left()
				nodes = append(nodes, n)
				lastnode = n
				i++
			}
		case 'U':
			for i < count {
				n := lastnode.Up()
				nodes = append(nodes, n)
				lastnode = n
				i++
			}
		case 'D':
			for i < count {
				n := lastnode.Down()
				nodes = append(nodes, n)
				lastnode = n
				i++
			}
		}
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
