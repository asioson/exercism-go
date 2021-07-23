// package space implements a routine to compute the age on a specific
// planet given age in Earth seconds

package space

type Planet string

// Age takes in age in seconds and planet name and computes the
// age in earth-years on the planet indicated
func Age(ageInSec float64, name Planet) float64 {
    period := map[Planet] float64 {
        "Mercury": 0.2408467,
        "Venus":   0.61519726,
        "Earth":   1.0,
        "Mars":    1.8808158,
        "Jupiter": 11.862615,
        "Saturn":  29.447498,
        "Uranus":  84.016846,
        "Neptune": 164.79132,
    }
    return ageInSec / (period[name] * 31557600)
}
