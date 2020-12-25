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

func findParens(tokens []string) (int, int) {
	var start int
	for i, t := range tokens {
		switch t {
		case "(":
			start = i
		case ")":
			return start, i
		}
	}
	return 0, 0
}

func processParens(tokens []string) []string {
	start, end := findParens(tokens)
	if start == end {
		return tokens
	}
	val := calc(tokens[start+1 : end])
	r := tokens[0:start]
	r = append(r, fmt.Sprintf("%d", val))
	r = append(r, tokens[end:len(tokens)-1]...)
	return r
}

func calc(tokens []string) int {
	extracted := processParens(tokens)
	for len(extracted) != len(tokens) {
		tokens = extracted
		extracted = processParens(tokens)
	}
	var stored int
	for i := len(extracted) - 1; i >= 0; i-- {
		t := extracted[i]
		switch t {
		case "+":
			return stored + calc(extracted[0:i])
		case "*":
			return stored * calc(extracted[0:i])
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
