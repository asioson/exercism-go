// Package collatzconjecture implements a routine to compute the
// number of steps before n becomes 1
package collatzconjecture

import "fmt"

// CollatzConjecture takes in a positive integer n and returns the
// number of steps before the collatz algorithm ends with 1
func CollatzConjecture(n int) (int, error) {
    if n < 1 {
        return 0, fmt.Errorf("n must be positive integer")
    }
    steps := 0
    for n != 1 {
        if n & 1 == 1 {
            n = n * 3 + 1
        } else {
            n = n >> 1
        }
        steps++
    }
    return steps, nil
}

