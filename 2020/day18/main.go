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

	var sum1, sum2 int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		sum1 += calculator.Calculate(scanner.Text())
		sum2 += calculator.Calculate2(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(infile, ": lines sum to", sum1, "(left-to-right)")
	fmt.Println(infile, ": lines sum to", sum2, "(addition has priority)")
}

func main() {
	process("input.txt")
}
