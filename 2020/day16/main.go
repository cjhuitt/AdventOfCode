package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"tickets"
)

func countLines(infile string) {
	file, err := os.Open(infile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	specs := []tickets.FieldSpec{}
	lines := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		specs = append(specs, tickets.ParseFieldSpec(line))
		lines++
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	for scanner.Scan() {
		line := scanner.Text()
		if line == "your ticket:" {
			continue
		}
		if line == "" {
			break
		}
		lines++
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	others := []tickets.Ticket{}
	for scanner.Scan() {
		line := scanner.Text()
		if line == "nearby tickets:" {
			continue
		}
		if line == "" {
			break
		}
		others = append(others, tickets.ParseTicket(line))
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
