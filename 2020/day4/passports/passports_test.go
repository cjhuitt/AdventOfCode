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
	}
	for i, tc := range tests {
		got := Parse(tc.input)
		if got != tc.want {
			t.Errorf("Expected Parse(%v) to result in %v, received %v (case %d)", tc.input, tc.want, got, i)
		}
	}
}
