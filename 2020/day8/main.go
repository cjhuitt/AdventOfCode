package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"vm"
)

func main() {
	file, err := os.Open("input.txt")
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

	_, program := vm.Parse(source).Execute()

	fmt.Println("Accumulator at first repeat instruction is", program.Accumulator())
}
