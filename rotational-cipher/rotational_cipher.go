// Package rotationalcipher implements routine to do ROTn
package rotationalcipher

import "unicode"

// RotationalCipher takes in a mesg and key n and returns
// the cipher subjected to ROTn
func RotationalCipher(mesg string, n int) string {
    code := ""
    for _, c := range mesg {
        if unicode.IsLetter(c) {
            if unicode.IsUpper(c) {
                c = rune((int(c - 'A') + n) % 26) + 'A'
            } else {
                c = rune((int(c - 'a') + n) % 26) + 'a'
            }
        }
        code += string(c)
    }
    return code
}
