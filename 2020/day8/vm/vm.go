package vm

import (
	"strconv"
	"strings"
)

type opcode struct {
	op  string
	val int
}

type program struct {
	code []opcode
	pos  int
	acc  int
}

func parseOp(in string) opcode {
	parts := strings.Fields(in)
	if len(parts) != 2 {
		return opcode{}
	}

	i, err := strconv.Atoi(parts[1])
	if err != nil {
		return opcode{}
	}

	return opcode{parts[0], i}
}

func Parse(in []string) program {
	p := program{}
	for _, line := range in {
		op := parseOp(line)
		if op.op != "" {
			p.code = append(p.code, op)
		}
	}

	return p
}

func (p program) Step() program {
	if p.pos < 0 || p.pos >= len(p.code) {
		p.pos = -1
		return p
	}

	op := p.code[p.pos]
	switch op.op {
	case "jmp":
		p.pos += op.val
	case "acc":
		p.acc += op.val
		p.pos++
	default:
		p.pos++
	}
	return p
}

func (p program) Execute() (bool, program) {
	trace := map[int]int{}
	instruction := p.Position()
	for instruction >= 0 {
		if trace[instruction] > 0 {
			break
		}
		trace[instruction] += 1
		p = p.Step()
		instruction = p.Position()
	}
	return false, p
}

func (p program) Position() int {
	return p.pos
}

func (p program) Accumulator() int {
	return p.acc
}
