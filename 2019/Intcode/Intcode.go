package Intcode

type program struct {
	code []int
}

func Default() program {
	return program{make([]int, 0)}
}

func New(p []int) program {
	return program{p}
}

func (p program) IsEmpty() bool {
	return len(p.code) == 0
}

func (p program) IsDone() bool {
	return len(p.code) == 0
}
