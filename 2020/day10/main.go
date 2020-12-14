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

func addDeviceAndPort(in []int) []int {
	return append([]int{0}, append(in, in[len(in)-1]+3)...)
}

func countSteps(in []int) []int {
	steps := make([]int, 4)
	for i := 1; i < len(in); i++ {
		diff := in[i] - in[i-1]
		if diff <= 3 {
			steps[diff] += 1
		}
	}

	return steps
}

func analyze(infile string) {
	data := readFrom(infile)
	data = arrange(data)
	data = addDeviceAndPort(data)
	steps := countSteps(data)
	fmt.Println(infile, ":", len(data), "lines read")
	fmt.Println(infile, ":", steps[1], "1-jolt steps,", steps[3], "3-jolt steps")
	fmt.Println(infile, "result:", steps[1]*steps[3])
}

func main() {
	analyze("test_input.txt")
	fmt.Println()
	analyze("input.txt")
}
