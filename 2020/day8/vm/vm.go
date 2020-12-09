package vm

import (
	"strconv"
	"strings"
)

func parseOp(in string) (string, int) {
	parts := strings.Fields(in)
	if len(parts) != 2 {
		return "", -1
	}

	i, err := strconv.Atoi(parts[1])
	if err != nil {
		return "", -1
	}

	return parts[0], i
}
