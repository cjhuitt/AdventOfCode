package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"tickets"
)

func readSpecs(scanner *bufio.Scanner) []tickets.FieldSpec {
	specs := []tickets.FieldSpec{}
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		specs = append(specs, tickets.ParseFieldSpec(line))
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return specs
}

func countLines(infile string) {
	file, err := os.Open(infile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	specs := readSpecs(scanner)

	for scanner.Scan() {
		line := scanner.Text()
		if line == "your ticket:" {
			continue
		}
		if line == "" {
			break
		}
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
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(infile, ":", len(specs), "specifications found")
	fmt.Println(infile, ":", len(others), "other tickets found")
}

func main() {
	countLines("test_input.txt")
	fmt.Println()
	countLines("input.txt")
}
