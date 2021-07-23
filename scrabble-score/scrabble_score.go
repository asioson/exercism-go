// package scrabble implements a routine to compute
// scrabble score given a word
package scrabble

import "strings"

// Score takes in a word and returns the Scabble score
func Score(word string) int {
    points := map[string] int {
        "AEIOULNRST": 1,
        "DG":         2,
        "BCMP":       3,
        "FHVWY":      4,
        "K":          5,
        "JX":         8,
        "QZ":        10,
    }
    word = strings.ToUpper(word)
    totalScore := 0
    for _, c := range word {
        for key, p := range points {
            if strings.ContainsAny(key, string(c)) {
                totalScore += p
            }
        }
    }
    return totalScore
}
