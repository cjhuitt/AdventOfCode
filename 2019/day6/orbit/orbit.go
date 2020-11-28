package orbit

import (
	"errors"
	"strings"
)

type body struct {
	id       string
	steps    int
	orbits   string
	orbiting *body
}

type bodylist = map[string]*body

func Parse(code string) (string, string, error) {
	parts := strings.Split(code, ")")
	if len(parts) != 2 {
		return "", "", errors.New("fail")
	}

	return parts[0], parts[1], nil
}

func Chart(codes []string) (bodylist, error) {
	chart := bodylist{}
	chart["COM"] = NewBody("", "COM")
	for _, in := range codes {
		c, id, err := Parse(in)
		if err != nil {
			return bodylist{}, err
		}
		chart[id] = NewBody(c, id)
	}
	for _, b := range chart {
		b.orbiting = chart[b.orbits]
	}

	return chart, nil
}

func NewBody(orbits, id string) *body {
	b := new(body)
	b.id = id
	b.orbits = orbits
	return b
}

func (b *body) StepsTo(t *body) int {
	i := 1
	c := b.orbiting
	for c != nil && c != t {
		i++
		c = c.orbiting
	}
	return i
}
