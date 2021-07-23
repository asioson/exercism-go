// package triangle implements a routine that checks if 
// three input numbers can be side lengths of a triangle
package triangle

import "math"
import "sort"

type Kind int

const NaT = 0  // not a triangle
const Equ = 1  // equilateral
const Iso = 2  // isosceles
const Sca = 3  // scalene

func valid(f float64) bool {
    return !math.IsNaN(f) && !math.IsInf(f, 0) && f > 0.0
}

// KindFromSides takes in three integers a, b, c, and 
// determines if those can form the side lengths of
// a triangle
func KindFromSides(a, b, c float64) Kind {
    if valid(a) && valid(b) && valid(c) {
	    sides := []float64{a, b, c}
	    sort.Float64s(sides)
	    a, b, c = sides[0], sides[1], sides[2]

	    if c > a+b {
		    return NaT
	    } else if a == b && a == c {
		    return Equ
	    } else if a == b || b == c {
		    return Iso
	    }
	    return Sca
    }
    return NaT
}
