package leap

// Determines if a given year is a leap year
func IsLeapYear(year int) bool {
	if isDivisibleByOneHundred(year) {
		return isDivisibleByFourHundred(year)
	} else {
		return isDivisibleByFour(year)
	}
}

func isDivisibleByFour(year int) bool {
	return year%4 == 0
}

func isDivisibleByOneHundred(year int) bool {
	return year%100 == 0
}

func isDivisibleByFourHundred(year int) bool {
	return year%400 == 0
}
