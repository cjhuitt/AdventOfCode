package orbit

import (
	"errors"
	"strings"
)

type body struct {
	steps  int
	orbits *body
}

func Parse(code string) (string, string, error) {
	parts := strings.Split(code, ")")
	if len(parts) != 2 {
		return "", "", errors.New("fail")
	}

	return parts[0], parts[1], nil
}
