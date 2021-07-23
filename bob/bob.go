// Package bob implements a routine that returns a specific string
// given a phrase/question

package bob

import "unicode"
import "strings"

// Hey returns the specific strings given certain phrases.
//
// phrases                     return value
// --------------------------  --------------------------
// A yell (in all caps)        'Whoa, chill out!'
// A yelled question           'Calm down, I know what I'm doing!'
// A normal question           'Sure.'
// Empty phrase                'Fine. Be that way!'
// Everything else             'Whatever.'
//
func Hey(remark string) string {
    remark = strings.TrimSpace(remark)
    if len(remark) == 0 {
        return "Fine. Be that way!" 
    }
    allCaps := true
    letterExists := false
    for _, r := range remark {
        if unicode.IsLetter(r) {
            letterExists = true
            if !unicode.IsUpper(r) {
                allCaps = false
                break
            }
        }
    }
    if letterExists {
        if remark[len(remark)-1] == '?' {
            if allCaps {
                return "Calm down, I know what I'm doing!"
            } else {
                return "Sure."
            }
        } else if allCaps {
            return "Whoa, chill out!"
        }
    } else if remark[len(remark)-1] == '?' {
        return "Sure."
    }
    return "Whatever."
}
