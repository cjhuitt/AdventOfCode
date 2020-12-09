package bags

import (
	"strconv"
	"strings"
)

type bag struct {
	style    string
	contents map[string]int
}

func parseStyle(in string) string {
	return strings.TrimSuffix(in, " bags")
}

func parseContents(in string) map[string]int {
	if strings.Contains(in, " no ") {
		return map[string]int{}
	}

	parts := strings.Fields(in)
	count, err := strconv.Atoi(parts[0])
	if err != nil {
		return map[string]int{}
	}

	style := strings.Join(parts[1:len(parts)-1], " ")
	return map[string]int{style: count}
}

func Parse(in string) bag {
	parts := strings.Split(in, " contain ")
	if len(parts) != 2 {
		return bag{}
	}

	return bag{style: parseStyle(parts[0]), contents: parseContents(parts[1])}
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
