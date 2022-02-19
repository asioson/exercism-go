// package connect implements routine to determine the winner of Hex game
package connect

import (
	"fmt"
	"strings"
)

// ResultOf takes in a state of game and returns the winner
func ResultOf(lines []string) (string, error) {
	board := []string{}
	for _, line := range lines {
		board = append(board, strings.ReplaceAll(line, " ", ""))
	}
	xmax := 0
	ymax := len(board)
	if ymax > 0 {
		xmax = len(board[0])
	} else {
		return "", fmt.Errorf("invalid board")
	}
	for y := 1; y < ymax; y++ {
		if xmax != len(board[y]) {
			return "", fmt.Errorf("invalid board")
		}
	}
	for y := 0; y < ymax; y++ {
		if board[y][0] == 'X' && ReachedEnd(board, y, 0, 'X') {
			return "X", nil
		}
	}
	for x := 0; x < xmax; x++ {
		if board[0][x] == 'O' && ReachedEnd(board, 0, x, 'O') {
			return "O", nil
		}
	}
	return "", nil
}

// ReachedEnd takes in a valid board, starting point (y,x) and token ch
// and determines if ch connects to the correct side
func ReachedEnd(b []string, y, x int, ch rune) bool {
	S := [][2]int{{y, x}}
	ymax := len(b)
	xmax := len(b[0])
	visited := make([][]bool, ymax)
	for y := 0; y < ymax; y++ {
		visited[y] = make([]bool, xmax)
	}
	dy := []int{1, 0, -1, 0, 1, -1}
	dx := []int{0, 1, 0, -1, -1, 1}
	for len(S) > 0 {
		ty, tx := S[0][0], S[0][1]
		visited[ty][tx] = true
		S = S[1:]
		switch {
		case ch == 'O' && ty == ymax-1:
			return true
		case ch == 'X' && tx == xmax-1:
			return true
		}
		for i := 0; i < len(dx); i++ {
			ny := ty + dy[i]
			nx := tx + dx[i]
			if ny < 0 || ny >= ymax || nx < 0 || nx >= xmax {
				continue
			}
			if rune(b[ny][nx]) == ch && !visited[ny][nx] {
				S = append(S, [2]int{ny, nx})
			}
		}

	}
	return false
}
