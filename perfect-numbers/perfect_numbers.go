package perfect

import "errors"

// Define the Classification type here.
type Classification string

const (
	ClassificationAbundant  Classification = "abundant"
	ClassificationDeficient Classification = "deficient"
	ClassificationPerfect   Classification = "perfect"
)

var ErrOnlyPositive error = errors.New("only positive numbers are allowed")

func Classify(n int64) (Classification, error) {
	if n <= 0 {
		return "", ErrOnlyPositive
	}
	sum := aliquotSum(n)
	switch {
	case sum == n:
		return ClassificationPerfect, nil
	case sum > n:
		return ClassificationAbundant, nil
	default:
		return ClassificationDeficient, nil
	}
}

func aliquotSum(n int64) int64 {
	sum := int64(0)
	for i := int64(1); i <= n/2; i++ {
		if n%i == 0 {
			sum += i
		}
	}
	return sum
}
