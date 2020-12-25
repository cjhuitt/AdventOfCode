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

func Calculate(in string) int {
	var s scanner.Scanner
	s.Init(strings.NewReader(in))
	var operand int
	var op string
	for tok := s.Scan(); tok != scanner.EOF; tok = s.Scan() {
		t := s.TokenText()
		switch t {
		case "+":
			op = "+"
		default:
			if op == "" {
				operand = toInt(t)
			} else {
				var b int
				b = toInt(t)
				operand = operand + b
			}
		}
	}
	fmt.Println(in, "=", operand)
	return operand
}
