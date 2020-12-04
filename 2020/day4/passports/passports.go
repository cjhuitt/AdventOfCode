package passports

import "strings"

type passport struct {
	byr, iyr, eyr, hgt, hcl, ecl, pid, cid string
}

func Parse(in string) passport {
	pp := passport{}
	for _, field := range strings.Fields(in) {
		parts := strings.Split(field, ":")
		if len(parts) != 2 {
			return passport{}
		}

		switch parts[0] {
		case "byr":
			pp.byr = parts[1]
		case "iyr":
			pp.iyr = parts[1]
		case "eyr":
			pp.eyr = parts[1]
		case "hgt":
			pp.hgt = parts[1]
		case "hcl":
			pp.hcl = parts[1]
		case "ecl":
			pp.ecl = parts[1]
		case "pid":
			pp.pid = parts[1]
		case "cid":
			pp.cid = parts[1]
		}
	}

	return pp
}
