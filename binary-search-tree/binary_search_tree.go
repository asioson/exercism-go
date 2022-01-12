// Package binarysearchtree implements binary search tree routines
package binarysearchtree

type SearchTreeData struct {
	left  *SearchTreeData
	data  int
	right *SearchTreeData
}

// NewBst creates and returns a new SearchTreeData.
func NewBst(i int) SearchTreeData {
    return SearchTreeData{data:i}
}

// InsertValue is a helper function to implement correct
// insertion of values in BST
func InsertValue(t *SearchTreeData, i int) *SearchTreeData {
    if t != nil {
        if i <= (*t).data {
            (*t).left = InsertValue((*t).left, i)
        } else {
            (*t).right = InsertValue((*t).right, i)
        }
    } else {
        t = new(SearchTreeData)
        *t = NewBst(i)
    }
    return t
}

// Insert inserts an int into the SearchTreeData.
// Inserts happen based on the rules of a BinarySearchTree
func (std *SearchTreeData) Insert(i int) {
    InsertValue(std, i)
}

// MapString returns the ordered contents of SearchTreeData as a []string.
// The values are in increasing order starting with the lowest int value.
// SearchTreeData that has the numbers [1,3,7,5] added will return the
// []string ["1", "3", "5", "7"].
func (std *SearchTreeData) MapString(f func(int) string) (result []string) {
    if std != nil {
        result = (std.left).MapString(f)
        result = append(result, f(std.data))
        temp := (std.right).MapString(f)
        result = append(result, temp ...)
        return result
    }
	return []string{}
}

// MapInt returns the ordered contents of SearchTreeData as an []int.
// The values are in increasing order starting with the lowest int value.
// SearchTreeData that has the numbers [1,3,7,5] added will return the
// []int [1,3,5,7].
func (std *SearchTreeData) MapInt(f func(int) int) (result []int) {
    if std != nil {
        result = (std.left).MapInt(f)
        result = append(result, f(std.data))
        temp := (std.right).MapInt(f)
        result = append(result, temp ...)
        return result
    }
	return []int{}
}
