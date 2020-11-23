package SpacecraftModule

type Module struct {
    mass int
}

func Default() Module {
    return New(0)
}

func New(mass int) Module {
    return Module{mass}
}

func (m Module) FuelRequired() int {
    if m.mass < 6 {
        return 0
    }
    return (m.mass / 3) - 2
}
