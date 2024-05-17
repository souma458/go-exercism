package gigasecond

// import path for the time package from the standard library
import "time"

const gigaseconds = 1000000000

// Adds 1000000000 seconds to a specific time
func AddGigasecond(t time.Time) time.Time {
	t = t.Add(time.Second * gigaseconds)
	return t
}
