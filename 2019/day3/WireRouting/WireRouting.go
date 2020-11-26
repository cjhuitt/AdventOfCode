package WireRouting

import (
	"strconv"
	"strings"
)

type path struct {
	nodes    []node
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
	nodes := []node{}
	sections := []section{}
	steps := strings.Split(r, ",")
	for _, step := range steps {
		dir := step[0]
		count, err := strconv.Atoi(step[1:])
		if err != nil {
			return empty()
		}
		a := lastnode
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
		sections = append(sections, section{a, lastnode})
	}
	return path{nodes, sections}
}

func (p path) Length() int {
	return len(p.nodes)
}

func (p path) Contains(n node) bool {
	return contains(p.nodes, n)
}

func (p path) Intersections(other path) []node {
	intersects := []node{}
	for _, n := range p.nodes {
		if contains(other.nodes, n) {
			intersects = append(intersects, n)
		}
	}
	return intersects
}

func Closest(nodes []node) node {
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
