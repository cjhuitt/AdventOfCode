package grid

import "testing"

func TestReadRow(t *testing.T) {
	tests := []struct {
		input string
		want  []*seat
	}{
		{input: "", want: []*seat{}},
		{input: "L", want: []*seat{newSeat('L')}},
		{input: ".", want: []*seat{newSeat('.')}},
		{input: "L.L", want: []*seat{newSeat('L'), newSeat('.'), newSeat('L')}},
		{input: "#.L", want: []*seat{newSeat('#'), newSeat('.'), newSeat('L')}},
		{input: "G.L", want: []*seat{}},
	}
	for i, tc := range tests {
		got := readRow(tc.input)
		if !seatSlicesEqual(got, tc.want) {
			t.Errorf("Expected readRow(%v) to result in %v, received %v (case %d)", tc.input, tc.want, got, i)
		}
	}
}

func TestReadSeating(t *testing.T) {
	tests := []struct {
		input []string
		want  deck
	}{
		{input: []string{}, want: deck{}},
		{input: []string{"L", "L"}, want: deck{[][]*seat{
			[]*seat{newSeat('L')},
			[]*seat{newSeat('L')}}, 1, 2}},
		{input: []string{"L", "G"}, want: deck{}},
		{input: []string{"L", "L."}, want: deck{}},
	}
	for i, tc := range tests {
		got := readSeating(tc.input)
		if !got.isEqualTo(tc.want) {
			t.Errorf("Expected readSeating(%v) to result in %v, received %v (case %d)", tc.input, tc.want, got, i)
		}
	}
}

func TestParse(t *testing.T) {
	tests := []struct {
		input []string
		want  deck
	}{
		{input: []string{}, want: deck{}},
		{input: []string{"L"}, want: deck{[][]*seat{
			[]*seat{newSeat('L')}}, 1, 1}},
	}
	for i, tc := range tests {
		got := Parse(tc.input)
		if !got.isEqualTo(tc.want) {
			t.Errorf("Expected Parse(%v) to result in %v, received %v (case %d)", tc.input, tc.want, got, i)
		}
	}
}

func TestStep(t *testing.T) {
	tests := []struct {
		input []string
		want  string
	}{
		{input: []string{}, want: ""},
		{input: []string{"L"}, want: "#\n"},
		{input: []string{"#"}, want: "#\n"},
		{input: []string{"."}, want: ".\n"},
	}
	for i, tc := range tests {
		d := Parse(tc.input)
		d.Step()
		got := d.Printable()
		if got != tc.want {
			t.Errorf("Expected stepping %v to result in %#v, received %#v (case %d)", tc.input, tc.want, got, i)
		}
	}
}
