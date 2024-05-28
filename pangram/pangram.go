package pangram

import "unicode"

func IsPangram(input string) bool {
	runeCountMap := make(map[rune]int, 0)
	for _, r := range input {
		if unicode.IsLetter(r) {
			runeCountMap[unicode.ToLower(r)]++
		}
	}
	return len(runeCountMap) == 26
}
