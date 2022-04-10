// package parsinglogfiles implements routines to process log texts
package parsinglogfiles

import (
	"regexp"
)

// IsValidLine returns true if the text starts with expected prefix
func IsValidLine(text string) bool {
	validPrefix := []string{
		"[TRC]", "[DBG]", "[INF]", "[WRN]", "[ERR]", "[FTL]"}
	for _, prefix := range validPrefix {
		if text[:5] == prefix {
			return true
		}
	}
	return false
}

// SplitLogLine splits an input string at <[~*=-]*> and returns
// the resulting slice of strings
func SplitLogLine(text string) []string {
	re := regexp.MustCompile(`<[~*=-]*>`)
	return re.Split(text, -1)
}

// CountQuotedPasswords return the number of instances the text
// "password" appeared (case insensitive matching)
func CountQuotedPasswords(lines []string) int {
	count := 0
	re := regexp.MustCompile(`".*[Pp][Aa][Ss][Ss][Ww][Oo][Rr][Dd].*"`)
	for _, s := range lines {
		if re.MatchString(s) {
			count++
		}
	}
	return count
}

// RemoveEndOfLineText removes `end-of-line\d*` from the input string
// and returns the resulting string
func RemoveEndOfLineText(text string) string {
	re := regexp.MustCompile(`end-of-line\d*`)
	return re.ReplaceAllString(text, "")
}

// TagWithUserName adds user information in message lines with
// matched user name
func TagWithUserName(lines []string) []string {
	updatedLines := []string{}
	re := regexp.MustCompile(`User[ ]+([a-zA-Z0-9]+)`)
	for _, s := range lines {
		if item := re.FindStringSubmatch(s); item != nil {
			s = "[USR] " + item[1] + " " + s
		}
		updatedLines = append(updatedLines, s)
	}
	return updatedLines
}
