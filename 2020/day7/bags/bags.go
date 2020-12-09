package bags

import (
	"strconv"
	"strings"
)

type parsedBag struct {
	style    string
	contents map[string]int
}

type bag struct {
	style    string
	contents map[*bag]int
}

type bagSpecList struct {
	contents map[string]*bag
}

func parseStyle(in string) string {
	return strings.TrimSuffix(in, " bags")
}

func parseContent(in string) (string, int, error) {
	parts := strings.Fields(in)
	count, err := strconv.Atoi(parts[0])
	if err != nil {
		return "", -1, err
	}

	style := strings.Join(parts[1:len(parts)-1], " ")
	return style, count, nil
}

func parseContents(in string) map[string]int {
	if strings.Contains(in, " no ") {
		return map[string]int{}
	}

	contents := make(map[string]int)
	for _, s := range strings.Split(in, ", ") {
		style, count, err := parseContent(s)
		if err != nil {
			return map[string]int{}
		}

		contents[style] = count
	}

	return contents
}

func ParseConstraint(in string) parsedBag {
	parts := strings.Split(in, " contain ")
	if len(parts) != 2 {
		return parsedBag{}
	}

	return parsedBag{style: parseStyle(parts[0]), contents: parseContents(parts[1])}
}

func (b parsedBag) isEqualTo(other parsedBag) bool {
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
