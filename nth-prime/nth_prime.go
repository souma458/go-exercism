package prime

import "errors"

// Nth returns the nth prime number. An error must be returned if the nth prime number can't be calculated ('n' is equal or less than zero)
func Nth(n int) (int, error) {
	if n <= 0 {
		return 0, errors.New("n must be >= 0")
	}

	primes := []int{2}
	for i := 3; len(primes) < n; i++ {
		if isPrime(i, primes) {
			primes = append(primes, i)
		}
	}
	return primes[len(primes)-1], nil
}

func isPrime(n int, primes []int) bool {
	for _, p := range primes {
		if n%p == 0 {
			return false
		}
	}
	return true
}
