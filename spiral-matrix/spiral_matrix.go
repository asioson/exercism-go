// package spiralmatrix implements routine to generate a spiral matrix
package spiralmatrix

const right, down, left, up = 0, 1, 2, 3

// SpiralMatrix generates a spiral matrix of given size
func SpiralMatrix(size int) [][]int {
	m := make([][]int, size)
	for i := 0; i < size; i++ {
		m[i] = make([]int, size)
	}
	xmin, ymin, xmax, ymax := 0, 0, size-1, size-1
	x, y := xmin, ymin
	direction := right
	for v := 1; v <= size*size; v++ {
		m[y][x] = v
		switch direction {
		case right:
			x++
		case down:
			y++
		case left:
			x--
		case up:
			y--
		}
		switch {
		case direction == right && x == xmax:
			direction = down
			ymin++
		case direction == down && y == ymax:
			direction = left
			xmax--
		case direction == left && x == xmin:
			direction = up
			ymax--
		case direction == up && y == ymin:
			direction = right
			xmin++
		}
	}
	return m
}
