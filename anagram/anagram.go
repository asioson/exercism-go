// package anagram implements routines to detect
// which among candidate words are anagrams of
// a specific word
package anagram

import "sort"
import "strings"

// SortLetters takes in a word and returns a string where the
// letters in the word are sorted
func SortLetters(word string) string {
    word = strings.ToLower(word)
    sList := []string{}
    for _, ch := range word { sList = append(sList, string(ch)) }
    sort.Strings(sList)
    sWord := ""
    for _, ch := range sList {
        sWord += string(ch)
    }
    return sWord 
}

// Detect takes in a word and candidate strings and
// returns a list of anagrams of the word
func Detect(word string, candidates []string) []string {
    anagrams := []string{}
    sWord := SortLetters(word)
    for _, w := range candidates {
        if strings.ToLower(word) == strings.ToLower(w) {
            continue
        }
        if sWord == SortLetters(w) {
            anagrams = append(anagrams, w)
        }
    }
    return anagrams
}
