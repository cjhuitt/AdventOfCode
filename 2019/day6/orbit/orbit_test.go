package orbit

import "testing"

func TestErrorsParsingEmptyString(t *testing.T) {
	c, id, err := Parse("")
	if err == nil {
		t.Errorf("Expected parsing %#v to error, got %#v, %#v", "", c, id)
	}
}

func TestParsing(t *testing.T) {
	tests := []struct {
		input       string
		want_center string
		want_id     string
		errors      bool
	}{
		{input: "COM)B", want_center: "COM", want_id: "B", errors: false},
	}
	for i, tc := range tests {
		c, id, err := Parse(tc.input)
		if tc.errors && err == nil {
			t.Errorf("Expected parsing %#v to error, got %#v, %#v (case %d)", tc.input, c, id, i)
		}
		if !tc.errors && err != nil {
			t.Errorf("Expected parsing %#v to succeed, got %v (case %d)", tc.input, err, i)
		}
		if c != tc.want_center || id != tc.want_id {
			t.Errorf("Expected Parse(%v) to give (%#v, %#v), got (%#v, %#v) (case %d)", tc.input, tc.want_center, tc.want_id, c, id, i)
		}
	}
}
