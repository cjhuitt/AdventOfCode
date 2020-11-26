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
	pp := WireRouting.ClosestPhysical(r0.Intersections(r1))
	pr := WireRouting.ClosestRouted(r0.Intersections(r1))

	fmt.Println("Closest physical intersection at", pp, "distance", pp.ManhattanLength())
	fmt.Println("Closest routed intersection at", pr, "distance", pr.RouteLength())
}
