package calculator

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"text/scanner"
)

func toInt(in string) int {
	i, err := strconv.Atoi(in)
	if err != nil {
		log.Fatal(err)
	}
	return i
}

func tokenize(in string) []string {
	var s scanner.Scanner
	s.Init(strings.NewReader(in))
	tokens := []string{}
	for tok := s.Scan(); tok != scanner.EOF; tok = s.Scan() {
		tokens = append(tokens, s.TokenText())
	}
	return tokens
}

func calc(tokens []string) int {
	var value int
	var op string
	for _, t := range tokens {
		switch t {
		case "+":
			op = "+"
		default:
			if op == "" {
				value = toInt(t)
			} else {
				var b int
				b = toInt(t)
				value = value + b
			}
		}
	}
	return value
}

func Calculate(in string) int {
	value := calc(tokenize(in))
	fmt.Println(in, "=", value)
	return value
}
