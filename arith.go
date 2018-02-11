package projecteuler

import (
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

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func MinOf(xs ...int) int {
	min := xs[0]
	for _, x := range xs {
		if x < min {
			min = x
		}
	}
	return min
}

func MaxOf(xs ...int) int {
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

func IsPrime(n int) bool {
	switch {
	case n <= 1:
		return false
	case n < 4:
		return true
	case n%2 == 0:
		return false
	case n < 9:
		return true
	case n%3 == 0:
		return false
	}

	for p := 5; p*p <= n; p += 6 {
		if n%p == 0 || n%(p+2) == 0 {
			return false
		}
	}
	return true
}
