// Package prime implements a routine that returns the value of the nth prime
package prime

// Nth takes in a value n and returns the nth prime number
func Nth(n int) (int, bool) {
    if n < 1 {
        return 0, false
    }
    prime := []int{2}
    if n > 1 {
        count := 1
        for x := 3; ; x += 2 {
            xIsPrime := true
            for j := 0; j < count; j++ {
                if x % prime[j] == 0 {
                    xIsPrime = false
                    break
                }
            }
            if xIsPrime {
                prime = append(prime, x)
                count++
                if count == n {
                    break
                }
            }
        }
    }
    return prime[n-1], true
}

