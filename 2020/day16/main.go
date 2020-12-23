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

func fieldEntries(good []tickets.Ticket, field int) []int {
	r := []int{}
	for _, t := range good {
		r = append(r, t.Field(field))
	}
	return r
}

func findFieldOrder(specs []tickets.FieldSpec, good []tickets.Ticket) []tickets.FieldSpec {
	r := []tickets.FieldSpec{}
	for i := 0; i < good[0].NumFields(); i++ {
		entries := fieldEntries(good, i)
		for _, s := range specs {
			if s.PassesAll(entries) {
				r = append(r, s)
				break
			}
		}
	}
	return r
}

func processFile(infile string) {
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
	good := []tickets.Ticket{}
	for _, t := range others {
		g, e := t.Validate(specs)
		if g {
			good = append(good, t)
		} else {
			rate += e
		}
	}
	fmt.Println(infile, ": error rate is", rate)
	fmt.Println(infile, ":", len(good), "good other tickets found")

	ordered := findFieldOrder(specs, good)
	fmt.Printf("%s : ", infile)
	for _, s := range ordered {
		fmt.Printf("%s, ", s.Name())
	}
	fmt.Printf("\n")
}

func main() {
	processFile("test_input.txt")
	fmt.Println()
	processFile("input.txt")
}
