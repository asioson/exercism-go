// package sieve implements a routine to determine a list 
// of primes
package sieve

// Sieve takes an integer n and returns an array of
// primes within the range [1,n]
func Sieve(n int) []int {
    primes := []int{}
    if n > 1 {
        composite := make([]bool,n+1)
        composite[0] = true
        composite[1] = true
        for d := 2; d <= n; d++ {
            if ! composite[d] {
                for i := d + d; i <= n; i += d {
                    composite[i] = true
                }
            }
        }
        for i := 2; i <= n; i++ {
            if ! composite[i] {
                primes = append(primes, i)
            }
        }
    }
    return primes
}
