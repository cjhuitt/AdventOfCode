package Intcode

type program struct {
	// the opcode/data store (intermixed)
	stack []int
	// the execution point
	xp int
	// The output (if any)
	output *int
	// Where to store input (if requested)
	input *int
}

const (
	ADD  = 1
	MULT = 2
	INP  = 3
	OUTP = 4
	TERM = 99
)

func invalid(stack []int) program {
	return program{stack, -1, nil, nil}
}

func Default() program {
	return program{make([]int, 0), -1, nil, nil}
}

func New(p []int) program {
	return program{p, 0, nil, nil}
}

func (p program) IsEmpty() bool {
	return len(p.stack) == 0
}

func (p program) IsDone() bool {
	return p.xp < 0 || p.xp >= len(p.stack) ||
		p.stack[p.xp] == TERM
}

func (p program) IsPaused() bool {
	return (p.output != nil || p.input != nil) && !p.IsDone()
}

// Step through execution of one opcode based on the current execution point
func (p program) Step() program {
	if p.xp < len(p.stack) && p.xp >= 0 {
		switch p.stack[p.xp] {
		case ADD:
			return add(p.stack, p.xp)
		case MULT:
			return mult(p.stack, p.xp)
		case INP:
			return in(p.stack, p.xp, p.input)
		case OUTP:
			return out(p.stack, p.xp)
		case TERM:
			break
		default:
			return invalid(p.stack)
		}
	}
	return p
}

// Execute the intcode program
func (p program) Execute() program {
	temp := p
	for !temp.IsDone() && !temp.IsPaused() {
		temp = temp.Step()
	}
	return temp
}

func (p program) Data() []int {
	return p.stack
}

func (p program) Output() *int {
	return p.output
}

func (p program) WithInput(input *int) program {
	if p.input == nil {
		return invalid(p.stack)
	}
	return program{p.stack, p.xp, nil, input}
}

func add(stack []int, xp int) program {
	if xp+3 >= len(stack) {
		return invalid(stack)
	}
	add1 := stack[xp+1]
	add2 := stack[xp+2]
	loc := stack[xp+3]
	if add1 >= len(stack) || add1 < 0 ||
		add2 >= len(stack) || add2 < 0 ||
		loc >= len(stack) || loc < 0 {
		return invalid(stack)
	}
	sum := stack[add1] + stack[add2]
	new_stack := stack
	new_stack[loc] = sum
	return program{new_stack, xp + 4, nil, nil}
}

func mult(stack []int, xp int) program {
	if xp+3 >= len(stack) {
		return invalid(stack)
	}
	mult1 := stack[xp+1]
	mult2 := stack[xp+2]
	loc := stack[xp+3]
	if mult1 >= len(stack) || mult1 < 0 ||
		mult2 >= len(stack) || mult2 < 0 ||
		loc >= len(stack) || loc < 0 {
		return invalid(stack)
	}
	mult := stack[mult1] * stack[mult2]
	new_stack := stack
	new_stack[loc] = mult
	return program{new_stack, xp + 4, nil, nil}
}

func in(stack []int, xp int, input *int) program {
	if xp+1 >= len(stack) {
		return invalid(stack)
	}
	loc := stack[xp+1]
	if loc >= len(stack) || loc < 0 {
		return invalid(stack)
	}

	if input == nil {
		// Prep and repeat instruction
		in := &stack[loc]
		return program{stack, xp, nil, in}
	}

	new_stack := stack
	new_stack[loc] = *input
	return program{new_stack, xp + 2, nil, nil}
}

func out(stack []int, xp int) program {
	if xp+1 >= len(stack) {
		return invalid(stack)
	}
	loc := stack[xp+1]
	if loc >= len(stack) || loc < 0 {
		return invalid(stack)
	}
	out := &stack[loc]
	return program{stack, xp + 2, out, nil}
}
