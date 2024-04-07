package booking

import "time"

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
	parsedDate, _ := time.Parse("1/2/2006 15:04:05", date)
	return parsedDate
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
	parsedDate, _ := time.Parse("January 2, 2006 15:04:05", date)
	return parsedDate.Before(time.Now())
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	parsedDate, _ := time.Parse("Monday, January 2, 2006 15:04:05", date)
	return parsedDate.Hour() >= 12 && parsedDate.Hour() <= 18
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	parsedDate := Schedule(date)
	return parsedDate.Format("You have an appointment on Monday, January 2, 2006, at 15:04.")
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	return time.Date(time.Now().Year(), time.September, 15, 0, 0, 0, 0, time.UTC)
}
