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

func timestampWorks(ts int64, routes map[int]int) bool {
	for id, offset := range routes {
		t := ts + int64(offset)
		if t%int64(id) != int64(0) {
			return false
		}
	}
	return true
}

func reverseSort(keys []int) []int {
	for i := 0; i < len(keys); i++ {
		for j := i + 1; j < len(keys); j++ {
			if keys[i] < keys[j] {
				keys[i], keys[j] = keys[j], keys[i]
			}
		}
	}
	return keys
}

func largestId(routes map[int]int) int {
	max := -1
	for id, _ := range routes {
		if id > max {
			max = id
		}
	}

	return max
}

func findFirstMatchMultiplier(a, b, offa, offb int) int {
	for i := 0; i <= b; i++ {
		t := a*i + offa - offb
		if t%b == 0 {
			return i
		}
	}

	return -1
}

func testMultiple(m int, offsets map[int]int) bool {
	works := 0
	for id, offset := range offsets {
		t := m - offset
		if t%id == 0 {
			works++
		}
	}

	return works == len(offsets)
}

func offsetsFor(ref int, offsets map[int]int) map[int]int {
	r := map[int]int{}
	for id, time_diff := range offsets {
		if id != ref {
			r[id] = findFirstMatchMultiplier(ref, id, offsets[ref], time_diff)
		}
	}
	return r
}

func findMultiple(id int, offsets map[int]int) int {
	if len(offsets) == 1 {
		for _, offset := range offsets {
			return offset
		}
	}

	a := largestId(offsets)
	o := offsetsFor(a, offsets)
	m := findMultiple(a, o)
	return a*m + offsets[a]
}

func invertOffsets(orig map[int]int) map[int]int {
	inverted := map[int]int{}
	for id, offset := range orig {
		inverted[id] = -offset
	}
	return inverted
}

func findMagicTimestamp(routes map[int]int) int {
	a := largestId(routes)
	o := offsetsFor(a, invertOffsets(routes))
	m := findMultiple(a, o)
	return a*m - routes[a]
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
	fmt.Println()
	countLines("input.txt")
}
