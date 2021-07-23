// package luhn implements a routine that checks if
// the input string is valid using the Luhn algorithm
package luhn

import "strings"

// Valid takes in a string. If the string has characters
// other than the digits 0 to 9 or space, the string is
// immediately declared invalid.  Otherwise, the Luhn
// formula is used to check if the string is valid.
//
// Luhn's algorithm
// 1. Double every second digit, starting from the right. 
// 2. If doubling the number results in a number greater 
//    than 9 then subtract 9 from the product. 
// 3. Get the sum all of the digits.
// 4. If the sum is evenly divisible by 10, then the 
//    number is valid else invalid
//
func Valid(card string) bool {
    n := len(card)
    sum := 0
    j := 0
    for i := n-1; i >= 0; i-- {
        d := card[i]
        if strings.ContainsAny("0123456789", string(d)) {
            if j % 2 == 1 {
                t := 2 * int(d - '0')
                if t > 9 {
                    sum += (t - 9)
                } else {
                    sum += t
                }
            } else {
                sum += int(d - '0')
            }
            j++
        } else if d != ' ' {
            return false
        }
    }
    return (j > 1) && (sum % 10 == 0)
}
