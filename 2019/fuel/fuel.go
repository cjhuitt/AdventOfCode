package fuel

// Calculate the fuel necessary for the given module weight
func FuelFor(weight int) int {
    if weight < 6 {
        return 0
    }
    return (weight / 3) - 2
}
