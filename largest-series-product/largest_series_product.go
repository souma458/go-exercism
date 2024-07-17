package lsproduct

import (
	"errors"
	"strconv"
	"unicode"
)

func LargestSeriesProduct(digits string, span int) (int64, error) {
	largest := int64(0)
	if span > len(digits) || span <= 0 {
		return 0, errors.New("span must be less or equal to digits' length and greater than 0")
	}
	for i := 0; i < len(digits)-span+1; i++ {
		product, err := calculateProduct(digits[i : i+span])
		if err != nil {
			return 0, err
		}
		if product > largest {
			largest = product
		}
	}
	return largest, nil
}

func calculateProduct(digits string) (int64, error) {
	product := int64(1)
	for i := 0; i < len(digits); i++ {
		if !unicode.IsDigit(rune(digits[i])) {
			return 0, errors.New("digits contains an invalid character")
		}
		digit, _ := strconv.Atoi(string(digits[i]))
		product *= int64(digit)
	}
	return product, nil
}
