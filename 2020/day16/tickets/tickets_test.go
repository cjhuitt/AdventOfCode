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

func TestValidateFields(t *testing.T) {
	c := constraint{1, 3}
	tests := []struct {
		input      []int
		want_valid bool
		want_error int
	}{
		{input: []int{}, want_valid: true, want_error: 0},
	}
	for i, tc := range tests {
		got_valid, got_error := c.validateFields(tc.input)
		if got_valid != tc.want_valid || got_error != tc.want_error {
			t.Errorf("Expected validating fields (%v) to be (%v, %v), received (%v, %v) (case %d)", tc.input, tc.want_valid, tc.want_error, got_valid, got_error, i)
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
	fields := []fieldspec{}
	fields = append(fields, fieldspec{"class", []constraint{{1, 3}, {5, 7}}})
	fields = append(fields, fieldspec{"row", []constraint{{6, 11}, {33, 44}}})
	fields = append(fields, fieldspec{"seat", []constraint{{13, 40}, {45, 50}}})

	tests := []struct {
		input      string
		want_valid bool
		want_error int
	}{
		{input: "", want_valid: true, want_error: 0},
	}
	for i, tc := range tests {
		tkt := parseTicket(tc.input)
		got_valid, got_error := tkt.Validate(fields)
		if got_valid != tc.want_valid || got_error != tc.want_error {
			t.Errorf("Expected validating ticket (%v) to be (%v, %v), received (%v, %v) (case %d)", tc.input, tc.want_valid, tc.want_error, got_valid, got_error, i)
		}
	}
}
