package main

import (
	"fmt"
)

type generator struct {
	min, current, mult int
}

func (g generator) Step() generator {
	c := g.current
	c++
	g.current = c
	return generator{g.min, c, g.mult}
}

func (g generator) AtEnd() bool {
	return g.current == 10
}

func (g generator) Value() int {
	return g.current * g.mult
}

func (g generator) ResetTo(min int) generator {
	return generator{min, min, g.mult}
}

type generators struct {
	places [6]generator
}

func InitGenerators(hun_k, ten_k, k, hun, ten, one int) generators {
	a := generator{0, 0, 100000}.ResetTo(hun_k)
	b := generator{0, 0, 10000}.ResetTo(ten_k)
	c := generator{0, 0, 1000}.ResetTo(k)
	d := generator{0, 0, 100}.ResetTo(hun)
	e := generator{0, 0, 10}.ResetTo(ten)
	f := generator{0, 0, 1}.ResetTo(one)

	return generators{[6]generator{a, b, c, d, e, f}}
}

func (set generators) GeneratedValue() int {
	v := 0
	for _, g := range set.places {
		v += g.Value()
	}
	return v
}

func main() {
	g := InitGenerators(1, 2, 8, 3, 9, 2)
	fmt.Println(g.GeneratedValue())
}
