// package piglatin implements routine to transform a phrase to piglatin
package piglatin

import "strings"

// support routines

func isVowel(ch rune) bool {
    return strings.ContainsRune("AEIOUaeiou", ch)
}

func soundsVowel(s string) bool {
    if len(s) < 2 {
        return false
    }
    if s[:2] == "xr" || s[:2] == "yt" {
        return true
    }
    return false
}

func startsWithQU(s string) bool {
    if len(s) < 2 {
        return false
    }
    if s[:2] == "qu" {
        return true
    }
    return false
}

func withQU(s string) bool {
    if len(s) < 3 {
        return false
    }
    if s[1:3] == "qu" {
        return true
    }
    return false
}

// Sentence takes in a string and returns a string
// with words in piglatin
func Sentence(sentence string) string {
    pigLatin := ""
    words := strings.Fields(sentence)
    wLen := len(words)
    for i := 0; i < wLen; i++ {
        w := words[i]
        if i > 0 {
            pigLatin += " "
        }
        if soundsVowel(w) {
            pigLatin += w + "ay"
        } else if isVowel(rune(w[0])) {
            pigLatin += w + "ay"
        } else if startsWithQU(w) {
            pigLatin += w[2:] + "quay"
        } else if withQU(w) {
            pigLatin += w[3:] + w[:3] + "ay"
        } else {
            j := 0
            for j = 0; j < len(w); j++ {
                if j > 0 && w[j] == 'y' {
                    break
                }
                if isVowel(rune(w[j])) {
                    break
                }
            }
            pigLatin += w[j:] + w[:j] + "ay"
        }
    }
    return pigLatin 
}
