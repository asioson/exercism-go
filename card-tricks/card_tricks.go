// Package cards implements routines to manipulate stack of items
package cards

// GetItem retrieves an item from a slice at given position. The second return value indicates whether
// a the given index existed in the slice or not.
func GetItem(slice []int, index int) (int, bool) {
    n := len(slice)
    if index < 0 || index >= n {
        return 0, false
    }
    return slice[index], true
}

// SetItem writes an item to a slice at given position overwriting an existing value.
// If the index is out of range it is be appended.
func SetItem(slice []int, index, value int) []int {
    n := len(slice)
    if index < 0 || index >= n {
        return append(slice,value)
    }
    slice[index] = value
    return slice
}

// PrefilledSlice creates a slice of given length and prefills it with the given value.
func PrefilledSlice(value, length int) []int {
    if length <= 0 {
        return nil
    }
    slice := []int{}
    for i := 0; i < length; i++ {
        slice = append(slice, value)
    }
    return slice
}

// RemoveItem removes an item from a slice by modifying the existing slice.
func RemoveItem(slice []int, index int) []int {
    n := len(slice)
    if index < 0 || index >= n {
        return slice
    }
    newSlice := append(slice[:index], slice[index+1:]...)
    return newSlice
}
