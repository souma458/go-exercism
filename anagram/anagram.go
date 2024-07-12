package anagram

import (
	"slices"
	"unicode"
)

func Detect(subject string, candidates []string) []string {
	subjectRunes := []rune(subject)
	toLowerCase(subjectRunes)

	anagrams := make([]string, 0)
	for _, c := range candidates {
		cRunes := []rune(c)
		toLowerCase(cRunes)
		if !slices.Equal(subjectRunes, cRunes) {
			slices.Sort(subjectRunes)
			slices.Sort(cRunes)
			if slices.Equal(subjectRunes, cRunes) {
				anagrams = append(anagrams, c)
			}
		}
	}

	return anagrams
}

func toLowerCase(runes []rune) {
	for i, r := range runes {
		runes[i] = unicode.ToLower(r)
	}
}
