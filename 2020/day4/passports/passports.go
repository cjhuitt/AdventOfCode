package passports

import (
	"regexp"
	"strconv"
	"strings"
)

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

func (pp passport) IsValid() bool {
	return isByrValid(pp.byr) && isIyrValid(pp.iyr) && isEyrValid(pp.eyr) &&
		isHgtValid(pp.hgt) && isHclValid(pp.hcl) && pp.ecl != "" && pp.pid != ""
}

func isByrValid(in string) bool {
	matched, err := regexp.MatchString(`\d\d\d\d`, in)
	if err != nil || !matched {
		return false
	}
	i, err := strconv.Atoi(in)
	if err != nil {
		return false
	}
	return i >= 1920 && i <= 2002
}

func isIyrValid(in string) bool {
	matched, err := regexp.MatchString(`20\d\d`, in)
	if err != nil || !matched {
		return false
	}
	i, err := strconv.Atoi(in)
	if err != nil {
		return false
	}
	return i >= 2010 && i <= 2020
}

func isEyrValid(in string) bool {
	matched, err := regexp.MatchString(`20\d\d`, in)
	if err != nil || !matched {
		return false
	}
	i, err := strconv.Atoi(in)
	if err != nil {
		return false
	}
	return i >= 2020 && i <= 2030
}

func isHgtValid(in string) bool {
	metric, err := regexp.MatchString(`1\d\dcm`, in)
	if err != nil {
		return false
	} else if metric {
		cm, err := strconv.Atoi(in[0:3])
		if err != nil {
			return false
		}
		return cm >= 150 && cm <= 193
	}
	imp, err := regexp.MatchString(`\d\din`, in)
	if err != nil {
		return false
	} else if imp {
		inches, err := strconv.Atoi(in[0:2])
		if err != nil {
			return false
		}
		return inches >= 59 && inches <= 76
	}
	return false
}

func isHclValid(in string) bool {
	matched, err := regexp.MatchString(`#[0-9a-f]{6}$`, in)
	return err == nil && matched
}
