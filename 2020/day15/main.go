package main

import (
	"fmt"
)

type game struct {
	mem map[int]int
}

func NewGame(start []int) game {
	g := game{}
	g.mem = make(map[int]int)
	for i, n := range start {
		g.mem[n] = i + 1
	}
	return g
}

func (g *game) Turn() int {
	return 0
}

func (g *game) NextTurn() int {
	return 0
}

func play(start []int, turns int) {
	g := NewGame(start)
	last := 0
	for g.Turn() > turns {
		last = g.NextTurn()
	}

	fmt.Println("Starting with", start, "after", turns, "turns the last number is", last)
	fmt.Println()
}

func main() {
	play([]int{0, 3, 6}, 4)
	//play([]int{1, 3, 2}, 2020)
	//play([]int{2, 1, 3}, 2020)
	//play([]int{1, 2, 3}, 2020)
	//play([]int{2, 3, 1}, 2020)
	//play([]int{3, 2, 1}, 2020)
	//play([]int{3, 1, 2}, 2020)

	//fmt.Println()
	//play([]int{9, 19, 1, 6, 0, 5, 4}, 2020)
}
