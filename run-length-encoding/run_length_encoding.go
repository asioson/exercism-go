// Package encode implements routines to do run-length encoding and decoding
package encode

import (
    "fmt"
    "unicode"
)

// RunLengthEncode encodes an input message to run-length encoded string
func RunLengthEncode(mesg string) string {
    code := ""
    n := len(mesg)
    if n > 0 {
        prevChar := mesg[0] 
        count := 1 
        i := 1 
        for i < n {
            if mesg[i] == prevChar {
                count++
            } else {
                if count > 1 {
                    code += fmt.Sprintf("%d", count)
                }
                code += fmt.Sprintf("%c", prevChar)
                prevChar = mesg[i]
                count = 1
            }
            i++
        }
        if count > 1 {
            code += fmt.Sprintf("%d", count)
        }
        code += fmt.Sprintf("%c", prevChar)
    }
    return code
}


// RunLengthDecode decodes an input run-length code to its original message
func RunLengthDecode(code string) string {
    mesg := ""
    n := len(code)
    if n > 0 {
        i := 0 
        for i < n {
            ch := code[i]
            if unicode.IsDigit(rune(ch)) {
                freq := 0
                for unicode.IsDigit(rune(ch)) {
                    freq = freq * 10 + int(ch - '0')
                    i++
                    if i == n {
                        break
                    }
                    ch = code[i]
                }
                for j := 0; j < freq; j++ {
                    mesg += fmt.Sprintf("%c", ch)
                }
            } else {
                mesg += fmt.Sprintf("%c", ch)
            }
            i++
        }
    }
    return mesg
}
