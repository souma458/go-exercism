package encode

import (
	"strconv"
	"unicode"
)

func RunLengthEncode(input string) string {
	if len(input) == 0 {
		return ""
	}

	previous := string(input[0])
	count := 0
	result := ""
	for i := 0; i < len(input); i++ {
		current := string(input[i])
		if current == previous {
			count++
			continue
		}
		addToResult(&result, previous, count)
		previous = current
		count = 1
	}

	addToResult(&result, previous, count)
	return result
}

func RunLengthDecode(input string) string {
	result := ""
	count := ""
	for _, c := range input {
		if unicode.IsDigit(c) {
			count += string(c)
		} else {
			if count == "" {
				count = "1"
			}
			countInt, _ := strconv.Atoi(count)
			for i := 0; i < countInt; i++ {
				result += string(c)
			}
			count = ""
		}
	}
	return result
}

func addToResult(result *string, previous string, count int) {
	if count > 1 {
		*result += strconv.Itoa(count) + previous
	} else {
		*result += previous
	}
}
