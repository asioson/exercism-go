// package matrix implements routines that manages a Matrix object
package matrix

import "strings"
import "fmt"
import "strconv"

type Matrix [][]int

// New creates a new Matrix object given data in an input string
func New(data string) (Matrix, error) {
    lines := strings.Split(data,"\n")
    nRows := len(lines)
    nCols := len(strings.Split(lines[0]," "))
    m := make(Matrix, nRows)
    for j := 0; j < nRows; j++ {
        row := strings.Split(strings.Trim(lines[j]," ")," ")
        if len(row) != nCols {
            return nil, fmt.Errorf("New: wrong number of values")
        }
        m[j] = make([]int, nCols)
        for i := 0; i < nCols; i++ {
            v, err := strconv.Atoi(row[i])
            if err != nil {
                return nil, err
            }
            m[j][i] = v
        }
    }
    return m, nil
}

// Rows function returns the matrix data in a separate
// matrix structure where data are stored in row-major
func (m Matrix) Rows() [][]int {
    if m == nil {
        return nil
    }
    nRows := len(m)
    nCols := len(m[0])
    rowData := make(Matrix, nRows)
    for j := 0; j < nRows; j++ {
        rowData[j] = make([]int, nCols)
        for i := 0; i < nCols; i++ {
            rowData[j][i] = m[j][i]
        }
    }
    return rowData
}

// Cols function returns the matrix data in a separate
// matrix structure where data are stored in col-major
func (m Matrix) Cols() [][]int {
    if m == nil {
        return nil
    }
    nRows := len(m)
    nCols := len(m[0])
    colData := make(Matrix, nCols)
    for i := 0; i < nCols; i++ {
        colData[i] = make([]int, nRows)
        for j := 0; j < nRows; j++ {
            colData[i][j] = m[j][i]
        }
    }
    return colData
}

// Set function sets the matrix value at row r
// and column c to value v
func (m Matrix) Set(r, c, v int) bool {
    nRows := len(m)
    nCols := len(m[0])
    if r < 0 || r >= nRows {
        return false
    }
    if c < 0 || c >= nCols {
        return false
    }
    m[r][c] = v 
    return true
}
