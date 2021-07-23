// package perfect implements a routine to
// classify natural numbers
package perfect

import "fmt"

var ErrOnlyPositive = fmt.Errorf("Only positives allowed")

type Classification int

const ClassificationDeficient = 1
const ClassificationPerfect   = 2
const ClassificationAbundant  = 3

// Classify takes in a natural number n and
// returns a classification
func Classify(n int64) (Classification, error) {
    if n <= 0 { return 0, ErrOnlyPositive }
    aliquotSum := int64(0)
    for d := int64(1); d < n; d++ {
        if n % d == 0 { aliquotSum += d }
    }
    if aliquotSum == n { return ClassificationPerfect, nil }
    if aliquotSum < n { return ClassificationDeficient, nil }
    return ClassificationAbundant, nil
}
