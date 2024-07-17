package atbash

import "unicode"

var alphabet = "abcdefghijklmnopqrstuvwxyz"
var cipherMap = buildCipherMap()

func Atbash(s string) string {
	result := ""
	spaceCtrl := 0
	for i, r := range s {
		if spaceCtrl == 5 && i != len(s)-1 {
			result += " "
			spaceCtrl = 0
		}
		addToString(&result, r, &spaceCtrl)
	}
	return result
}

func buildCipherMap() map[rune]rune {
	backwardMap := make(map[rune]rune)
	for i, char := range alphabet {
		backwardMap[char] = rune(alphabet[len(alphabet)-1-i])
	}
	return backwardMap
}

func addToString(s *string, r rune, ctrl *int) {
	if unicode.IsLetter(r) {
		*s += string(cipherMap[unicode.ToLower(r)])
		*ctrl++
	} else {
		if unicode.IsDigit(r) {
			*s += string(r)
			*ctrl++
		}
	}
}
