package projecteuler

import (
	"math"
	"math/big"
)

type Int = big.Int

func DivisorsCount(x int) int {
	if x == 1 {
		return 1
	}
	divisor := 0
	n := x
	for i := 1; i < n; i++ {
		if x%i == 0 {
			divisor += 2
			if i*i == x {
				divisor--
			}
			x = x / i
		}
	}
	return divisor
}

func Divisors(x int) []int {
	if x == 1 {
		return []int{1}
	}
	divisor := []int{}
	n := x
	for i := 1; i < n; i++ {
		if x%i == 0 {
			// divisor += 2
			if i*i == x {
				// divisor--
			}
			x = x / i
		}
	}
	return divisor
}

func Min(xs ...int) int {
	min := xs[0]
	for _, x := range xs {
		if x < min {
			min = x
		}
	}
	return min
}

func Max(xs ...int) int {
	max := xs[0]
	for _, x := range xs {
		if x > max {
			max = x
		}
	}
	return max
}

func Sum(xs ...int) int {
	sum := 0
	for _, x := range xs {
		sum += x
	}
	return sum
}

func Factorial(n int) *Int {
	x := new(big.Int)
	x.MulRange(1, int64(n))
	return x
}

func IsPrime(n int) bool {
	if n < 3 {
		return n == 2
	}
	if n%2 == 0 {
		return false
	}
	sqrt := int(math.Sqrt(float64(n))) + 1
	for i := 3; i < sqrt; i += 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func GCD(x, y int) int {
	for y != 0 {
		x, y = y, x%y
	}
	return x
}

func LCM(x, y int) int {
	return x / GCD(x, y) * y
}

func Pow(a, b int) int {
	p := 1
	for b > 0 {
		if b&1 != 0 {
			p *= a
		}
		b >>= 1
		a *= a
	}
	return p
}
