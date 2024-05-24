package diffsquares

func SquareOfSum(n int) int {
	return Sum(n) * Sum(n)
}

func Sum(n int) int {
	if n == 0 {
		return 0
	}
	return n + Sum(n-1)
}

func SumOfSquares(n int) int {
	if n == 1 {
		return n
	}
	return n*n + SumOfSquares(n-1)
}

func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
