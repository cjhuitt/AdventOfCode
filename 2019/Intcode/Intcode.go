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
		switch p.stack[p.xp] % 100 {
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

func cutoff(pnum int) (int, int) {
	cut := 100
	for i := 1; i < pnum; i++ {
		cut = cut * 10
	}
	return cut, cut * 10
}

func find_read(stack []int, opcode, pos, pnum int) *int {
	if pos+pnum >= len(stack) {
		return nil
	}

	cut, mod := cutoff(pnum)
	if (opcode % mod) > cut {
		return &stack[pos+pnum]
	}

	loc := stack[pos+pnum]
	if loc >= len(stack) || loc < 0 {
		return nil
	}
	return &stack[loc]
}

func find_write(stack []int, opcode, pos, pnum int) *int {
	if pos+pnum >= len(stack) {
		return nil
	}

	threshold, _ := cutoff(pnum)
	if opcode > threshold {
		return nil
	}

	loc := stack[pos+pnum]
	if loc >= len(stack) || loc < 0 {
		return nil
	}
	return &stack[loc]
}

func add(stack []int, xp int) program {
	opcode := stack[xp]
	new_stack := stack

	read1 := find_read(stack, opcode, xp, 1)
	read2 := find_read(stack, opcode, xp, 2)
	write := find_write(new_stack, opcode, xp, 3)
	if read1 == nil || read2 == nil || write == nil {
		return invalid(stack)
	}

	*write = *read1 + *read2
	return program{new_stack, xp + 4, nil, nil}
}

func mult(stack []int, xp int) program {
	opcode := stack[xp]
	new_stack := stack

	read1 := find_read(stack, opcode, xp, 1)
	read2 := find_read(stack, opcode, xp, 2)
	write := find_write(new_stack, opcode, xp, 3)
	if read1 == nil || read2 == nil || write == nil {
		return invalid(stack)
	}

	*write = *read1 * *read2
	return program{new_stack, xp + 4, nil, nil}
}

func in(stack []int, xp int, input *int) program {
	opcode := stack[xp]

	if input == nil {
		// Prep and repeat instruction
		write := find_write(stack, opcode, xp, 1)
		if write == nil {
			return invalid(stack)
		}

		return program{stack, xp, nil, write}
	}

	new_stack := stack
	write := find_write(new_stack, opcode, xp, 1)
	if write == nil {
		return invalid(stack)
	}
	*write = *input
	return program{new_stack, xp + 2, nil, nil}
}

func out(stack []int, xp int) program {
	opcode := stack[xp]

	out := find_read(stack, opcode, xp, 1)
	if out == nil {
		return invalid(stack)
	}

	return program{stack, xp + 2, out, nil}
}
