// package minesweeper implements routines to process minesweeper puzzles
package minesweeper

import "fmt"

// IsValid checks if the board is valid
func (b Board) IsValid() bool {
	ylen := len(b)
	xlen := len(b[0])
	for x := 1; x < xlen-1; x++ {
		if b[0][x] != '-' || b[ylen-1][x] != '-' {
			return false
		}
	}
	if b[0][0] != '+' || b[0][xlen-1] != '+' {
		return false
	}
	if b[ylen-1][0] != '+' || b[ylen-1][xlen-1] != '+' {
		return false
	}
	for y := 1; y < ylen-1; y++ {
		if len(b[y]) != xlen {
			return false
		}
		if b[y][0] != '|' || b[y][xlen-1] != '|' {
			return false
		}
	}
	for y := 1; y < ylen-1; y++ {
		for x := 1; x < xlen-1; x++ {
			if b[y][x] != '*' && b[y][x] != ' ' {
				return false
			}
		}
	}
	return true
}

// Update increments the value of a valid location
// on the board
func (b Board) Update(y, x int) {
	if b[y][x] == ' ' {
		b[y][x] = '1'
	} else if b[y][x] >= '1' && b[y][x] < '8' {
		b[y][x]++
	}
}

// Count process the board and shows the number of mines
// in blank locations
func (b Board) Count() error {
	ylen := len(b)
	xlen := len(b[0])
	dy := []int{-1, 0, 1, 1, 0, -1, -1, 1}
	dx := []int{-1, -1, -1, 1, 1, 1, 0, 0}
	if !b.IsValid() {
		return fmt.Errorf("invalid board")
	}
	for y := 1; y < ylen-1; y++ {
		for x := 1; x < xlen-1; x++ {
			if b[y][x] == '*' {
				for i := 0; i < len(dx); i++ {
					b.Update(y+dy[i], x+dx[i])
				}
			}
		}
	}
	return nil
}
