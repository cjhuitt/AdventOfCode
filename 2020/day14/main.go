package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"decoder"
)

func executeLines(infile string) {
	file, err := os.Open(infile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	p1 := decoder.Program()
	p2 := decoder.Program2()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		p1.Execute(scanner.Text())
		p2.Execute(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(infile, ": Variation 1 memory sum is", p1.SumMemory())
	fmt.Println(infile, ": Variation 2 memory sum is", p2.SumMemory())
}

func main() {
	executeLines("test_input.txt")
	fmt.Println()
	executeLines("input.txt")
}
