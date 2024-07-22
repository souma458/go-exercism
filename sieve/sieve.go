package sieve

// Sieve returns all prime numbers less than or equal to the limit.
func Sieve(limit int) []int {
	if limit < 2 {
		return nil
	}

	// create a slice to hold the numbers up to the limit
	numbers := make([]bool, limit+1)
	for i := 2; i*i <= limit; i++ {
		if !numbers[i] {
			for j := i * i; j <= limit; j += i {
				numbers[j] = true
			}
		}
	}

	// create a slice to hold the primes
	primes := []int{}
	for i := 2; i <= limit; i++ {
		if !numbers[i] {
			primes = append(primes, i)
		}
	}

	return primes
}
