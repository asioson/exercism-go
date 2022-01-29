// package wordsearch implements the routines to match words in a puzzle matrix
package wordsearch

import "fmt"

// Match takes a word and puzzle including the following parameters:
//    ly, lx : number of rows and cols in puzzle
//    sy, sx : starting location in puzzle
//    dy, dx : change in movement
// and then returns the following:
//    ey, ex : ending location in the puzzle
//    res : true if there's a match, false otherwise
func Match(w string, puzzle []string, ly, lx, sy, sx, dy, dx int) (ey, ex int, res bool) {
	matched := true
	x := sx
	y := sy
	for i := 0; i < len(w); i++ {
		if w[i] != puzzle[y][x] {
			matched = false
			break
		}
		x += dx
		y += dy
		if x < 0 || x >= lx || y < 0 || y >= ly {
			matched = (i == len(w)-1)
			break
		}
	}
	if matched {
		return y - dy, x - dx, true
	}
	return 0, 0, false
}

// Solve takes a slice of words and puzzle and returns the location
// of each word in the puzzle.  Returns error when at least a word
// is not found.
func Solve(words []string, puzzle []string) (map[string][2][2]int, error) {
	loc := map[string][2][2]int{}
	ylen := len(puzzle)
	xlen := 0
	if ylen > 0 {
		xlen = len(puzzle[0])
	}
	dy := []int{0, 0, 1, -1, 1, 1, -1, -1}
	dx := []int{1, -1, 0, 0, 1, -1, 1, -1}
	for _, w := range words {
		foundWord := false
		for y := 0; y < ylen; y++ {
			for x := 0; x < xlen; x++ {
				if w[0] != puzzle[y][x] {
					continue
				}
				for i := 0; i < len(dx); i++ {
					if ey, ex, ok := Match(w, puzzle, ylen, xlen, y, x, dy[i], dx[i]); ok {
						loc[w] = [2][2]int{{x, y}, {ex, ey}}
						foundWord = true
					}
				}
			}
		}
		if !foundWord {
			return loc, fmt.Errorf("word %s not found", w)
		}
	}
	return loc, nil
}
