package rotationalcipher

import (
	"slices"
	"strings"
)

var alphabet = []string{
	"a",
	"b",
	"c",
	"d",
	"e",
	"f",
	"g",
	"h",
	"i",
	"j",
	"k",
	"l",
	"m",
	"n",
	"o",
	"p",
	"q",
	"r",
	"s",
	"t",
	"u",
	"v",
	"w",
	"x",
	"y",
	"z",
}

func RotationalCipher(plain string, shiftKey int) string {
	nrOfAlphabetLetters := len(alphabet)

	for i := 0; i < len(plain); i++ {
		current := string(plain[i])
		currentAlphabetIndex := slices.Index(alphabet, strings.ToLower(current))

		if currentAlphabetIndex != -1 {
			shift := slices.Index(alphabet, strings.ToLower(current)) + shiftKey

			if shift > nrOfAlphabetLetters-1 {
				shift -= nrOfAlphabetLetters
			}

			replacement := string(alphabet[shift])
			if strings.ToUpper(current) == current {
				replacement = strings.ToUpper(replacement)
			}

			plain = plain[:i] + replacement + plain[i+1:]
		}
	}

	return plain
}
