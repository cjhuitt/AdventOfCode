package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
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

func parse(input []string) (int, []int) {
	if len(input) != 2 {
		return 0, []int{}
	}

	th, err := strconv.Atoi(input[0])
	if err != nil {
		return 0, []int{}
	}

	ids := []int{}
	for _, s := range strings.Split(input[1], ",") {
		id, err := strconv.Atoi(s)
		if err == nil {
			ids = append(ids, id)
		}
	}
	return th, ids
}

func countLines(infile string) {
	lines := read(infile)
	threshold, routes := parse(lines)
	fmt.Println(infile, ": At time", threshold, "the bus routes are", routes)
}

func main() {
	countLines("test_input.txt")
	countLines("input.txt")

}
