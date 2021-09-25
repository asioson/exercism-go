// Package techpalace implements routines for tech palace store app
package techpalace

import "strings"

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
    return "Welcome to the Tech Palace, " + strings.ToUpper(customer)
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
    starLine := strings.Repeat("*", numStarsPerLine)
    return starLine + "\n" + welcomeMsg + "\n" + starLine 
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
    return strings.Trim(strings.Split(oldMsg,"\n")[1]," *\n")
}
