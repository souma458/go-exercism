package wordcount

import (
	"regexp"
	"strings"
)

type Frequency map[string]int

func WordCount(phrase string) Frequency {
	words := strings.FieldsFunc(phrase, func(r rune) bool {
		return r == ':' || r == '!' || r == '?' || r == ',' || r == ';' || r == '\t' || r == '\n' || r == ' ' || r == '.'
	})
	frequency := make(map[string]int, 0)
	for _, w := range words {
		w = strings.Trim(w, "'")
		if len(w) > 0 && isValidWord(w) {
			frequency[strings.TrimSpace(strings.ToLower(w))]++
		}
	}

	return frequency
}

func isValidWord(word string) bool {
	return regexp.MustCompile("^[a-zA-Z0-9'`-]*$").MatchString(word)
}
