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

func findSoonestAfter(earliest int, frequencies []int) (int, int) {
	if len(frequencies) < 1 {
		return -1, -1
	}
	freq := frequencies[0]
	wait := earliest % freq
	for _, f := range frequencies {
		w := earliest % f
		if w < wait {
			freq = f
			wait = w
		}
	}

	return freq, wait
}

func countLines(infile string) {
	lines := read(infile)
	threshold, routes := parse(lines)
	id, wait := findSoonestAfter(threshold, routes)
	fmt.Println(infile, ": After", threshold, "the earliest bus route is", id,
		"after waiting", wait, "(", id*wait, ")")
}

func main() {
	countLines("test_input.txt")
	countLines("input.txt")

}
