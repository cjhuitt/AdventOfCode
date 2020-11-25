package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"WireRouting"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	paths := make([]string, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		paths = append(paths, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	if len(paths) < 2 {
		log.Fatal("Not enough paths specified")
	}

	r0 := WireRouting.Route(paths[0])
	r1 := WireRouting.Route(paths[1])
	p := WireRouting.Closest(r0.Intersections(r1))

	fmt.Println("Closest intersection at", p, "distance", p.ManhattanLength())
}
