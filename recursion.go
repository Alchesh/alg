package main

// Fakk - factorial
func Fakk(val int) int {
	if val < 2 {
		return 1
	}

	return val * Fakk(val-1)
}

// Sum - n first sum
func Sum(val int) int {

	if val == 1 {
		return 1
	}

	return val + Sum(val-1)
}

// Fib - fibonacci
func Fib(val int) int {

	if val == 1 || val == 2 {
		return 1
	}

	return Fib(val-1) + Fib(val-2)
}

// SumMas - sum of array elements
func SumMas(val []int) int {

	n := len(val)

	if n == 0 {
		return 0
	}

	return val[n-1] + SumMas(val[0:n-1])
}
