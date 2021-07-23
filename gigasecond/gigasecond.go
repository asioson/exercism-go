// package gigasecond implements a routine that returns time t plus 1 gigasecond
package gigasecond

// import path for the time package from the standard library
import "time"

// AddGigasecond returns time t plus 1 gigasecond
func AddGigasecond(t time.Time) time.Time {
	return t.Add(time.Second * time.Duration(1000000000))
}
