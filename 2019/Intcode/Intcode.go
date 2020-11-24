package Intcode

type program struct {
	code []int
	xp   int
}

const (
	ADD  = 1
	MULT = 2
	TERM = 99
)

func Default() program {
	return program{make([]int, 0), 0}
}

func New(p []int) program {
	return program{p, 0}
}

func (p program) IsEmpty() bool {
	return len(p.code) == 0
}

func (p program) IsDone() bool {
	return len(p.code) == 0 || p.code[p.xp] == TERM
}
