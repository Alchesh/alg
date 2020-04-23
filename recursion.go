package main

import (
	"math"
)

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

// Fib - fibonacci (not 0!)
// 1 1 2 3 5 8 13
func Fib(val int) int {

	if val <= 2 {
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

// PowerN - num**pow CPU = O(n)
func PowerN(num, pow int) float32 {
	if pow == 0 {
		return 1
	}

	if pow > 0 {
		return float32(num) * PowerN(num, pow-1)
	}
	return float32(PowerN(num, pow+1)) / float32(num)

}

// PowerN - num**pow CPU = O(logn)
func PowerN2(num, pow int) int {
	if pow == 0 {
		return 1
	}

	if pow%2 == 0 {
		return int(math.Pow(float64(PowerN2(num, pow/2)), 2))
	} else {
		return num * int(math.Pow(float64(PowerN2(num, (pow-1)/2)), 2))
	}
}

//SlowAdd - add by increment
func SlowAdd(a, b int) int {
	if a == 0 {
		return b
	}
	if b == 0 {
		return a
	}

	return SlowAdd(a-1, b) + 1
}
