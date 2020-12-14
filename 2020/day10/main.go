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

func arrange(in []int) []int {
	out := make([]int, len(in))
	copy(out, in)
	for i := 0; i < len(out); i++ {
		for j := i; j < len(out); j++ {
			if out[i] > out[j] {
				out[i], out[j] = out[j], out[i]
			}
		}
	}

	return out
}

func countSteps(in []int) []int {
	steps := make([]int, 3)
	for i := 1; i < len(in); i++ {
		diff := in[i] - in[i-1]
		if diff <= 3 {
			steps[diff] += 1
		}
	}

	return steps
}

func analyze(infile string) {
	input := readFrom(infile)
	fmt.Println(infile, ":", len(input), "lines read")
}

func main() {
	analyze("test_input.txt")
	analyze("input.txt")
}
