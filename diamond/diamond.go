// package diamond implements a routine that generates a diamond kata
package diamond

import "fmt"

// Gen generates a diamond kata given an input letter
func Gen(char byte) (string, error) {
	if int(char) < int('A') || int(char) > int('Z') {
		return "", fmt.Errorf("invalid character")
	}
	baseChar := int('A')
	output := ""
	offset := 0
	halfSize := int(char) - baseChar
	for i := 0; i <= halfSize; i++ {
		strLetter := string(rune(baseChar + i))
		for j := 0; j < halfSize-offset; j++ {
			output += " "
		}
		output += strLetter
		if offset > 0 {
			for j := 0; j < 2*offset-1; j++ {
				output += " "
			}
			output += strLetter
		}
		for j := 0; j < halfSize-offset; j++ {
			output += " "
		}
		output += "\n"
		offset++
	}
	for i := halfSize - 1; i >= 0; i-- {
		strLetter := string(rune(baseChar + i))
		for j := 0; j < halfSize-offset+2; j++ {
			output += " "
		}
		output += strLetter
		if offset > 2 {
			for j := 0; j < 2*offset-5; j++ {
				output += " "
			}
			output += strLetter
		}
		for j := 0; j < halfSize-offset+2; j++ {
			output += " "
		}
		output += "\n"
		offset--
	}
	return output, nil
}
