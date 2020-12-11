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

func run(infile string, preamble int) {
	file, err := os.Open(infile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	buffer := []int{}
	lines := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		i, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		lines++
		buffer = addValue(buffer, i, preamble)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(infile, ":", lines, "lines read")
}

func main() {
	run("test_input.txt", 5)
	run("input.txt", 25)
}
