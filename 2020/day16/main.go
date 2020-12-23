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

func readMyTicket(scanner *bufio.Scanner) tickets.Ticket {
	t := tickets.Ticket{}
	for scanner.Scan() {
		line := scanner.Text()
		if line == "your ticket:" {
			continue
		}
		if line == "" {
			break
		}
		t = tickets.ParseTicket(line)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return t
}

func readOtherTickets(scanner *bufio.Scanner) []tickets.Ticket {
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

	return others
}

func countLines(infile string) {
	file, err := os.Open(infile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	specs := readSpecs(scanner)
	_ = readMyTicket(scanner)
	others := readOtherTickets(scanner)

	fmt.Println(infile, ":", len(specs), "specifications found")
	fmt.Println(infile, ":", len(others), "other tickets found")

	rate := 0
	for _, t := range others {
		g, e := t.Validate(specs)
		if !g {
			rate += e
		}
	}
	fmt.Println(infile, ": error rate is", rate)
}

func main() {
	countLines("test_input.txt")
	fmt.Println()
	countLines("input.txt")
}
