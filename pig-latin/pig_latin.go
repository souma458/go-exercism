package piglatin

import (
	"strings"
	"unicode"
)

func Sentence(sentence string) string {
	words := strings.Split(sentence, " ")
	result := ""

	for _, word := range words {
		result += translateEnglishToPig(word) + " "
	}
	return strings.TrimSpace(result)
}

func translateEnglishToPig(word string) string {
	prefixQu, okQu := startsWithZeroOrMoreConsoantsFollowedByQu(word)
	prefixY, okY := startsWithOneOrMoreConsoantsFollowedByY(word)
	switch {
	case isVowel(rune(word[0])) || startsWithXr(word[0:2]) || startsWithYt(word[0:2]):
		return word + "ay"
	case okQu:
		return word[len(prefixQu):] + prefixQu + "ay"
	case okY:
		return word[len(prefixY):] + prefixY + "ay"
	default:
		return word[len(findStartingConsoants(word)):] + findStartingConsoants(word) + "ay"
	}
}

func isVowel(r rune) bool {
	return unicode.ToLower(r) == 'a' || unicode.ToLower(r) == 'e' || unicode.ToLower(r) == 'i' || unicode.ToLower(r) == 'o' || unicode.ToLower(r) == 'u'
}

func startsWithXr(firstTwo string) bool {
	return strings.ToLower(firstTwo) == "xr"
}

func startsWithYt(firstTwo string) bool {
	return strings.ToLower(firstTwo) == "yt"
}

func startsWithZeroOrMoreConsoantsFollowedByQu(s string) (string, bool) {
	for i, r := range s {
		if isVowel(r) {
			return "", false
		}
		if r == 'q' && i+2 < len(s) && s[i+1] == 'u' {
			return s[:i+2], true
		}
	}
	return "", false
}

func startsWithOneOrMoreConsoantsFollowedByY(s string) (string, bool) {
	for i, r := range s {
		if isVowel(r) {
			return "", false
		}
		if i+1 < len(s) && s[i+1] == 'y' {
			return s[:i+1], true
		}
	}
	return "", false
}

func findStartingConsoants(s string) string {
	for i, r := range s {
		if isVowel(r) {
			return s[:i]
		}
	}
	return s
}
