package projecteuler

import "testing"

func TestProblemXXX(t *testing.T) {
	args := [2]struct {
		want  int64
		limit int
	}{
		{21, 6},
	}

	for i, v := range args {
		limit := v.limit
		want := v.want
		if res := solveXXX(limit); res != want {
			t.Errorf("Test %v: %v != %v", i, res, want)
		}
	}

	limit := 1000
	res := solveXXX(limit)
	t.Logf("Answer: %v", res)
}

func solveXXX(n int) (res int64) {
	for i := 0; i <= n; i++ {
		res += int64(i)
	}
	return res
}
