package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"vm"
)

func main() {
	file, err := os.Open("test_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	source := []string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		source = append(source, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	trace := map[int]int{}
	program := vm.Parse(source)
	instruction := program.Position()
	for instruction >= 0 {
		if trace[instruction] > 0 {
			break
		}
		trace[instruction] += 1
		program = program.Step()
		instruction = program.Position()
	}

	fmt.Println("Accumulator at first repeat instruction is", program.Accumulator())
}
