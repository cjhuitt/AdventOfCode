package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func addValue(buffer []int, value int, max_size int) []int {
	if len(buffer) >= max_size {
		buffer = buffer[1:max_size]
	}

	return append(buffer, value)
}

func valueAllowed(buffer []int, value int, max_size int) bool {
	if len(buffer) < max_size {
		return true
	}
	for i := 0; i < max_size; i++ {
		for j := i + 1; j < max_size; j++ {
			if buffer[i]+buffer[j] == value {
				return true
			}
		}
	}
	return false
}

func findInvalid(infile string, preamble int) int {
	file, err := os.Open(infile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	buffer := []int{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		i, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		if !valueAllowed(buffer, i, preamble) {
			fmt.Println(infile, ":", i, "not allowed")
			return i
		}
		buffer = addValue(buffer, i, preamble)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(infile, ": No invalid values")
	return -1
}

func readTo(infile string, target int) []int {
	file, err := os.Open(infile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	buffer := []int{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		i, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		if i == target {
			return buffer
		}

		buffer = append(buffer, i)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return buffer
}

func findSumRange(buffer []int, target int) (int, int) {
	size := len(buffer)
	for i := 0; i < size; i++ {
		curr := buffer[i]
		for j := i + 1; j < size; j++ {
			curr += buffer[j]
			if curr == target {
				return i, j
			} else if curr > target {
				break
			}
		}
	}

	return -1, -1
}

func findWeakness(infile string, preamble int, target int) {
	file, err := os.Open(infile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	buffer := readTo(infile, target)
	start, end := findSumRange(buffer, target)
	min := buffer[start]
	max := buffer[start]
	for i := start + 1; i <= end; i++ {
		if buffer[i] < min {
			min = buffer[i]
		}
		if buffer[i] > max {
			max = buffer[i]
		}
	}
	fmt.Println(infile, ":", buffer[start], "through", buffer[end], "give weakness:", min+max)
}

func main() {
	invalid := findInvalid("test_input.txt", 5)
	findWeakness("test_input.txt", 5, invalid)

	invalid = findInvalid("input.txt", 25)
	findWeakness("input.txt", 25, invalid)

}
