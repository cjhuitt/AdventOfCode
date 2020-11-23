package fuel

import "testing"

func TestZero(t *testing.T) {
    got := FuelFor(0)
    if got != 0 {
        t.Errorf("FuelFor(0) = %d, want 1", got)
    }
}

func TestCalculation(t *testing.T) {
    tests := []struct {
        input int
        want int
    }{
        {input:12, want:2},
        {input:14, want:2},
        {input:1969, want:654},
        {input:100756, want:33583},
    }
    for _, tc := range tests {
        got := FuelFor(tc.input)
        if got != tc.want {
            t.Errorf("FuelFor(%d) = %d, want %d", tc.input, got, tc.want)
        }
    }
}
