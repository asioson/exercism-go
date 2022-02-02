// package change implements routines to produce the minimum change needed
// to generate a value
package change

import (
	"fmt"
)

// Change produces the change needed to generate the target value given
// the coin values available
func Change(coins []int, target int) ([]int, error) {
	return BestSum(target, coins, &map[int][]int{})
}

// BestSum does the actual work of producing the change needed to generate the
// targetSum.  The solution is recursive and uses a memo of partial solutions.
func BestSum(targetSum int, numList []int, memo *map[int][]int) ([]int, error) {
	if list, ok := (*memo)[targetSum]; ok {
		return list, nil
	}
	if targetSum == 0 {
		return []int{}, nil
	}
	if targetSum < 0 {
		return []int{}, fmt.Errorf("negative target")
	}
	var bestCombination *[]int
	for _, x := range numList {
		remainder := targetSum - x
		remainderResult, err := BestSum(remainder, numList, memo)
		if err == nil {
			if bestCombination == nil {
				bestCombination = new([]int)
				*bestCombination = append([]int{x}, remainderResult...)
			} else if len(*bestCombination) > 1+len(remainderResult) {
				*bestCombination = append([]int{x}, remainderResult...)
			}
		}
	}
	if bestCombination != nil {
		(*memo)[targetSum] = *bestCombination
		return (*memo)[targetSum], nil
	}
	return []int{}, fmt.Errorf("can't best combination for %d", targetSum)
}
