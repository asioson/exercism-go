// package brackets implements a routine that checks
// if an expression has balanced brackets
package brackets

import "strings"

// Bracket takes in an expression and checks if it
// has balanced brackets
func Bracket(expr string) bool {
    stack := []rune{} 
    for _, ch := range expr {
        if strings.Contains("{([", string(ch)) {
            stack = append(stack, ch)
        } else if strings.Contains("})]", string(ch)) {
            n := len(stack) - 1
            if n < 0 { return false }
            if ch == ']' {
                if stack[n] != '[' { return false }
            } else if ch == ')' {
                if stack[n] != '(' { return false }
            } else {
                if stack[n] != '{' { return false }
            }
            stack = stack[:n]
        }
    }
    return len(stack) == 0
}
