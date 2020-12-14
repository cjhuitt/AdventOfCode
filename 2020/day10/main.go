package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func readFrom(infile string) []int {
	file, err := os.Open(infile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	in := []int{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		i, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		in = append(in, i)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return in
}

func main() {
	input := readFrom("test_input.txt")
	fmt.Println("test_input.txt:", len(input), "lines read")
	input = readFrom("input.txt")
	fmt.Println("input.txt:", len(input), "lines read")
}
