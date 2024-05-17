package isogram

import "unicode"

func IsIsogram(word string) bool {
	runeCount := make(map[rune]int)
	for _, r := range word {
		if r != '-' && r != ' ' {
			runeCount[unicode.ToLower(r)]++
		}
	}
	for _, v := range runeCount {
		if v > 1 {
			return false
		}
	}
	return true
}
