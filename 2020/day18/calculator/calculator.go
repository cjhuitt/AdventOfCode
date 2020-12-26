package calculator

import (
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

func tokenize(in string) []string {
	var s scanner.Scanner
	s.Init(strings.NewReader(in))
	tokens := []string{}
	for tok := s.Scan(); tok != scanner.EOF; tok = s.Scan() {
		tokens = append(tokens, s.TokenText())
	}
	return tokens
}

type node interface {
	add(node) node
	calculate() int
	transform() node
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

func (n *operand) transform() node {
	return n
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

func (n *plus) transform() node {
	return n
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

func (n *mult) transform() node {
	return n
}

//==============================================================================
type subexpr struct {
	top node
}

func (n *subexpr) add(other node) node {
	if n.top == nil {
		n.top = other
		return n
	}
	return n.top.add(other)
}

func (n *subexpr) calculate() int {
	return n.top.calculate()
}

func (n *subexpr) transform() node {
	return n
}

//==============================================================================
func newNode(op string) node {
	switch op {
	case "+":
		return &plus{}
	case "*":
		return &mult{}
	case "(":
		return &subexpr{}
	default:
	}
	return &operand{op: op}
}

func build(tokens []string) (node, int) {
	var top node
	var i int
	for i = 0; i < len(tokens); i++ {
		t := tokens[i]
		n := newNode(t)
		if t == "(" {
			o, diff := build(tokens[i+1 : len(tokens)])
			n.add(o)
			if top == nil {
				top = n
			} else {
				_ = top.add(n)
			}
			i += diff + 1
		} else if t == ")" {
			return top, i
		} else {
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

func Calculate(in string) int {
	top, _ := build(tokenize(in))
	return top.calculate()
}

func Calculate2(in string) int {
	top, _ := build(tokenize(in))
	top = top.transform()
	return top.calculate()
}
