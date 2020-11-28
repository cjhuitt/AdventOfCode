package orbit

import "testing"

func TestParsing(t *testing.T) {
	tests := []struct {
		input       string
		want_center string
		want_id     string
		errors      bool
	}{
		{input: "", want_center: "", want_id: "", errors: true},
		{input: "COM)B", want_center: "COM", want_id: "B", errors: false},
		{input: "B)C", want_center: "B", want_id: "C", errors: false},
		{input: "D)I", want_center: "D", want_id: "I", errors: false},
		{input: "ASDF", want_center: "", want_id: "", errors: true},
	}
	for i, tc := range tests {
		c, id, err := Parse(tc.input)
		if tc.errors && err == nil {
			t.Errorf("Expected parsing %#v to error, got %#v, %#v (case %d)", tc.input, c, id, i)
		} else if !tc.errors && err != nil {
			t.Errorf("Expected parsing %#v to succeed, got %v (case %d)", tc.input, err, i)
		} else if c != tc.want_center || id != tc.want_id {
			t.Errorf("Expected Parse(%v) to give (%#v, %#v), got (%#v, %#v) (case %d)", tc.input, tc.want_center, tc.want_id, c, id, i)
		}
	}
}

func TestStepsTo(t *testing.T) {
	tests := []struct {
		input []string
		from  string
		to    string
		want  int
	}{
		{input: []string{"COM)B", "B)C"}, from: "C", to: "COM", want: 2},
	}
	for i, tc := range tests {
		chart := bodylist{}
		chart["COM"] = NewBody("", "COM")
		for _, in := range tc.input {
			c, id, err := Parse(in)
			if err != nil {
				t.Errorf("Error on input %#v (case %d)", in, i)
			}
			chart[id] = NewBody(c, id)
		}
		for _, b := range chart {
			b.orbiting = chart[b.orbits]
		}
		got := chart[tc.from].StepsTo(chart[tc.to])
		if got != tc.want {
			t.Errorf("Expected (%#v).StepsTo(%#v) to be %d, got %d (case %d)", tc.from, tc.to, tc.want, got, i)
		}
	}
}
