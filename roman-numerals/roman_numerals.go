package romannumerals

import "errors"

var letterValue = []struct {
	letter string
	value  int
}{
	{"M", 1000},
	{"CM", 900},
	{"D", 500},
	{"CD", 400},
	{"C", 100},
	{"XC", 90},
	{"L", 50},
	{"XL", 40},
	{"X", 10},
	{"IX", 9},
	{"V", 5},
	{"IV", 4},
	{"I", 1},
}

func ToRomanNumeral(input int) (string, error) {
	if !isValidNumber(input) {
		return "", errors.New("invalid input; must be greater than 0 and less than 4000")
	}
	result := ""
	for _, entry := range letterValue {
		div := input / entry.value
		for i := 0; i < div; i++ {
			result += entry.letter
		}
		input = input % entry.value
	}
	return result, nil
}

func isValidNumber(input int) bool {
	return input > 0 && input < 4000
}
