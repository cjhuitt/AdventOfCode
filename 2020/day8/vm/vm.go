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
	i := p.pos
	for i >= 0 {
		if trace[i] > 0 {
			return false, p
		}
		trace[i] += 1
		p = p.Step()
		i = p.pos
	}
	return true, p
}

func (p program) Accumulator() int {
	return p.acc
}

func (p program) dup() program {
	r := program{}
	r.code = make([]opcode, len(p.code))
	_ = copy(r.code, p.code)
	r.pos = p.pos
	r.acc = p.acc
	return r
}

func (p program) FixNextNop(start int) (int, program, bool) {
	for i := start; i < len(p.code); i++ {
		if p.code[i].op == "nop" {
			r := p.dup()
			r.code[i].op = "jmp"
			return i, r, true
		}
	}
	return -1, p, false
}

func (p program) FixNextJmp(start int) (int, program, bool) {
	for i := start; i < len(p.code); i++ {
		if p.code[i].op == "jmp" {
			r := p.dup()
			r.code[i].op = "nop"
			return i, r, true
		}
	}
	return -1, p, false
}
