package decoder

import (
	"strconv"
	"strings"
	"text/scanner"
)

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

func parseStore(in string) (uint64, uint64) {
	var s scanner.Scanner
	s.Init(strings.NewReader(in))
	s.Filename = "example"
	var loc, val uint64
	var err error
	loc_next, val_next := false, false
	for tok := s.Scan(); tok != scanner.EOF; tok = s.Scan() {
		t := s.TokenText()
		if t == "[" {
			loc_next = true
		} else if t == "=" {
			val_next = true
		} else if loc_next {
			loc_next = false
			loc, err = strconv.ParseUint(t, 10, 64)
			if err != nil {
				return 0, 0
			}
		} else if val_next {
			val_next = false
			val, err = strconv.ParseUint(t, 10, 64)
			if err != nil {
				return 0, 0
			}
		}
	}
	return loc, val
}

func (m mask) processed(in uint64) uint64 {
	return in&m.allowed + m.predefined
}

type loc_mask struct {
	base   uint64
	varies []int
}

func parseLocMask(in string) loc_mask {
	var l loc_mask
	t := len(in) - 1
	for i, c := range in {
		l.base = l.base << 1
		switch c {
		case 'X':
			l.varies = append(l.varies, t-i)
		case '1':
			l.base += 1
		}
	}

	return l
}

func getMemLocs(base uint64, varying []int) []uint64 {
	if len(varying) == 0 {
		return []uint64{base}
	}
	orig := getMemLocs(base, varying[1:])
	flip := []uint64{}
	f := uint64(1) << varying[0]
	for _, m := range orig {
		flip = append(flip, m|f)
	}
	return append(orig, flip...)
}

func (l *loc_mask) set(mem map[uint64]uint64, val uint64) {
	for _, m := range getMemLocs(l.base, l.varies) {
		mem[m] = val
	}
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

func (p *program) Execute(in string) {
	switch in[1] {
	case 'a':
		p.filter = parseMask(in)
	case 'e':
		loc, val := parseStore(in)
		p.store(loc, val)
	}
}

func (p *program) SumMemory() uint64 {
	var sum uint64
	for _, val := range p.mem {
		sum += val
	}

	return sum
}
