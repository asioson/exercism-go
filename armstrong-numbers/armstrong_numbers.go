// Package armstrong implements a routine to check if
// a number is an armstrong number
package armstrong

import "fmt"

// IsNumber takes in an integer a and checks if a is
// an armstrong number
func IsNumber(a int) bool {
    if 0 <= a && a < 10 {
        return true
    }
    n := len(fmt.Sprintf("%d", a))
    sum := 0
    x := a
    for x > 0 {
        d := x % 10
        x = x / 10
        y := 1
        for i := 0; i < n; i++ {
            y *= d
        }
        sum += y
    }
    return sum == a
}
