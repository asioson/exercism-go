// package strain implements methods of sample collections
package strain

// sample collections
type Ints []int
type Lists [][]int
type Strings []string

// Keep returns subset of a collection that satisfies mustKeep condition
func (collection Ints) Keep(mustKeep func(int) bool) Ints {
    result := Ints(nil)
    for _, item := range collection {
        if mustKeep(item) {
            result = append(result,item)
        }
    }
    return result
}

// Discard returns subset of a collection that does not satisfy
// the mustDiscard condition
func (collection Ints) Discard(mustDiscard func(int) bool) Ints {
    result := Ints(nil)
    for _, item := range collection {
        if !mustDiscard(item) {
            result = append(result,item)
        }
    }
    return result
}

// Keep returns subset of a collection that satisfies mustKeep condition
func (collection Lists) Keep(mustKeep func([]int) bool) Lists {
    result := Lists(nil)
    for _, item := range collection {
        if mustKeep(item) {
            result = append(result,item)
        }
    }
    return result
}

// Keep returns subset of a collection that satisfies mustKeep condition
func (collection Strings) Keep(mustKeep func(string) bool) Strings {
    result := Strings(nil)
    for _, item := range collection {
        if mustKeep(item) {
            result = append(result,item)
        }
    }
    return result
}
