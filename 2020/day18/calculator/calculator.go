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

func Calculate(in string) int {
	tokens := tokenize(in)
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
	fmt.Println(in, "=", value)
	return value
}
