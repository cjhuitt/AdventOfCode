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
	step_count := 0
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
		sections = append(sections, section{next, lastnode, step_count, i})
		step_count += i
	}
	return path{sections}
}

func (p path) Intersections(other path) []intersectPoint {
	found := []intersectPoint{}
	for _, s := range p.sections {
		found = append(found, intersects(other.sections, s)...)
	}
	return found
}

func ClosestPhysical(points []intersectPoint) node {
	if len(points) == 0 {
		return node{}
	}

	closest := points[0]
	for _, p := range points {
		if p.loc.ManhattanLength() < closest.loc.ManhattanLength() {
			closest = p
		}
	}
	return closest.loc
}

func contains(points []intersectPoint, n node) bool {
	for _, test := range points {
		if n.EqualTo(test.loc) {
			return true
		}
	}
	return false
}

func intersects(sections []section, s section) []intersectPoint {
	points := []intersectPoint{}
	for _, test := range sections {
		good, node := test.Intersect(s)
		if good {
			points = append(points, node)
		}
	}
	return points
}
