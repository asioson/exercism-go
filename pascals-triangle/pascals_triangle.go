// package pascal implements a routine that generates a pascal triangle
package pascal

// Triangle takes in an integer n and produce a pascal traingle of
// size n
func Triangle(n int) (result [][]int) {
    pascal := [][]int{}
    if n == 1 {
        pascal = append(pascal, []int{1})
    } else if n > 1 {
        pascal = Triangle(n-1)
        newRow := []int{1}
        for i := 0; i < n-2; i++ {
            newRow = append(newRow, pascal[n-2][i]+pascal[n-2][i+1])
        }
        newRow = append(newRow,1)
        pascal = append(pascal, newRow)
    }
    return pascal
}
