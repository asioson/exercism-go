// package grains implements routines to compute the number
// of grains in each square of a chessboard and total number
// of grains overall.
package grains

import "errors"

type errorString struct {
    s string
}

func (e *errorString) Error() string {
    return e.s
}

func New(text string) error {
    return &errorString{text}
}
// Square takes an integer n and returns the number of
// grains on square n.  Valid n values are 1 to 64 only
func Square(n int) (uint64,error) {
    square := uint64(1)
    if (n > 64) || (n <= 0) {
        return 0, errors.New("input can't be larger than 64") 
    } else {
        for i := 2; i <= n; i++ {
            square *= 2
        }
    }
    return square, nil 
}

// Total returns the total number of grains on a chessboard
func Total() uint64 {
    sum := uint64(0)
    for i := 1; i <= 64; i++ {
        v, _ := Square(i)
        sum += v 
    }
    return sum
}
