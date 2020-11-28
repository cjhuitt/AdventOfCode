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
	chart["COM"] = newBody("", "COM")
	for _, in := range codes {
		c, id, err := Parse(in)
		if err != nil {
			return bodylist{}, err
		}
		chart[id] = newBody(c, id)
	}
	for _, b := range chart {
		b.orbiting = chart[b.orbits]
	}

	return chart, nil
}

func newBody(orbits, id string) *body {
	b := new(body)
	b.id = id
	b.orbits = orbits
	return b
}

func (b *body) StepsTo(t *body) int {
	if t == nil || b == nil {
		return -1
	}

	if b == t {
		return 0
	}

	if b.orbiting == nil {
		return -1
	}

	next := b.orbiting.StepsTo(t)
	if next >= 0 {
		return next + 1
	}

	return -1
}

func (b *body) StepsToCenter() int {
	if b == nil {
		return -1
	}

	if b.id == "COM" {
		return 0
	}

	if b.orbiting == nil {
		return -1
	}

	if b.steps > 0 {
		return b.steps
	}

	c := b.orbiting.StepsToCenter()
	if c >= 0 {
		b.steps = c + 1
		return b.steps
	}
	return -1
}

func TotalStepsIn(c bodylist) (int, error) {
	sum := 0
	for _, b := range c {
		s := b.StepsToCenter()
		if s < 0 {
			return -1, errors.New("No route from " + b.id + " to center")
		}
		sum += b.StepsToCenter()
	}
	return sum, nil
}

func (b *body) toCenter() ([]*body, error) {
	r := []*body{}
	last := b
	for last != nil && last.id != "COM" {
		r = append([]*body{last}, r...)
		last = last.orbiting
	}

	if last != nil {
		return r, nil
	}

	return []*body{}, errors.New("No route to center")
}

func TransfersBetween(c bodylist, from, to string) (int, error) {
	f := c[from]
	if f == nil {
		return -1, errors.New(from + " (from) not in chart")
	}

	t := c[to]
	if t == nil {
		return -1, errors.New(to + " (to) not in chart")
	}

	ftoc, err := f.toCenter()
	if err != nil {
		return -1, err
	}

	ttoc, err := t.toCenter()
	if err != nil {
		return -1, err
	}

	for ftoc[0] == ttoc[0] {
		ftoc = ftoc[1:]
		ttoc = ttoc[1:]
	}

	return len(ftoc) + len(ttoc), nil
}
