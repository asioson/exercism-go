// package isogram implements a routine that checks
// if the input word is an isogram
package isogram

import (
    "strings"
    "unicode"
)

// IsIsogram takes a word and returns true if the
// word is an isogram else false
func IsIsogram(word string) bool {
    word = strings.ToUpper(word)
    seenLetters := ""
    for _, c := range word {
        if strings.ContainsAny(seenLetters, string(c)) { 
            return false
        }
        if unicode.IsLetter(c) {
            seenLetters += string(c)
        }
    } 
    return true
}
