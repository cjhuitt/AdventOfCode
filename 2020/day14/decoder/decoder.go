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

func (p *program) execute(in string) {
	switch in[1] {
	case 'a':
		p.filter = parseMask(in)
	case 'e':
		loc, val := parseStore(in)
		p.store(loc, val)
	}
}
