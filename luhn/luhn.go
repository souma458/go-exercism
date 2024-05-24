package luhn

import (
	"strconv"
	"strings"
	"unicode"
)

func Valid(id string) bool {
	id = clearSpaces(id)
	return isLengthValid(id) && isDigitsOnly(id) && luhn(id)
}

func isLengthValid(id string) bool {
	return len(id) > 1
}

func clearSpaces(s string) string {
	return strings.ReplaceAll(s, " ", "")
}

func isDigitsOnly(s string) bool {
	for _, r := range s {
		if !unicode.IsDigit(r) {
			return false
		}
	}
	return true
}

func luhn(s string) bool {
	total := 0
	secondDigit := false

	for i := len(s) - 1; i >= 0; i-- {
		v, _ := strconv.Atoi(string(s[i]))
		if secondDigit {
			v = 2 * v
			if v > 9 {
				v = v - 9
			}
		}
		total += v
		secondDigit = !secondDigit
	}

	return total%10 == 0
}
