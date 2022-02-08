// package rectangles implements routine to count rectangles in a given diagram
package rectangles

// Count returns the number of rectangles in a given diagram
func Count(diagram []string) int {
	rows, cols := len(diagram), 0
	if rows > 0 {
		cols = len(diagram[0])
	}
	count := 0
	for i := 0; i < rows-1; i++ {
		for j := 0; j < cols-1; j++ {
			if diagram[i][j] == '+' {
				for k := j + 1; k < cols; k++ {
					if diagram[i][k] == '+' {
						for l := i + 1; l < rows; l++ {
							if diagram[l][j] == '+' && diagram[l][k] == '+' {
								count++
							} else if diagram[l][j] != '+' && diagram[l][j] != '|' {
								break
							} else if diagram[l][k] != '+' && diagram[l][k] != '|' {
								break
							}
						}
					} else if diagram[i][k] != '+' && diagram[i][k] != '-' {
						break
					}
				}
			}
		}
	}
	return count
}
