// package transpose transposes a set of strings
package transpose

// Transpose receives a slice of strings and returns 
// a slice of string with transposed characters
func Transpose(input []string) []string {
    output := []string{}
    n := len(input)
    if n > 0 {
        maxLen := len(input[0])
        j := 0
        for j < maxLen {
            str := ""
            space := ""
            for i := 0; i < n; i++ {
                if j < len(input[i]) {
                    if len(space) > 0 {
                        str += space
                        space = ""
                    }
                    str += string(input[i][j])
                } else {
                    space += " "
                }
            }
            output = append(output, str)
            j++
        }
    }
    return output
}
