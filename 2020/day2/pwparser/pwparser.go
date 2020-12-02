package pwparser

import (
	"errors"
	"strconv"
	"strings"
)

type span struct{ min, max int }

func parseSpan(in string) (span, error) {
	parts := strings.Split(in, "-")
	if len(parts) != 2 {
		return span{}, errors.New("Invalid range specification")
	}
	min, err := strconv.Atoi(parts[0])
	if err != nil {
		return span{}, err
	}
	max, err := strconv.Atoi(parts[1])
	if err != nil {
		return span{}, err
	}
	return span{min, max}, nil
}

type rule struct {
	allowed span
	char    string
}

func parseRule(in string) (rule, error) {
	parts := strings.Split(in, ":")
	if len(parts) != 2 {
		return rule{}, errors.New("Invalid rule specification")
	}
	s, err := parseSpan(parts[0])
	if err != nil {
		return rule{}, err
	}
	return rule{s, strings.TrimSpace(parts[1])}, nil
}
