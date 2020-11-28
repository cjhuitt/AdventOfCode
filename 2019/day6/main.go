package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"orbit"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	specs := []string{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		specs = append(specs, scanner.Text())
	}

	chart, err := orbit.Chart(specs)
	if err != nil {
		log.Fatal(err)
	}
	steps, err := orbit.TotalStepsIn(chart)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Found", steps, "steps")

	steps, err = orbit.TransfersBetween(chart, "YOU", "SAN")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Found", steps-2, "transfers")
}
