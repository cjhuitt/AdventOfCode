package orbit

import "testing"

func TestErrorsParsingEmptyString(t *testing.T) {
	a, b, err := Parse("")
	if err == nil {
		t.Errorf("Expected parsing %#v to error, got %#v, %#v", "", a, b)
	}
}
