package isbn

import (
	"strconv"
	"strings"
	"unicode"
)

func IsValidISBN(isbn string) bool {
	isbn = removeDashes(isbn)
	if len(isbn) != 10 {
		return false
	}

	if !isLastCharacterValid(isbn) {
		return false
	}

	return validateIsbn(isbn)
}

func removeDashes(isbn string) string {
	return strings.ReplaceAll(isbn, "-", "")
}

func isLastCharacterValid(isbn string) bool {
	lastChar := rune(isbn[len(isbn)-1])
	return lastChar == 'X' || isCharacterValid(lastChar)
}

func isCharacterValid(isbnCharacter rune) bool {
	return unicode.IsDigit(isbnCharacter)
}

func validateIsbn(isbn string) bool {
	count := 10
	total := 0
	for count > 0 {
		isbnCharacter := rune(isbn[10-count])
		if count > 1 && !isCharacterValid(isbnCharacter) {
			return false
		}

		isbnNumber, _ := strconv.Atoi(string(isbn[10-count]))

		if isbnCharacter == 'X' {
			isbnNumber = 10
		}

		total += isbnNumber * count
		count--
	}
	return total%11 == 0
}
