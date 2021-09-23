// Package listops implements basic list operations
package listops

type IntList []int
type binFunc func(x, y int) int
type predFunc func(n int) bool
type unaryFunc func(x int) int


// Length returns the length of list
func (list IntList) Length() int {
	return len(list)
}

// Foldl takes a function bf, a list, and an initial
// accumulator and fold each list item into the 
// accumulator from left to right.
func (list IntList) Foldl(bf binFunc, ac int) int {
	for _, item := range list {
		ac = bf(ac, item)
	}
	return ac
}

// Foldr takes a function bf, a list, and an initial
// accumulator and fold each list item into the 
// accumulator from right to left.
func (list IntList) Foldr(bf binFunc, ac int) int {
    n := list.Length()
	for i := n - 1; i >= 0; i-- {
		ac = bf(list[i], ac)
	}
	return ac
}

// Filter takes a function pf and a list 
// and extracts items from the list that
// satisfy the predicate pf.
func (list IntList) Filter(pf predFunc) IntList {
    result := IntList{}
	for _, item := range list {
		if pf(item) {
			result = append(result, item)
		}
	}
	return result
}

// Map takes a function uf and a list
// and applies uf to each item of the list
// to produce elements of another list.
func (list IntList) Map(uf unaryFunc) IntList {
	result := IntList{}
	for _, item := range list {
		result = append(result, uf(item))
	}
	return result
}

// Reverse takes a list and produce a
// another list with elements in reverse
// order.
func (list IntList) Reverse() IntList {
    n := list.Length()
	result := make(IntList, n)
	for i, item := range list {
		result[n - 1 - i] = item
	}
	return result 
}

// Append takes a list and appends to it
// the elements of another list.
func (list IntList) Append(l IntList) IntList {
	result := IntList{}
	result = append(result, list...)
	result = append(result, l...)
	return result
}

// Concat takes a list and list of lists and
// produce a flattened list containing the
// all elements of the lists.
func (list IntList) Concat(listSeq []IntList) IntList {
	result := IntList{}
	result = append(result, list...)
	for _, l := range listSeq {
		result = append(result, l...)
	}
	return result
}

