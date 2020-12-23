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

func (c *constraint) isValid() bool {
	return c.min >= 0 && c.max >= 0
}

//==============================================================================
type fieldspec struct {
	name  string
	rules []constraint
}

func parseName(in string) (string, string) {
	parts := strings.Split(in, ":")
	if len(parts) != 2 {
		return "", ""
	}
	return parts[0], parts[1]
}

func parseConstraints(in string) []constraint {
	r := []constraint{}
	parts := strings.Split(in, " ")
	for _, p := range parts {
		if p != "or" {
			c := parseConstraint(p)
			if c.isValid() {
				r = append(r, c)
			}
		}
	}
	return r
}

func parseFieldSpec(in string) fieldspec {
	f := fieldspec{}
	name, other := parseName(in)
	f.name = name
	f.rules = parseConstraints(other)
	return f
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

//==============================================================================
type ticket struct {
	fields []int
}

func parseTicket(in string) ticket {
	t := ticket{}
	parts := strings.Split(in, ",")
	if len(parts) <= 1 {
		return t
	}

	for _, p := range parts {
		t.fields = append(t.fields, extractInt(p))
	}
	return t
}

func (t *ticket) Equal(other ticket) bool {
	if len(t.fields) != len(other.fields) {
		return false
	}

	for i := 0; i < len(t.fields); i++ {
		if t.fields[i] != other.fields[i] {
			return false
		}
	}

	return true
}

func (t *ticket) Validate(fields []fieldspec) (bool, int) {
	return true, 0
}
