package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func read(infile string) []string {
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

func findStasisOccupation(infile string) {
	lines := read(infile)

	fmt.Println(infile, ":", len(lines), "lines found")
}

func main() {
	findStasisOccupation("test_input.txt")
	findStasisOccupation("input.txt")

}
