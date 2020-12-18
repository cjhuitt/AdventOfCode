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

func parse(input []string) (int, map[int]int) {
	if len(input) != 2 {
		return 0, map[int]int{}
	}

	th, err := strconv.Atoi(input[0])
	if err != nil {
		return 0, map[int]int{}
	}

	ids := map[int]int{}
	for i, s := range strings.Split(input[1], ",") {
		id, err := strconv.Atoi(s)
		if err == nil {
			ids[id] = i
		}
	}
	return th, ids
}

func findSoonestAfter(earliest int, frequencies []int) (int, int) {
	if len(frequencies) < 1 {
		return -1, -1
	}
	freq := frequencies[0]
	m := earliest % freq
	wait := freq - m
	for _, f := range frequencies {
		m = earliest % f
		w := f - m
		if w < wait {
			freq = f
			wait = w
		}
	}

	return freq, wait
}

func largestId(routes map[int]int) int {
	largest := 0
	for id := range routes {
		if id > largest {
			largest = id
		}
	}
	return largest
}

func timestampWorks(ts int64, routes map[int]int) bool {
	for id, offset := range routes {
		t := ts + int64(offset)
		if t%int64(id) != int64(0) {
			return false
		}
	}
	return true
}

func findMagicTimestamp(routes map[int]int) int64 {
	largest := largestId(routes)
	i := int64(1)
	for true {
		val := int64(largest) * i
		ts := val - int64(routes[largest])
		if timestampWorks(ts, routes) {
			return ts
		}
		i++
	}
	return -1
}

func countLines(infile string) {
	lines := read(infile)
	threshold, routes := parse(lines)

	frequencies := make([]int, len(routes))
	i := 0
	for k := range routes {
		frequencies[i] = k
		i++
	}
	id, wait := findSoonestAfter(threshold, frequencies)
	fmt.Println(infile, ": After", threshold, "the earliest bus route is", id,
		"after waiting", wait, "(", id*wait, ")")

	fmt.Println(infile, ": Magic timestamp is", findMagicTimestamp(routes))
}

func main() {
	countLines("test_input.txt")
	countLines("input.txt")

}
