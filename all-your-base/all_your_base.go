// Package allyourbase implements base conversion of a given integer value
package allyourbase

import "fmt"

// ConvertToBase takes in a digit list (dList) in input base (inBase) 
// and returns a digit list in output base (outBase)
func ConvertToBase(inBase int, dList []int, outBase int) ([]int, error) {
    outList := []int{}
    if inBase < 2 {
        return outList, fmt.Errorf("input base must be >= 2")
    }
    if outBase < 2 {
        return outList, fmt.Errorf("output base must be >= 2")
    }
    if inBase == outBase {
        return dList, nil
    }
    if len(dList) == 0 {
        return []int{0}, nil
    }
    value := 0
    for _, d := range dList {
        if d < 0 || d >= inBase {
            return []int{}, fmt.Errorf("all digits must satisfy 0 <= d < input base")
        }
        value = value * inBase + d
    }
    for value > 0 {
        outList = append([]int{value % outBase}, outList...)
        value /= outBase
    }
    if len(outList) == 0 {
        outList = append(outList,0)
    }
    return outList, nil
}

