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

	p := decoder.Program()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		p.Execute(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(infile, ": Memory sum is", p.SumMemory())
}

func main() {
	executeLines("test_input.txt")
	executeLines("input.txt")

}
