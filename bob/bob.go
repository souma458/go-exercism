package bob

import (
	"strings"
	"unicode"
)

func Hey(remark string) string {
	remark = strings.TrimSpace(remark)
	switch {
	case remark == "":
		return "Fine. Be that way!"
	case isAllUpperCase(remark) && isQuestion(remark):
		return "Calm down, I know what I'm doing!"
	case isQuestion(remark):
		return "Sure."
	case isAllUpperCase(remark):
		return "Whoa, chill out!"
	default:
		return "Whatever."
	}
}

func isAllUpperCase(remark string) bool {
	hasLetter := false

	for _, r := range remark {
		if unicode.IsLetter(r) {
			hasLetter = true
			if !unicode.IsUpper(r) {
				return false
			}
		}
	}

	return hasLetter
}

func isQuestion(remark string) bool {
	return len(remark) > 0 && rune(remark[len(remark)-1]) == '?'
}
