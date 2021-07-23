// package summultiples implements a routine that
// computes the sum of multiples of elements in 
// a list
package summultiples

// SumMultiples takes in an integer n and variable 
// list of divisors.  The function computes the sum
// of integers in the range (0,n) that are multiples
// of the divisors 
func SumMultiples(n int, args ... int) int {
    sum := 0
    for i := 1; i < n; i++ {
        for _, d := range args {
            if d == 0 { continue }
            if i % d == 0 {
                sum += i
                break
            }
        }
    }
    return sum
}

