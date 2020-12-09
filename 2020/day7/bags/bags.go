package bags

type bag struct {
	style    string
	contents map[string]int
}

func Parse(in string) bag {
	return bag{}
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
