package grains

import (
	"errors"
)

func Square(number int) (uint64, error) {
	if number <= 0 || number > 64 {
		return 0, errors.New("number must be > 0 and <= 64")
	}
	grains := uint64(1)
	for i := 1; i < number; i++ {
		grains *= 2
	}

	return grains, nil
}

func Total() uint64 {
	var result uint64
	for i := 1; i <= 64; i++ {
		grains, _ := Square(i)
		result += grains
	}
	return result
}
