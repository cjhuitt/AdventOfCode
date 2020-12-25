package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"grid"
)

func readFile(infile string) []string {
	file, err := os.Open(infile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	lines := []string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return lines
}

func process(infile string, steps int) {
	lines := readFile(infile)
	g := grid.Parse(lines)
	fmt.Println(infile, ":", g.NumActive(), "active at start")

	for i := 0; i < steps; i++ {
		g.Step()
	}
	fmt.Println(infile, ":", g.NumActive(), "active after", steps, "steps")
}

func main() {
	process("test_input.txt", 6)
	fmt.Println()
	process("input.txt", 6)
}
