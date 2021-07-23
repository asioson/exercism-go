// package lsproduct implements a routine to compute the
// largest product of a sequence of digits
package lsproduct

import "fmt"
import "unicode"

// LargestSeriesProduct taken in a string of digits and calculates
// the largest product of a contiguous substring of digits
func LargestSeriesProduct(digitList string, k int) (int, error) {
    getProduct := func(s string) int {
        p := 1
        for _, d := range s {
            p *= int(d) - '0'
        }
        return p
    }
    if k < 0 { return 0, fmt.Errorf("invalid span") }
    maxProd := 0
    n := len(digitList)
    if n < k { 
        return 0, fmt.Errorf("incorrect span") 
    } else {
        for _, d := range digitList {
            if !unicode.IsDigit(d) {
                return 0, fmt.Errorf("invalid digit")
            }
        }
        for i := 0; i < n - k + 1; i++ {
            x := getProduct(digitList[i:i+k])
            if x > maxProd { maxProd = x }
        }
    }
    return maxProd, nil
}

