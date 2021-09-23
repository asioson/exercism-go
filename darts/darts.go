// Package darts implements a routine to score a darts game
package darts

// Score takes in the coordinates x, y and returns a score
// according to the table:
// - - - - - - - - - - - - - - - - - - - -
//     inner circle (radius 1) : 10 pts
//     next circle (radius 5)  :  5 pts
//     next circle (radius 10) :  1 pt
//     ouside the 3rd circle   :  0 pts
// - - - - - - - - - - - - - - - - - - - -
func Score(x, y float64) int {
    sqrz := x * x + y * y
    switch {
    case sqrz > 100:
        return 0
    case sqrz <= 1:
        return 10
    case sqrz <= 25:
        return 5
    default:
        return 1
    }
}
