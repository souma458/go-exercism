package scrabble

import "unicode"

var valueLetterMap = map[int][]rune{
	1:  {'a', 'e', 'i', 'o', 'u', 'l', 'n', 'r', 's', 't'},
	2:  {'d', 'g'},
	3:  {'b', 'c', 'm', 'p'},
	4:  {'f', 'h', 'v', 'w', 'y'},
	5:  {'k'},
	8:  {'j', 'x'},
	10: {'q', 'z'},
}

var letterValueMap = createLetterValueMap(valueLetterMap)

func Score(word string) int {
	score := 0
	for _, r := range word {
		score += letterValueMap[unicode.ToLower(r)]
	}
	return score
}

func createLetterValueMap(inputMap map[int][]rune) map[rune]int {
	lookupMap := make(map[rune]int)
	for value, runes := range inputMap {
		for _, r := range runes {
			lookupMap[r] = value
		}
	}
	return lookupMap
}
