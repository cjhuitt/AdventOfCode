package tickets

import (
	"strconv"
	"strings"
)

type constraint struct {
	min, max int
}

func extractInt(in string) int {
	i, err := strconv.Atoi(in)
	if err != nil {
		return -1
	}
	return i
}

func parseConstraint(in string) constraint {
	parts := strings.Split(in, "-")
	if len(parts) != 2 {
		return constraint{-1, -1}
	}

	return constraint{extractInt(parts[0]), extractInt(parts[1])}
}

//==============================================================================
type fieldspec struct {
	name  string
	rules []constraint
}

func parseFieldSpec(in string) fieldspec {
	return fieldspec{}
}

func (f *fieldspec) Equal(other fieldspec) bool {
	if f.name != other.name || len(f.rules) != len(other.rules) {
		return false
	}
	for i := 0; i < len(f.rules); i++ {
		if f.rules[i] != other.rules[i] {
			return false
		}
	}

	return true
}
