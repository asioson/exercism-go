// package matrix implements routines to create a matrix and generate saddle points
package matrix

import (
	"errors"
	"strconv"
	"strings"
)

// Define the Matrix and Pair types here.
type Matrix [][]int

type Pair [2]int

// New creates a new Matrix given data in string
func New(s string) (*Matrix, error) {
	rows := strings.Split(s, "\n")
	nrows := len(rows)
	m := make(Matrix, nrows)
	for i, item := range rows {
		for _, s := range strings.Split(item, " ") {
			if v, err := strconv.Atoi(s); err == nil {
				m[i] = append(m[i], v)
			} else {
				return nil, errors.New("invalid value")
			}
		}
	}
	return &m, nil
}

// Saddle determines the set of saddle points in a Matrix object
func (m *Matrix) Saddle() []Pair {
	nrows, ncols := len(*m), 0
	if nrows > 0 {
		ncols = len((*m)[0])
	}
	maxRow := make([]int, nrows)
	for j := 0; j < nrows; j++ {
		maxRow[j] = (*m)[j][0]
		for i := 1; i < ncols; i++ {
			if (*m)[j][i] > maxRow[j] {
				maxRow[j] = (*m)[j][i]
			}
		}
	}
	minCol := make([]int, ncols)
	for i := 0; i < ncols; i++ {
		minCol[i] = (*m)[0][i]
		for j := 1; j < nrows; j++ {
			if (*m)[j][i] < minCol[i] {
				minCol[i] = (*m)[j][i]
			}
		}
	}
	pairs := []Pair{}
	for j := 0; j < nrows; j++ {
		for i := 0; i < ncols; i++ {
			if (*m)[j][i] == maxRow[j] && (*m)[j][i] == minCol[i] {
				pairs = append(pairs, Pair{j, i})
			}
		}
	}
	return pairs
}
