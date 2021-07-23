// package raindrops implements a routine that converts
// a number to raindrop sound (as string)
package raindrops

import "strconv"

// Convert takes an integer and converts it to raindrop
// sound (as string)
func Convert(x int) string {
    mustReturnDigits := true;
    sound := ""
    if x % 3 == 0 {
        sound += "Pling"
        mustReturnDigits = false;
    }
    if x % 5 == 0 {
        sound += "Plang"
        mustReturnDigits = false;
    }
    if x % 7 == 0 {
        sound += "Plong"
        mustReturnDigits = false;
    }
    if mustReturnDigits {
       return strconv.Itoa(x)
    } else {
       return sound;
    }
}

