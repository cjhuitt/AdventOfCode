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


func (m Module) FuelRequired() int {
    if m.mass < 6 {
        return 0
    }
    return (m.mass / 3) - 2
}
