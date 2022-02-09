// package ocr implements routines to recognize digits in a given data sequence
package ocr

import (
	"fmt"
	"strings"
)

// recognizeDigit returns the recognized digit value given a data string
func recognizeDigit(s string) int {
	switch s {
	case " _ | ||_|":
		return 0
	case "     |  |":
		return 1
	case " _  _||_ ":
		return 2
	case " _  _| _|":
		return 3
	case "   |_|  |":
		return 4
	case " _ |_  _|":
		return 5
	case " _ |_ |_|":
		return 6
	case " _   |  |":
		return 7
	case " _ |_||_|":
		return 8
	case " _ |_| _|":
		return 9
	}
	return -1
}

// Recognize returns the digits recognized in a given data string
func Recognize(s string) []string {
	data := strings.Split(s, "\n")
	rows := len(data)
	result := []string{}
	for j := 0; j < rows; j++ {
		n := len(data[j]) / 3
		if n > 0 {
			readVal := ""
			for i := 0; i < n; i++ {
				digitStr := ""
				for r := 0; r < 3; r++ {
					digitStr += data[j+r][i*3 : i*3+3]
				}
				v := recognizeDigit(digitStr)
				if v >= 0 {
					readVal += fmt.Sprintf("%d", v)
				} else {
					readVal += "?"
				}
			}
			j += 3
			result = append(result, readVal)
		}
	}
	return result
}
