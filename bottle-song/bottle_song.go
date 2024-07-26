package bottlesong

import (
	"fmt"
	"strings"
)

var numbers = []string{"One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten"}

func Recite(startBottles, takeDown int) []string {
	result := make([]string, 0)
	for takeDown > 0 {
		bottleWord := "bottles"
		if startBottles == 1 {
			bottleWord = "bottle"
		}

		firstAndSecondLine := fmt.Sprintf("%s green %s hanging on the wall,", numbers[startBottles-1], bottleWord)
		thirdLine := "And if one green bottle should accidentally fall,"
		var finalLine string

		switch startBottles {
		case 1:
			finalLine = "There'll be no green bottles hanging on the wall."
		case 2:
			finalLine = fmt.Sprintf("There'll be %s green bottle hanging on the wall.", strings.ToLower(numbers[startBottles-2]))
		default:
			finalLine = fmt.Sprintf("There'll be %s green bottles hanging on the wall.", strings.ToLower(numbers[startBottles-2]))
		}

		result = append(result, firstAndSecondLine, firstAndSecondLine, thirdLine, finalLine)

		takeDown--
		startBottles--

		if takeDown > 0 {
			result = append(result, "")
		}
	}
	return result
}
