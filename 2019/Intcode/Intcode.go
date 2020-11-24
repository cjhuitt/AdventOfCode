package Intcode

type program struct {
	// the opcode/data store (intermixed)
	stack []int
	// the execution point
	xp int
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
	return len(p.stack) == 0
}

func (p program) IsDone() bool {
	return len(p.stack) == 0 || p.stack[p.xp] == TERM
}

// Step through execution of one opcode based on the current execution point
func (p program) Step() program {
	if p.xp >= len(p.stack) {
		return p
	}
	switch p.stack[p.xp] {
	case ADD:
		return add(p.stack, p.xp)
	case TERM:
		break
	}
	return p
}

func (p program) Data() []int {
	return p.stack
}

func add(stack []int, xp int) program {
	if xp+3 >= len(stack) {
		return program{stack, xp}
	}
	sum := stack[xp+1] + stack[xp+2]
	loc := stack[xp+3] - 1 // convert 1-based to 0-based
	if loc >= len(stack) || loc < 0 {
		return program{stack, xp}
	}
	new_stack := stack
	new_stack[loc] = sum
	return program{new_stack, xp + 4}
}
