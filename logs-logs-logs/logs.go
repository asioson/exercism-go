// Package logs implements the routines to process log-lines
package logs

import (
    "strings"
    "unicode"
    "unicode/utf8"
)

// Message extracts the message from the provided log line.
func Message(line string) string {
    i := 0
    j := len(line) - 1
    for i < j {
        if line[i] == ':' {
            i++
            break
        } else {
            i++
        }
    }
    return strings.TrimFunc(line[i:], unicode.IsSpace)
}

// MessageLen counts the amount of characters (runes) in the message of the log line.
func MessageLen(line string) int {
	return utf8.RuneCountInString(Message(line)) 
}

// LogLevel extracts the log level string from the provided log line.
func LogLevel(line string) string {
    if line[1] == 'E' {
        return "error"
    } else if line[1] == 'I' {
        return "info"
    } else if line[1] == 'W' {
        return "warning"
    }
	return ""
}

// Reformat reformats the log line in the format `message (logLevel)`.
func Reformat(line string) string {
	return Message(line) + " (" + LogLevel(line) + ")"
}
