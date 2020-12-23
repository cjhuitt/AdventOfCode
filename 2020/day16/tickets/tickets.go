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

func (c *constraint) passes(test int) bool {
	return test >= c.min && test <= c.max
}

//==============================================================================
type FieldSpec struct {
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

func ParseFieldSpec(in string) FieldSpec {
	f := FieldSpec{}
	name, other := parseName(in)
	f.name = name
	f.rules = parseConstraints(other)
	return f
}

func (f *FieldSpec) Equal(other FieldSpec) bool {
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

func (f *FieldSpec) passes(test int) bool {
	for _, r := range f.rules {
		if r.passes(test) {
			return true
		}
	}

	return false
}

//==============================================================================
type Ticket struct {
	fields []int
}

func ParseTicket(in string) Ticket {
	t := Ticket{}
	parts := strings.Split(in, ",")
	if len(parts) <= 1 {
		return t
	}

	for _, p := range parts {
		t.fields = append(t.fields, extractInt(p))
	}
	return t
}

func (t *Ticket) Equal(other Ticket) bool {
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

func passes(test int, specs []FieldSpec) bool {
	for _, s := range specs {
		if s.passes(test) {
			return true
		}
	}

	return false
}

func (t *Ticket) Validate(specs []FieldSpec) (bool, int) {
	for _, f := range t.fields {
		if !passes(f, specs) {
			return false, f
		}
	}
	return true, 0
}
