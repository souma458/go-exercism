package meetup

import "time"

type WeekSchedule int

const (
	First WeekSchedule = iota
	Second
	Third
	Fourth
	Teenth
	Last
)

func Day(wSched WeekSchedule, wDay time.Weekday, month time.Month, year int) int {
	dayOfMonth := time.Date(year, month, 1, 0, 0, 0, 0, time.UTC)
	for dayOfMonth.Weekday() != wDay {
		dayOfMonth = dayOfMonth.AddDate(0, 0, 1)
	}

	switch wSched {
	case First:
		return dayOfMonth.Day()
	case Second:
		return dayOfMonth.AddDate(0, 0, 7).Day()
	case Third:
		return dayOfMonth.AddDate(0, 0, 14).Day()
	case Fourth:
		return dayOfMonth.AddDate(0, 0, 21).Day()
	case Teenth:
		switch {
		case time.Date(year, month, 13, 0, 0, 0, 0, time.UTC).Weekday() == wDay:
			return 13
		case time.Date(year, month, 14, 0, 0, 0, 0, time.UTC).Weekday() == wDay:
			return 14
		case time.Date(year, month, 15, 0, 0, 0, 0, time.UTC).Weekday() == wDay:
			return 15
		case time.Date(year, month, 16, 0, 0, 0, 0, time.UTC).Weekday() == wDay:
			return 16
		case time.Date(year, month, 17, 0, 0, 0, 0, time.UTC).Weekday() == wDay:
			return 17
		case time.Date(year, month, 18, 0, 0, 0, 0, time.UTC).Weekday() == wDay:
			return 18
		case time.Date(year, month, 19, 0, 0, 0, 0, time.UTC).Weekday() == wDay:
			return 19
		default:
			return 0
		}
	case Last:
		newDayOfMonth := dayOfMonth.AddDate(0, 0, 21)
		for newDayOfMonth.Month() == dayOfMonth.Month() {
			newDayOfMonth = dayOfMonth.AddDate(0, 0, 7)
			if newDayOfMonth.Month() == dayOfMonth.Month() {
				dayOfMonth = newDayOfMonth
			}
		}
		return dayOfMonth.Day()
	default:
		return 0
	}
}
