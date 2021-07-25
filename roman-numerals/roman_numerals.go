// package romannumerals implement a routine to determine
// the roman numeral equivalent of a given natural number
package romannumerals

import "fmt"
import "strings"

// ToRomanNumeral takes in a decimal number and returns
// the equivalent roman numeral expression
func ToRomanNumeral(d int) (string, error) {
	if d <= 0 || d > 3000 {
		return "", fmt.Errorf("input must be in range 1 to 3000 only")
	}
    decimal := []int{
        1, 4, 5, 9, 10, 40, 50, 90, 100, 400, 500, 900, 1000,
    }
    roman := []string{
        "I","IV","V","IX","X","XL","L","XC","C","CD","D","CM","M",
    }
	r := ""
	for i := len(roman) - 1; i >= 0; i-- {
        n := d / decimal[i]
        d = d % decimal[i]
		r += strings.Repeat(roman[i], n)
	}
	return r, nil
}
