// package clock implements routines that creates a clock and add/subtract
// minutes from it 
package clock

import "strconv"

type Clock struct {
    hour    int
    minute  int
}

// New creates a new clock given hour and minute values
func New(hour, minute int) Clock {
    if minute < 0 {
        deltaHours := (-minute) / 60 
        hour -= deltaHours + 1 
        minute = 60 - ((-minute) % 60)
    }
    if minute >= 60 {
        deltaHours := minute / 60
        hour += deltaHours
        hour %= 24
        minute %= 60
    }
    if hour < 0 {
        hour = 24 - ((-hour) % 24)
    }
    hour %= 24
    return Clock{hour: hour, minute: minute}
}

// String returns the time in string format hh:mm 
func (c Clock) String() string {
    timeStr := ""
    if c.hour < 10 {
        timeStr = "0"
    }
    timeStr += strconv.Itoa(c.hour) + ":"
    if c.minute < 10 {
        timeStr += "0"
    }
    timeStr += strconv.Itoa(c.minute)
    return timeStr
}

// Add minutes from the clock object
func (c Clock) Add(minute int) Clock {
    return New(c.hour, c.minute + minute)
}

// Subtract minutes from the clock object
func (c Clock) Subtract(minute int) Clock {
    return New(c.hour, c.minute - minute)
}
