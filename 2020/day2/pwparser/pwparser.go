package pwparser

import (
	"errors"
	"strconv"
	"strings"
)

type pair struct{ min, max int }

func parsePair(in string) (pair, error) {
	parts := strings.Split(in, "-")
	if len(parts) != 2 {
		return pair{}, errors.New("Invalid range specification")
	}
	min, err := strconv.Atoi(parts[0])
	if err != nil {
		return pair{}, err
	}
	max, err := strconv.Atoi(parts[1])
	if err != nil {
		return pair{}, err
	}
	return pair{min, max}, nil
}

type rule struct {
	allowed pair
	char    string
}

func ParseRule(in string) (rule, error) {
	parts := strings.Split(in, " ")
	if len(parts) != 2 {
		return rule{}, errors.New("Invalid rule specification")
	}
	s, err := parsePair(parts[0])
	if err != nil {
		return rule{}, err
	}
	return rule{s, strings.TrimSpace(parts[1])}, nil
}

func (r rule) Matches(pw string) bool {
	c := strings.Count(pw, r.char)
	return c >= r.allowed.min && c <= r.allowed.max
}
