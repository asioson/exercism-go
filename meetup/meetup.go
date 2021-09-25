// Package meetup implements routine to return calendar day given
// week schedule
package meetup

import "time"

type WeekSchedule int

const First  = 0
const Second = 7
const Third  = 14
const Fourth = 21
const Teenth = 12
const Last   = -7

// Day takes a specified week schedule, weekday, month, and year
// and then returns the calendar day.
func Day(s WeekSchedule, d time.Weekday, m time.Month, y int) int {
    if s == Last {
        m++
    }
    t := time.Date(y, m, int(s) + 1, 12, 0, 0, 0, time.UTC)
    offset := int(d + 7 - t.Weekday()) % 7
    t = t.AddDate(0, 0, offset)
    return t.Day()
}
