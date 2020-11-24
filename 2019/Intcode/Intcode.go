package Intcode

type program struct {
	code []int
}

func New() program {
	return program{make([]int, 0)}
}

func (p program) IsEmpty() bool {
	return len(p.code) == 0
}
