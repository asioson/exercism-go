// package diffsquares implements three routines to
// compute the difference of square of sum and sum of squares
// of integers from 1 to n
package diffsquares


// SquareOfSum takes an integer n and returns the
// square of the sum 1 + 2 + ... + n
func SquareOfSum(n int) int {
    s := (n * (n + 1)) / 2
    return s * s
}

// SumOfSquares takes an integer n and returns the
// sum 1^2 + 2^2 + ... + n^2
func SumOfSquares(n int) int {
    return (n * (n + 1) * (n + n + 1)) / 6
}

// Difference takens an integer n and returns the 
// difference of square of sum and sum of squares
// of integers from 1 to n
func Difference(n int) int {
    return SquareOfSum(n) - SumOfSquares(n);
}
