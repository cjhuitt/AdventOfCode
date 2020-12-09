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
	}
	return p
}
