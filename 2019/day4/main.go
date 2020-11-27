package main

import (
	"errors"
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

func (g generator) Position() int {
	return g.current
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

func (set generators) Value() int {
	v := 0
	for _, g := range set.places {
		v += g.Value()
	}
	return v
}

func (set generators) Next() (generators, error) {
	r := generators{}
	pivot := len(set.places) - 1
	for ; pivot > 0; pivot-- {
		if !set.places[pivot].AtEnd() {
			break
		}
	}
	for true {
		if pivot < 0 {
			return r, errors.New("exceeded generator capacity")
		}
		r.places[pivot] = set.places[pivot].Step()
		if !r.places[pivot].AtEnd() {
			break
		}
		pivot--
	}
	for i := pivot + 1; i < len(r.places); i++ {
		r.places[i] = r.places[i].ResetTo(r.places[i-1].Position())
	}
	for i := pivot - 1; i >= 0; i-- {
		r.places[i] = set.places[i]
	}
	return r, nil
}

func main() {
	g := InitGenerators(1, 2, 8, 3, 9, 2)
	var err error
	for true {
		fmt.Println(g.Value())
		g, err = g.Next()
		if err != nil {
			fmt.Println("Generated too far")
			break
		}
	}
}
