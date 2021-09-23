// Package prime implements a routine to return the
// prime factors of a positive integer
package prime

// Factors takes in a 64bit integer x and returns the
// prime factors of x
func Factors(x int64) []int64 {
    f := []int64{}
    if x < 2 {
        return f
    }
    if x & 1 == 0 {
        for x & 1 == 0 {
            x /= 2
            f = append(f, 2)
        }
    }
    for d := int64(3); x >= d ; d += 2 {
        if x % d == 0 {
            for x % d == 0 {
                x /= d
                f = append(f, d)
            }
        }
    }
    return f
}
