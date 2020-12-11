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

func findSumSequence(infile string, preamble int, target int) {
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
			return
		}
		buffer = addValue(buffer, i, preamble)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(infile, ": No invalid values")
}

func main() {
	invalid := findInvalid("test_input.txt", 5)
	findSumSequence("test_input.txt", 5, invalid)

	invalid = findInvalid("input.txt", 25)
	findSumSequence("input.txt", 25, invalid)

}
