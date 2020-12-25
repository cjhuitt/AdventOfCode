package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"text/scanner"
)

func calculate(in string) int {
	var s scanner.Scanner
	s.Init(strings.NewReader(in))
	var operand int
	var op string
	var err error
	for tok := s.Scan(); tok != scanner.EOF; tok = s.Scan() {
		t := s.TokenText()
		switch t {
		case "+":
			op = "+"
		default:
			if op == "" {
				operand, err = strconv.Atoi(t)
				if err != nil {
					log.Fatal(err)
				}
			} else {
				var b int
				b, err = strconv.Atoi(t)
				if err != nil {
					log.Fatal(err)
				}
				operand = operand + b
			}
		}
	}
	fmt.Println(in, "=", operand)
	return operand
}

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
	calculate("1 + 2")
	//countLines("input.txt")
}
