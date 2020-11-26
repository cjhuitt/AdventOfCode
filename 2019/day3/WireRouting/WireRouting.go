package WireRouting

import (
	"strconv"
	"strings"
)

type path struct {
	sections []section
}

func empty() path {
	return path{}
}

func Route(r string) path {
	if r == "" {
		return empty()
	}
	lastnode := Node(0, 0)
	sections := []section{}
	steps := strings.Split(r, ",")
	for _, step := range steps {
		dir := step[0]
		count, err := strconv.Atoi(step[1:])
		if err != nil {
			return empty()
		}
		var next node
		i := 0
		switch dir {
		case 'R':
			next = lastnode.Right()
			for i < count {
				lastnode = lastnode.Right()
				i++
			}
		case 'L':
			next = lastnode.Left()
			for i < count {
				lastnode = lastnode.Left()
				i++
			}
		case 'U':
			next = lastnode.Up()
			for i < count {
				lastnode = lastnode.Up()
				i++
			}
		case 'D':
			next = lastnode.Down()
			for i < count {
				lastnode = lastnode.Down()
				i++
			}
		}
		sections = append(sections, section{next, lastnode})
	}
	return path{sections}
}

func (p path) Intersections(other path) []node {
	found := []node{}
	for _, s := range p.sections {
		found = append(found, intersects(other.sections, s)...)
	}
	return found
}

func ClosestPhysical(nodes []node) node {
	if len(nodes) == 0 {
		return node{}
	}

	closest := nodes[0]
	for _, n := range nodes {
		if n.ManhattanLength() < closest.ManhattanLength() {
			closest = n
		}
	}
	return closest
}

func contains(nodes []node, n node) bool {
	for _, test := range nodes {
		if n.EqualTo(test) {
			return true
		}
	}
	return false
}

func intersects(sections []section, s section) []node {
	nodes := []node{}
	for _, test := range sections {
		good, node := test.Intersect(s)
		if good {
			nodes = append(nodes, node)
		}
	}
	return nodes
}
