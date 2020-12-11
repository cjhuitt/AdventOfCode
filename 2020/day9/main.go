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

func run(infile string, preamble int) {
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
	run("test_input.txt", 5)
	run("input.txt", 25)
}
