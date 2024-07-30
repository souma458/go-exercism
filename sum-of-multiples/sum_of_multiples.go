package summultiples

func SumMultiples(limit int, divisors ...int) int {
	energy := 0
	multiples := make(map[int]struct{})
	for _, div := range divisors {
		for _, multiple := range calculateMultiples(limit, div) {
			_, ok := multiples[multiple]
			if !ok {
				energy += multiple
				multiples[multiple] = struct{}{}
			}
		}
	}
	return energy
}

func calculateMultiples(limit, divisor int) []int {
	if divisor == 0 {
		return []int{0}
	}

	multiples := make([]int, 0)
	for i := divisor; i < limit; i += divisor {
		multiples = append(multiples, i)
	}
	return multiples
}
