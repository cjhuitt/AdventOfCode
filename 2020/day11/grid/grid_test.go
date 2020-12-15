package grid

import "testing"

func TestReadRow(t *testing.T) {
	tests := []struct {
		input string
		want  []seat
	}{
		{input: "", want: []seat{}},
		{input: "L", want: []seat{newSeat('L')}},
		{input: ".", want: []seat{newSeat('.')}},
		{input: "L.L", want: []seat{newSeat('L'), newSeat('.'), newSeat('L')}},
		{input: "#.L", want: []seat{}},
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
		want  state
	}{
		{input: []string{}, want: state{}},
		{input: []string{"L", "L"}, want: state{[][]seat{
			[]seat{newSeat('L')},
			[]seat{newSeat('L')}}}},
		{input: []string{"L", "G"}, want: state{}},
		{input: []string{"L", "L."}, want: state{}},
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
		want  state
	}{
		{input: []string{}, want: state{}},
	}
	for i, tc := range tests {
		got := Parse(tc.input)
		if !got.isEqualTo(tc.want) {
			t.Errorf("Expected readSeating(%v) to result in %v, received %v (case %d)", tc.input, tc.want, got, i)
		}
	}
}
