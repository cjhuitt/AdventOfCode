package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"routing"
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

func countLines(infile string) {
	s := routing.Ship().Moved(read(infile))
	fmt.Println(infile, ":", s.Position(), "ends", s.Position().ManhattanLength(), "away")
}

func main() {
	countLines("test_input.txt")
	countLines("input.txt")

}
