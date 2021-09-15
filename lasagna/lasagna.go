package lasagna

func OvenTime() int {
    return 40
}

func RemainingOvenTime(t int) int {
    return OvenTime() - t
}

func PreparationTime(layers int) int {
    return 2 * layers
}

func ElapsedTime(layers int, t int) int {
    return t + PreparationTime(layers)
}
