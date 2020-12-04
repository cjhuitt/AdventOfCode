package passports

type passport struct {
	byr, iyr, eyr, hgt, hcl, ecl, pid, cid string
}

func Parse(in string) passport {
	return passport{}
}
