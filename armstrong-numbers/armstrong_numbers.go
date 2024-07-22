package armstrong

import (
	"math"
	"strconv"
)

func IsNumber(n int) bool {
	s := strconv.Itoa(n)
	total := 0
	pow := len(s)
	for _, v := range s {
		intValue, _ := strconv.Atoi(string(v))
		total += int(math.Pow(float64(intValue), float64(pow)))
	}
	return n == total
}
