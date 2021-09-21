// Package booking implements routines needed for appointment
// scheduler for a beauty salon
package booking

import "fmt"
import "time"

// Schedule returns a time.Time from a string containing a date
func Schedule(date string) time.Time {
    var m, d, y, hr, mn, sc int
    fmt.Sscanf(date,"%d/%d/%d %d:%d:%d", &m, &d, &y, &hr, &mn, &sc)
    return time.Date(y, time.Month(m), d, hr, mn, sc, 0, time.UTC)
}

// HasPassed returns whether a date has passed
func HasPassed(date string) bool {
    var d, y, hr, mn, sc int
    var m string
    var toMonth = map[string]time.Month {
        "January":   time.January,
        "February":  time.February,
        "March":     time.March,
        "April":     time.April,
        "May":       time.May,
        "June":      time.June,
        "July":      time.July,
        "August":    time.August,
        "September": time.September,
        "October":   time.October,
        "November":  time.November,
        "December":  time.December,
    }
    fmt.Sscanf(date,"%s %d, %d %d:%d:%d", &m, &d, &y, &hr, &mn, &sc)
    thisDate := time.Date(y, toMonth[m], d, hr, mn, sc, 0, time.UTC)
    return thisDate.Before(time.Now())
}

// IsAfternoonAppointment returns whether a time is in the afternoon
func IsAfternoonAppointment(date string) bool {
    var d, y, hr, mn, sc int
    var wd, m string
    fmt.Sscanf(date,"%s %s %d, %d %d:%d:%d", &wd, &m, &d, &y, &hr, &mn, &sc)
    return (hr >= 12) && (hr < 18)
}

// Description returns a formatted string of the appointment time
func Description(date string) string {
    message := "You have an appointment on %s, %s %d, %d, at %d:%d."
    t := Schedule(date)
    return fmt.Sprintf(message, t.Weekday(), t.Month(), t.Day(), t.Year(), t.Hour(), t.Minute()) 
}

// AnniversaryDate returns a Time with this year's anniversary
func AnniversaryDate() time.Time {
    return time.Date(time.Now().Year(), time.Month(9), 15, 0, 0, 0, 0, time.UTC)
}
