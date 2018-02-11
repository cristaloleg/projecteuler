package projecteuler

import "testing"

func TestProblem002(t *testing.T) {
	res := solve002()
	t.Logf("Answer: %v", res)
}

func solve002() (res int64) {
	res, a, b := int64(0), int64(0), int64(1)
	for a < 4000000 {
		a, b = b, a+b
		if a%2 == 0 {
			res += a
		}
	}
	return res
}
