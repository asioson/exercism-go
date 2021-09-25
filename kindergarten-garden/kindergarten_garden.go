// Package kindergarten implements routines to track
// the plants of kinder students 
package kindergarten

import (
    "fmt"
    "sort"
    "strings"
)

// Garden maps each student to a list of their plants
type Garden map[string][]string


// NewGarden takes in a diagram and a list of children
// and returns a new kindergarten garden.
func NewGarden(diagram string, children []string) (*Garden, error) {
    n := len(children)
	sortedChildren := make([]string, n)
	copy(sortedChildren, children)
	sort.Strings(sortedChildren)
    for i := 0; i < n - 1; i++ {
		if sortedChildren[i] == sortedChildren[i+1] {
			return nil, fmt.Errorf("%s has duplicate", sortedChildren[i])
		}
	}
	if len(diagram) == 0 || diagram[0] != '\n' {
		return nil, fmt.Errorf("invalid garden format")
	}
    row := strings.Split(diagram[1:], "\n") 
    if len(row) != 2 {
		return nil, fmt.Errorf("invalid garden format")
	}
    if len(row[0]) != len(row[1]) {
	    return nil, fmt.Errorf("uneven garden")
    }
    plantName := map[rune]string{
	    'C': "clover",
	    'G': "grass",
	    'R': "radishes",
	    'V': "violets",
    }
	garden := Garden{}
    for i := 0; i < 2; i++ {
        if 2 * n != len(row[i]) {
			return nil, fmt.Errorf("each child must have 2 plants per row")
        }
		for j, code := range row[i] {
			name, ok := plantName[code]
			if !ok {
				return nil, fmt.Errorf("unknown plant")
			}
			child := sortedChildren[j/2]
			plantList, ok := garden[child]
			if !ok {
				plantList = []string{}
			}
			garden[child] = append(plantList, name)
		}
	}
	return &garden, nil
}

// Plants returns a map where each child maps to a list of plants
func (g Garden) Plants(child string) ([]string, bool) {
	plantList, ok := g[child]
	return plantList, ok 
}
