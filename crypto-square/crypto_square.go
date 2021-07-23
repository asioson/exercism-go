// package cryptosquare implements a routine that 
// encrypts a string using crypto square method
package cryptosquare

import "strings"
import "unicode"
import "math"

// Encode takes in a message and produces an encrypted
// message using crypto square method
func Encode(data string) string {
    data = strings.ToLower(data)
    mesg := ""
    for _, ch := range data {
        if unicode.IsDigit(ch) || unicode.IsLetter(ch) {
            mesg += string(ch)
        }
    }
    n := len(mesg)
    c := int(math.Sqrt(float64(n)))
    r := c
    if c * r < n {
        r++
        if c * r < n {
            c++
        }
    }
    list := make([]string, r)
    k := 0
    for j := 0; j < c; j++ {
        for i := 0; i < r; i++ {
            if k < n {
                list[i] += string(mesg[k])
            } else {
                list[i] += " "
            }
            k++
        }
    }
    return strings.Join(list," ") 
}
