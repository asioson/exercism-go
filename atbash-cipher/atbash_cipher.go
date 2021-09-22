// Package atbash implements the atbash cipher
package atbash

import "unicode"

// Atbash takes in a message and returns the encrypted message
// using atbash cipher
func Atbash(mesg string) string {
    code := ""
    count := 0
    for _, c := range mesg {
        if unicode.IsLetter(c) {
            if count == 5 {
                code += " "
                count = 0
            }
            if unicode.IsUpper(c) {
                c = rune(25 - int(c - 'A')) + 'a'
            } else {
                c = rune(25 - int(c - 'a')) + 'a'
            }
            code += string(c)
            count++
        } else if unicode.IsDigit(c) {
            if count == 5 {
                code += " "
                count = 0
            }
            code += string(c)
            count++
        }
    }
    return code
}
