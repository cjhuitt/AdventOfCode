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

func twoLargestIds(routes map[int]int) (int, int) {
	keys := []int{}
	for id, _ := range routes {
		keys = append(keys, id)
	}

	keys = reverseSort(keys)
	return keys[0], keys[1]
}

func findFirstMatchMultiplier(largest, other, large_offset, other_offset int) int {
	for i := 0; i <= other; i++ {
		t := largest*i - large_offset + other_offset
		if t%other == 0 {
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
	mult := offsets[id]
	for true {
		if testMultiple(mult, offsets) {
			break
		}
		if mult < 0 {
			return -1
		}
		mult += id
	}

	return mult
}

func findMagicTimestamp(routes map[int]int) int64 {
	fmt.Println(routes)

	largest, other := twoLargestIds(routes)
	offsets := offsetsFor(largest, routes)

	mult := findMultiple(other, offsets)

	ts := int64(largest) * int64(mult)
	ts -= int64(routes[largest])
	return ts
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
	//countLines("input.txt")
}
