package collatzconjecture

import "errors"

func CollatzConjecture(n int) (int, error) {
	if n <= 0 {
		return 0, errors.New("n must be greater than 0")
	}
	return howManySteps(n, 0), nil
}

func howManySteps(n int, steps int) int {
	if n == 1 {
		return steps
	}
	if n%2 == 0 {
		return howManySteps(n/2, steps+1)
	}
	return howManySteps(3*n+1, steps+1)
}
