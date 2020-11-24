package SpacecraftModule

type Module struct {
    mass int
}

func Default() Module {
    return NewModule(0)
}

func NewModule(mass int) Module {
    return Module{mass}
}

func NewModules(masses []int) []Module {
    modules := make([]Module, 0)
    for _, m := range masses {
        modules = append(modules, Module{ m })
    }
    return modules
}


// The base fuel for launching the module
func (m Module) BaseFuelRequired() int {
    if m.mass < 6 {
        return 0
    }
    return (m.mass / 3) - 2
}

// The total fuel for launching the module (including the mass of the fuel)
func (mod Module) TotalFuelRequired() int {
    if mod.mass < 6 {
        return 0
    }

    m := (mod.mass / 3) - 2
    t := m
    for m >= 6 {
        m = (m / 3) - 2
        t += m
    }

    return t;
}
