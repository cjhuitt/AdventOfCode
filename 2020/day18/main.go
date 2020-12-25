package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"calculator"
)

func process(infile string) {
	file, err := os.Open(infile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	sum := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		sum += calculator.Calculate(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(infile, ": lines sum to", sum)
}

func main() {
	process("input.txt")
}
