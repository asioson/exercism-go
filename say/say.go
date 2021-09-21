// package say implements routines to translate a number
// from 0 to 999,999,999,999 to words
package say

var digit = []string {
    "zero","one","two","three","four","five","six","seven","eight","nine",
}

var teen = []string {
    "","eleven","twelve","thirteen","fourteen","fifteen","sixteen","seventeen",
    "eighteen","nineteen",
}

var tens = []string {
    "", "ten", "twenty","thirty","forty","fifty","sixty","seventy","eighty",
    "ninety","hundred",
}

var units = []string {
    "", "thousand","million","billion",
}

// SayLess1000 is a helper function that takes in a number from 0 to 999
// and returns the value in english words
func SayLess1000(x int64) string {
    numWords := ""
    if x >= 1000 || x < 0 {
        return ""
    }
    if x < 10 {
        return digit[x]
    }
    h  := x / 100
    hr := x % 100
    if h > 0 {
        numWords = digit[h] + " hundred"
    }
    t  := hr / 10
    tr := hr % 10
    if t > 0 {
        if h > 0 {
            numWords += " "
        }
        if t > 1 {
            numWords += tens[t]
            if tr > 0 {
                numWords += "-" + digit[tr] 
            }
        } else if t == 1 {
            numWords += teen[tr]
        } else {
            numWords += digit[tr]
        }
    } 
    return numWords
}

// Say takes in a value from 0 to 999,999,999,999 and returns the 
// value in english words
func Say(x int64) (string, bool) {
    if x < 0 || x > 999999999999 {
        return "", false
    }
    if x < 1000 {
        return SayLess1000(x), true
    }
    u := 0
    numWords := ""
    for (x > 0) {
        y := x % 1000
        x = x / 1000
        if y > 0 {
            if len(numWords) > 0 && numWords[0] != ' ' {
                numWords = " " + numWords
            }
            if u > 0 {
                numWords = SayLess1000(y) + " " + units[u] + numWords
            } else {
                numWords = SayLess1000(y) + numWords
            }
        }
        u++
    }
    return numWords, true 
}
