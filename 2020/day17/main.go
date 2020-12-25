package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func countLines(infile string) {
	file, err := os.Open(infile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	lines := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(infile, ":", lines, "lines found")
}

func main() {
	countLines("test_input.txt")
	countLines("input.txt")
}
