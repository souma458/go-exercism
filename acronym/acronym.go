package acronym

import (
	"regexp"
	"strings"
	"unicode"
)

var regex = regexp.MustCompile(`[^a-zA-Z0-9\s-]+`)

func Abbreviate(s string) string {
	s = regex.ReplaceAllString(s, "")
	acronym := string(s[0])
	for i, r := range s {
		if i != 0 && (r == ' ' || r == '-') && i+1 < len(s) && unicode.IsLetter(rune(s[i+1])) {
			acronym += string(s[i+1])
		}
	}
	return strings.ToUpper(acronym)
}
