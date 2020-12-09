package bags

import (
	"strings"
)

type bag struct {
	style    string
	contents map[string]int
}

func parseStyle(in string) string {
	return strings.TrimSuffix(in, " bags")
}

func Parse(in string) bag {
	parts := strings.Split(in, " contain ")
	if len(parts) != 2 {
		return bag{}
	}

	return bag{style: parseStyle(parts[0])}
}

func (b bag) isEqualTo(other bag) bool {
	if b.style != other.style || len(b.contents) != len(other.contents) {
		return false
	}

	for style, count := range b.contents {
		if val, ok := other.contents[style]; !ok || val != count {
			return false
		}
	}
	return true
}
