package Intcode

type program struct {
	// the opcode/data store (intermixed)
	stack []int
	// the execution point
	xp int
	// The most recent output (if any)
	output *int
}

const (
	ADD  = 1
	MULT = 2
	TERM = 99
)

func Default() program {
	return program{make([]int, 0), -1, nil}
}

func New(p []int) program {
	return program{p, 0, nil}
}

func (p program) IsEmpty() bool {
	return len(p.stack) == 0
}

func (p program) IsDone() bool {
	return p.xp < 0 || p.xp >= len(p.stack) ||
		p.stack[p.xp] == TERM
}

// Step through execution of one opcode based on the current execution point
func (p program) Step() program {
	if p.xp < len(p.stack) && p.xp >= 0 {
		switch p.stack[p.xp] {
		case ADD:
			return add(p.stack, p.xp)
		case MULT:
			return mult(p.stack, p.xp)
		case TERM:
			break
		default:
			return program{p.stack, -1, nil}
		}
	}
	return p
}

// Execute the intcode program
func (p program) Execute() program {
	temp := p
	for !temp.IsDone() {
		temp = temp.Step()
	}
	return temp
}

func (p program) Data() []int {
	return p.stack
}

func (p program) Output() (int, error) {
	return 0, nil
}

func add(stack []int, xp int) program {
	if xp+3 >= len(stack) {
		return program{stack, -1, nil}
	}
	add1 := stack[xp+1]
	add2 := stack[xp+2]
	loc := stack[xp+3]
	if add1 >= len(stack) || add1 < 0 ||
		add2 >= len(stack) || add2 < 0 ||
		loc >= len(stack) || loc < 0 {
		return program{stack, -1, nil}
	}
	sum := stack[add1] + stack[add2]
	new_stack := stack
	new_stack[loc] = sum
	return program{new_stack, xp + 4, nil}
}

func mult(stack []int, xp int) program {
	if xp+3 >= len(stack) {
		return program{stack, -1, nil}
	}
	mult1 := stack[xp+1]
	mult2 := stack[xp+2]
	loc := stack[xp+3]
	if mult1 >= len(stack) || mult1 < 0 ||
		mult2 >= len(stack) || mult2 < 0 ||
		loc >= len(stack) || loc < 0 {
		return program{stack, -1, nil}
	}
	mult := stack[mult1] * stack[mult2]
	new_stack := stack
	new_stack[loc] = mult
	return program{new_stack, xp + 4, nil}
}
