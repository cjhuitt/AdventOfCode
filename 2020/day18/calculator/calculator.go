package calculator

import (
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
	var stored int
	for i := len(tokens) - 1; i >= 0; i-- {
		t := tokens[i]
		switch t {
		case "+":
			return stored + calc(tokens[0:i])
		case "*":
			return stored * calc(tokens[0:i])
		default:
			stored = toInt(t)
		}
	}
	return stored
}

func Calculate(in string) int {
	value := calc(tokenize(in))
	return value
}
