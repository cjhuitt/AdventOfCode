package main

import (
	"fmt"
)

func play(start []int, turns int) {
	fmt.Println("Starting with", start, "after", turns, "turns the number is")
}

func main() {
	play([]int{1, 3, 2}, 2020)
}
