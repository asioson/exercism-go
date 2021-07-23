// package hamming implements a routine that computes the
// number of letters that differs given two strings

package hamming

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

// Distance counts the differences between strings a and b
// and returns it as a result.  This distance is called
// Hamming distance
func Distance(a, b string) (int, error) {
    if len(a) == len(b) {
        count := 0
        for i := 0; i < len(a); i++ {
            if a[i] != b[i] {
                count++
            }
        }
        return count, nil
    } else {
        return 0, errors.New("Distance: input string lengths are not equal") 
    }
}
