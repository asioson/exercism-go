// package binarysearch implements a search routine using
// binary search method
package binarysearch


// SearchInts takes in a sorted list and a key and then
// returns the index of the key in the list (when it exists).
// An index value of -1 is returned when the key is not found.
func SearchInts(list []int, key int) int {
    i := 0
    j := len(list) - 1
    for i <= j {
        m := (i + j) >> 1
        if list[m] == key { 
            return m 
        } else if key < list[m] {
            j = m - 1
        } else {
            i = m + 1
        }
    }
    return -1
}
