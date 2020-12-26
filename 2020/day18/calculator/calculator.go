package calculator

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"text/scanner"
)

//==============================================================================
func toInt(in string) int {
	i, err := strconv.Atoi(in)
	if err != nil {
		log.Fatal(err)
	}
	return i
}

//==============================================================================
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
	temp := tokens[start+1 : end]
	val := calc(temp)
	r := make([]string, len(tokens)-(end-start))
	if start > 0 {
		copy(r, tokens[0:start])
	}
	r[start] = fmt.Sprintf("%d", val)
	if end < len(tokens)-1 {
		copy(r[start+1:len(r)], tokens[end+1:len(tokens)])
	}
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

//==============================================================================
type node interface {
	add(node) node
	calculate() int
}

//==============================================================================
type operand struct {
	op string
}

func (n *operand) add(other node) node {
	return nil
}

func (n *operand) calculate() int {
	return toInt(n.op)
}

//==============================================================================
type plus struct {
	left  node
	right node
}

func (n *plus) add(other node) node {
	if n.left == nil {
		n.left = other
		return n
	}
	r := n.left.add(other)
	if r != nil {
		return r
	}
	if n.right == nil {
		n.right = other
		return n
	}
	return n.right.add(other)
}

func (n *plus) calculate() int {
	return n.left.calculate() + n.right.calculate()
}

//==============================================================================
type mult struct {
	left  node
	right node
}

func (n *mult) add(other node) node {
	if n.left == nil {
		n.left = other
		return n
	}
	r := n.left.add(other)
	if r != nil {
		return r
	}
	if n.right == nil {
		n.right = other
		return n
	}
	return n.right.add(other)
}

func (n *mult) calculate() int {
	return n.left.calculate() * n.right.calculate()
}

//==============================================================================
func newNode(op string) node {
	var n node
	switch op {
	case "+":
		n = &plus{}
	case "*":
		n = &mult{}
	default:
		n = &operand{op: op}
	}
	return n
}

func build(tokens []string) (node, int) {
	var top node
	var i int
	for i = 0; i < len(tokens); i++ {
		t := tokens[i]
		if t == "(" {
			n, diff := build(tokens[i+1 : len(tokens)])
			if top == nil {
				top = n
			} else {
				_ = top.add(n)
			}
			i += diff + 1
		} else if t == ")" {
			return top, i
		} else {
			n := newNode(t)
			if top == nil {
				top = n
			} else {
				r := top.add(n)
				if r == nil {
					n.add(top)
					top = n
				}
			}
		}
	}
	return top, i
}

func CalculateWithTree(in string) int {
	top, _ := build(tokenize(in))
	return top.calculate()
}
