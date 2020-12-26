package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"seats"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	occupied := make(map[int]bool)
	max := -1
	min := 999
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		id := seats.Find(scanner.Text()).Id()
		occupied[id] = true
		if id > max {
			max = id
		}
		if id < min {
			min = id
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	for i := min; i < max; i++ {
		if !occupied[i] {
			fmt.Println("Unoccupied:", i)
			break
		}
	}

	fmt.Println("Max id:", max)
}
