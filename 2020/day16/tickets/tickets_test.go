package tickets

import "testing"

func TestParseConstraint(t *testing.T) {
	tests := []struct {
		input string
		want  constraint
	}{
		{input: "", want: constraint{-1, -1}},
		{input: "1-3", want: constraint{min: 1, max: 3}},
	}
	for i, tc := range tests {
		got := parseConstraint(tc.input)
		if got != tc.want {
			t.Errorf("Expected parseConstraint(%v) to result in %v, received %v (case %d)", tc.input, tc.want, got, i)
		}
	}
}

func TestPassConstraint(t *testing.T) {
	c := constraint{1, 3}
	tests := []struct {
		input int
		want  bool
	}{
		{input: -1, want: false},
		{input: 0, want: false},
		{input: 1, want: true},
		{input: 2, want: true},
		{input: 3, want: true},
		{input: 4, want: false},
	}
	for i, tc := range tests {
		got := c.passes(tc.input)
		if got != tc.want {
			t.Errorf("Expected %#v.passes(%v) to be %v, received %v (case %d)", c, tc.input, tc.want, got, i)
		}
	}
}

func TestParseFieldSpec(t *testing.T) {
	tests := []struct {
		input string
		want  fieldspec
	}{
		{input: "", want: fieldspec{}},
		{input: "class: 1-3 or 5-7", want: fieldspec{"class", []constraint{constraint{1, 3}, constraint{5, 7}}}},
	}
	for i, tc := range tests {
		got := parseFieldSpec(tc.input)
		if !got.Equal(tc.want) {
			t.Errorf("Expected parseFieldSpec(%v) to result in %v, received %v (case %d)", tc.input, tc.want, got, i)
		}
	}
}

func TestPassField(t *testing.T) {
	spec := fieldspec{"class", []constraint{{1, 3}, {5, 7}}}
	tests := []struct {
		input int
		want  bool
	}{
		{input: -1, want: false},
		{input: 0, want: false},
		{input: 1, want: true},
		{input: 2, want: true},
		{input: 3, want: true},
		{input: 4, want: false},
		{input: 5, want: true},
		{input: 6, want: true},
		{input: 7, want: true},
		{input: 8, want: false},
	}
	for i, tc := range tests {
		got := spec.passes(tc.input)
		if got != tc.want {
			t.Errorf("Expected %#v.passes(%v) to be %v, received %v (case %d)", spec, tc.input, tc.want, got, i)
		}
	}
}

func TestParseTicket(t *testing.T) {
	tests := []struct {
		input string
		want  ticket
	}{
		{input: "", want: ticket{}},
		{input: "7,1,14", want: ticket{[]int{7, 1, 14}}},
	}
	for i, tc := range tests {
		got := parseTicket(tc.input)
		if !got.Equal(tc.want) {
			t.Errorf("Expected parseTicket(%v) to result in %v, received %v (case %d)", tc.input, tc.want, got, i)
		}
	}
}

func TestValidateTicket(t *testing.T) {
	specs := []fieldspec{}
	specs = append(specs, fieldspec{"class", []constraint{{1, 3}, {5, 7}}})
	specs = append(specs, fieldspec{"row", []constraint{{6, 11}, {33, 44}}})
	specs = append(specs, fieldspec{"seat", []constraint{{13, 40}, {45, 50}}})

	tests := []struct {
		input      string
		want_valid bool
		want_error int
	}{
		{input: "", want_valid: true, want_error: 0},
		{input: "7,3,47", want_valid: true, want_error: 0},
		{input: "40,4,50", want_valid: false, want_error: 4},
		{input: "55,2,20", want_valid: false, want_error: 55},
		{input: "38,6,12", want_valid: false, want_error: 12},
	}
	for i, tc := range tests {
		tkt := parseTicket(tc.input)
		got_valid, got_error := tkt.Validate(specs)
		if got_valid != tc.want_valid || got_error != tc.want_error {
			t.Errorf("Expected validating ticket (%v) to be (%v, %v), received (%v, %v) (case %d)", tc.input, tc.want_valid, tc.want_error, got_valid, got_error, i)
		}
	}
}
