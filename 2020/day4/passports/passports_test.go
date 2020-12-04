package passports

import "testing"

func TestParseInfoString(t *testing.T) {
	tests := []struct {
		input string
		want  passport
	}{
		// empty
		{input: "", want: passport{}},

		// Individual items
		{input: "byr:1937", want: passport{byr: "1937"}},
		{input: "iyr:2017", want: passport{iyr: "2017"}},
		{input: "eyr:2020", want: passport{eyr: "2020"}},
		{input: "hgt:183cm", want: passport{hgt: "183cm"}},
		{input: "hcl:#fffffd", want: passport{hcl: "#fffffd"}},
		{input: "ecl:gry", want: passport{ecl: "gry"}},
		{input: "pid:860033327", want: passport{pid: "860033327"}},
		{input: "cid:147", want: passport{cid: "147"}},

		// combination
		{input: "ecl:gry cid:147", want: passport{ecl: "gry", cid: "147"}},
		{input: "iyr:2013 ecl:amb cid:350 eyr:2023 pid:028048884 hcl:#cfa07d byr:1929",
			want: passport{iyr: "2013", ecl: "amb", cid: "350", eyr: "2023", pid: "028048884", hcl: "#cfa07d", byr: "1929"}},
	}
	for i, tc := range tests {
		got := Parse(tc.input)
		if got != tc.want {
			t.Errorf("Expected Parse(%v) to result in %v, received %v (case %d)", tc.input, tc.want, got, i)
		}
	}
}

func TestValidity(t *testing.T) {
	tests := []struct {
		input string
		want  bool
	}{
		// empty
		{input: "", want: false},

		{input: "ecl:gry pid:860033327 eyr:2020 hcl:#fffffd byr:1937 iyr:2017 cid:147 hgt:183cm", want: true},
		{input: "iyr:2013 ecl:amb cid:350 eyr:2023 pid:028048884 hcl:#cfa07d byr:1929", want: false},
		{input: "hcl:#ae17e1 iyr:2013 eyr:2024	ecl:brn pid:760753108 byr:1931 hgt:179cm", want: true},
		{input: "hcl:#cfa07d eyr:2025 pid:166559648 iyr:2011 ecl:brn hgt:59in", want: false},
	}
	for i, tc := range tests {
		got := Parse(tc.input).IsValid()
		if got != tc.want {
			t.Errorf("Expected Parse(%v).IsValid() to be %v, received %v (case %d)", tc.input, tc.want, got, i)
		}
	}
}

func TestBirthYearValidity(t *testing.T) {
	tests := []struct {
		input string
		want  bool
	}{
		// empty
		{input: "", want: false},

		{input: "1919", want: false},
		{input: "1920", want: true},
		{input: "2002", want: true},
		{input: "2003", want: false},
	}
	for i, tc := range tests {
		got := isByrValid(tc.input)
		if got != tc.want {
			t.Errorf("Expected IsByrValid(%v) to be %v, received %v (case %d)", tc.input, tc.want, got, i)
		}
	}
}
