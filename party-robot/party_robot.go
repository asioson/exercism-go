// Package partyrobot implements routines that returns
// party messages
package partyrobot

import "fmt"

// Welcome greets a person by name.
func Welcome(name string) string {
    return "Welcome to my party, " + name + "!"
}

// HappyBirthday wishes happy birthday to the birthday person and stands out their age.
func HappyBirthday(name string, age int) string {
    return fmt.Sprintf("Happy birthday %s! You are now %d years old!", name, age)
}

// AssignTable assigns a table to each guest.
func AssignTable(name string, table int, neighbour, direction string, distance float64) string {
    message := "Welcome to my party, %s!\n" 
    message += "You have been assigned to table %X. "
    message += "Your table is %s, exactly %.1f meters from here.\n" 
    message += "You will be sitting next to %s."
    return fmt.Sprintf(message, name, table, direction, distance, neighbour)
}
