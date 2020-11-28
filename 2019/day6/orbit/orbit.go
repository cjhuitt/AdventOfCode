package orbit

import "errors"

type body struct {
	steps  int
	orbits *body
}

func Parse(code string) (string, string, error) {
	return "", "", errors.New("fail")
}
