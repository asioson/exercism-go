// Package leap implements a routine that checks if a year
// is leap year or not
package leap

// IsLeapYear returns true if the input year is a leap year
func IsLeapYear(year int) bool {
    if year % 100 == 0 {
        return year % 400 == 0
    } else if year % 4 == 0 {
        return true
    } 
	return false
}
