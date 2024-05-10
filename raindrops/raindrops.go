package raindrops

import (
	"strconv"
)

func Convert(number int) string {
	result := ""
	if isDivisible(number, 3) {
		result += "Pling"
	}
	if isDivisible(number, 5) {
		result += "Plang"
	}
	if isDivisible(number, 7) {
		result += "Plong"
	}
	if len(result) == 0 {
		result += strconv.Itoa(number)
	}
	return result
}

func isDivisible(number int, divisor int) bool {
	return number%divisor == 0
}
