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

func main() {
	tens := generator{1, 2, 10}
	fmt.Println(tens, tens.Value())
	tens = tens.Step()
	fmt.Println(tens, tens.Value())
	tens = tens.ResetTo(3)
	fmt.Println(tens, tens.Value())
}
