package grid

import "fmt"
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
			[]*seat{newSeat('L')}}, 1, 2, true}},
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
			[]*seat{newSeat('L')}}, 1, 1, true}},
		{input: []string{"L", "L"}, want: deck{[][]*seat{
			[]*seat{newSeat('L')}, []*seat{newSeat('L')}}, 1, 1, true}},
		{input: []string{"LL"}, want: deck{[][]*seat{
			[]*seat{newSeat('L'), newSeat('L')}}, 1, 1, true}},
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
		input   []string
		want    string
		changed bool
	}{
		{input: []string{}, want: "", changed: false},
		{input: []string{"L"}, want: "#\n", changed: true},
		{input: []string{"#"}, want: "#\n", changed: false},
		{input: []string{"."}, want: ".\n", changed: false},
		{input: []string{"###", "###", "###"}, want: "#L#\nLLL\n#L#\n", changed: true},
		{input: []string{"##", "#.", "##"}, want: "##\n#.\n##\n", changed: false},
		{input: []string{"##", "##", "##"}, want: "##\nLL\n##\n", changed: true},
		{input: []string{"#..#..#", ".......", ".......", "#..#..#", ".......", ".......", "#..#..#"},
			want: "#..L..#\n.......\n.......\nL..L..L\n.......\n.......\n#..L..#\n", changed: true},
	}
	for i, tc := range tests {
		fmt.Println()
		fmt.Println(i)
		d := Parse(tc.input)
		d.Step()
		got := d.Printable()
		if d.Changed() != tc.changed {
			t.Errorf("Expected stepping %v to show %v change, received %#v (case %d)", tc.input, tc.changed, d.Changed(), i)
		} else if got != tc.want {
			t.Errorf("Expected stepping %v to result in %#v, received %#v (case %d)", tc.input, tc.want, got, i)
		}
	}
}
