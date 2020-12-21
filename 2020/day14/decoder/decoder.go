package decoder

type mask struct {
	allowed    uint64
	predefined uint64
}

func parseMask(in string) mask {
	var allowed, predef uint64
	for _, c := range in {
		allowed = allowed << 1
		predef = predef << 1
		switch c {
		case 'X':
			allowed += 1
		case '1':
			allowed += 0
			predef += 1
		case '0':
			allowed += 0
		}
	}

	return mask{allowed, predef}
}

func (m mask) processed(in uint64) uint64 {
	return in&m.allowed + m.predefined
}

type program struct {
	filter mask
	mem    map[uint64]uint64
}

func Program() program {
	p := program{}
	p.mem = make(map[uint64]uint64)
	return p
}

func (p *program) store(loc, val uint64) {
	p.mem[loc] = p.filter.processed(val)
}
