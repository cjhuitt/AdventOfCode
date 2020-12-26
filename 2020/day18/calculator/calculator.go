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
type everythingNode struct {
	left  node
	op    string
	right node
}

func (n *everythingNode) add(other node) node {
	switch n.op {
	case "*", "+":
	default:
		return nil
	}
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

func (n *everythingNode) calculate() int {
	switch n.op {
	case "+":
		return n.left.calculate() + n.right.calculate()
	case "*":
		return n.left.calculate() * n.right.calculate()
	}
	return toInt(n.op)
}

func newNode(op string) node {
	var n node
	switch op {
	case "+":
		n = &plus{}
	case "*":
		n = &everythingNode{op: op}
	default:
		n = &operand{op: op}
	}
	return n
}

func build(tokens []string) node {
	var top node
	for _, t := range tokens {
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
	return top
}

func CalculateWithTree(in string) int {
	top := build(tokenize(in))
	return top.calculate()
}
