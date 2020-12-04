package passports

import "strings"

type passport struct {
	byr, iyr, eyr, hgt, hcl, ecl, pid, cid string
}

func Parse(in string) passport {
	parts := strings.Split(in, ":")
	if len(parts) != 2 {
		return passport{}
	}

	pp := passport{}
	switch parts[0] {
	case "iyr":
		pp.iyr = parts[1]
	}

	return pp
}
