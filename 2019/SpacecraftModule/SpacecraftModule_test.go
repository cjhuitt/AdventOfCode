package SpacecraftModule

import "testing"

// The default module has zero mass
func TestDefaultModuleRequiresNoFuel(t *testing.T) {
    got := Default().FuelRequired()
    if got != 0 {
        t.Errorf("Default().FuelRequired() = %d, want 0", got)
    }
}

func TestFuelCalculation(t *testing.T) {
    tests := []struct {
        mass int
        want int
    }{
        {mass:12, want:2},
        {mass:14, want:2},
        {mass:1969, want:654},
        {mass:100756, want:33583},
    }
    for _, tc := range tests {
        got := New(tc.mass).FuelRequired()
        if got != tc.want {
            t.Errorf("New(%d).FuelRequired() = %d, want %d", tc.mass, got, tc.want)
        }
    }
}
