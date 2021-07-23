// package wordcount implements routines that
// returns a frequency table of words found in 
// a phrase
package wordcount

import "strings"
import "unicode"

type Frequency map[string]int 


// WordCount takes in a phrase and returns a frequency
// table of words found in the phrase
func WordCount(phrase string) Frequency {
    freq := Frequency{}
    insideWord := false
    phrase = strings.ToLower(phrase)
    word := ""
    for _, ch := range phrase {
        if insideWord { 
            if unicode.IsLetter(ch) || unicode.IsDigit(ch) || ch == '\'' {
                word += string(ch) 
            } else {
                n := len(word)-1
                if n > 0 && word[n] == '\'' { word = word[:n] }
                freq[word]++
                insideWord = false 
                word = ""
            }
        } else {
            if unicode.IsLetter(ch) || unicode.IsDigit(ch) {
                word += string(ch)
                insideWord = true
            }
        }
    }
    if insideWord { 
        n := len(word)-1
        if n > 0 && word[n] == '\'' { word = word[:n] }
        freq[word]++
    }
    return freq
}

