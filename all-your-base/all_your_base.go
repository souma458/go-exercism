package allyourbase

import (
	"errors"
	"math"
)

func ConvertToBase(inputBase int, inputDigits []int, outputBase int) ([]int, error) {
	if inputBase < 2 {
		return nil, errors.New("input base must be >= 2")
	}

	if outputBase < 2 {
		return nil, errors.New("output base must be >= 2")
	}

	baseTen, convError := convertToBaseTen(inputBase, inputDigits)
	if convError != nil {
		return nil, convError
	}

	return convertFromBaseTen(outputBase, baseTen), nil
}

func convertToBaseTen(inputBase int, inputDigits []int) (int, error) {
	result := 0
	for i := 0; i < len(inputDigits); i++ {
		if inputDigits[i] < 0 || inputDigits[i] >= inputBase {
			return 0, errors.New("all digits must satisfy 0 <= d < input base")
		}
		result += inputDigits[i] * int(math.Pow(float64(inputBase), float64(len(inputDigits)-i-1)))
	}
	return result, nil
}

func convertFromBaseTen(outputBase int, input int) []int {
	result := []int{}
	if input == 0 {
		return append(result, 0)
	}
	for input > 0 {
		digit := input % outputBase
		result = append([]int{digit}, result...)
		input = (input - digit) / outputBase
	}
	return result
}
