// package pythagorean implements the routines 
// Range(min,max) and Sum(p)
package pythagorean

type Triplet [3]int

// Range takes in min and max and returns a list of 
// pythagorean triples (a,b,c) such that 
// min <= a, b, c <= max
func Range(min, max int) []Triplet {
    list := []Triplet{}
    for a := min; a <= max-2; a++ {
        for b := a + 1; b <= max-1; b++ {
            for c := b + 1; c <= max; c++ {
                if a * a + b * b == c * c {
                    list = append(list, Triplet{a,b,c})
                }
            }
        }
    }
    return list
}

// Sum takes in an integer p and returns a list
// of pythagorean triples (a,b,c) such that
// a + b + c == p
func Sum(p int) []Triplet {
    result := []Triplet{}
    list := Range(1,p)
    for _, x := range list {
        if x[0] + x[1] + x[2] == p {
            result = append(result, x)
        }
    }
    return result
}
