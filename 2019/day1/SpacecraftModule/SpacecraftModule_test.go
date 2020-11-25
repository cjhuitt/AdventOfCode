package SpacecraftModule

import "testing"

// The default module has zero mass
func TestDefaultModuleRequiresNoFuel(t *testing.T) {
    got := Default().BaseFuelRequired()
    if got != 0 {
        t.Errorf("Default().BaseFuelRequired() = %d, want 0", got)
    }
}

func TestBaseFuelCalculation(t *testing.T) {
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
        got := NewModule(tc.mass).BaseFuelRequired()
        if got != tc.want {
            t.Errorf("New(%d).BaseFuelRequired() = %d, want %d", tc.mass, got, tc.want)
        }
    }
}

func TestMultipleModuleInitialization(t *testing.T) {
    masses := []int{10, 14}
    modules := NewModules(masses)
    if len(modules) != 2 {
        t.Errorf("len(modules) = %d, want 2", len(modules))
    }

    got := modules[0].BaseFuelRequired() + modules[1].BaseFuelRequired()
    if got != 3 {
        t.Errorf("sum(modules BaseFuelRequired) = %d, want 3", got)
    }
}

func TestTotalFuelCalculation(t *testing.T) {
    tests := []struct {
        mass int
        want int
    }{
        {mass:14, want:2},
        {mass:1969, want:966},
        {mass:100756, want:50346},
    }
    for _, tc := range tests {
        got := NewModule(tc.mass).TotalFuelRequired()
        if got != tc.want {
            t.Errorf("New(%d).TotalFuelRequired() = %d, want %d", tc.mass, got, tc.want)
        }
    }
}

