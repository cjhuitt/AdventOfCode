package main

import (
	"fmt"
)

func play(start []int, turns int) {
	fmt.Println("Starting with", start, "after", turns, "turns the number is")
	fmt.Println()
}

func main() {
	play([]int{0, 3, 6}, 2020)
	//play([]int{1, 3, 2}, 2020)
	//play([]int{2, 1, 3}, 2020)
	//play([]int{1, 2, 3}, 2020)
	//play([]int{2, 3, 1}, 2020)
	//play([]int{3, 2, 1}, 2020)
	//play([]int{3, 1, 2}, 2020)

	//fmt.Println()
	//play([]int{9, 19, 1, 6, 0, 5, 4}, 2020)
}
