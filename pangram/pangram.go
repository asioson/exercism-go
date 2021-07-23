// package pangram implements a routine that checks if
// a string is a pangram or not
package pangram

import "unicode"

// Pangram takes in an input string and 
// checks if it is a pangram or not
func IsPangram(input string) bool {
    tally := map[rune]int{}
    for _, ch := range input {
        ch = unicode.ToLower(ch)
        if unicode.IsLetter(ch) {
            tally[ch]++
        }
    }
    for ch := 'a'; ch <= 'z'; ch++ {
        if tally[ch] == 0 { return false }
    }
    return true
}
