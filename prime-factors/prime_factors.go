package prime

// Factors returns the prime factors of n.
func Factors(n int64) []int64 {
	var result []int64
	for div := int64(2); div*div <= n; div++ {
		for n%div == 0 {
			result = append(result, div)
			n /= div
		}
	}
	if n > 1 {
		result = append(result, n)
	}
	return result
}

// IsPrime checks if a given number is prime.
func IsPrime(value int64) bool {
	if value <= 1 {
		return false
	}
	for i := int64(2); i*i <= value; i++ {
		if value%i == 0 {
			return false
		}
	}
	return true
}
