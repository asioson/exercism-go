// Package acronym implements a routine that returns an acronym of a given string
package acronym

import "unicode"
import "strings"

// Abbreviate takes the first letter of words in string s to form an acronym
// and returns the acronym with all letters in uppercase
func Abbreviate(s string) string {
	insideWord := false
	acronym := ""
	for _, c := range s {
		if unicode.IsLetter(c) || c == '\'' {
			if (! insideWord) {
				insideWord = true
				acronym = acronym + string(c)
			}
		} else {
			insideWord = false
		}
	}
	return strings.ToUpper(acronym)
}
