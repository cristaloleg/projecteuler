package projecteuler

import "testing"

func TestProblem1(t *testing.T) {
	args := [2]struct {
		param int
		want  int64
	}{
		{10, 23},
	}

	for i, v := range args {
		param := v.param
		want := v.want
		if res := solve1(param); res != want {
			t.Errorf("Test %v: %v != %v", i, res, want)
		}
	}

	res := solve1(1000)
	t.Logf("Answer: %v", res)
}

func solve1(n int) (res int64) {
	for i := 1; i < n; i++ {
		if i%3 == 0 || i%5 == 0 {
			res += int64(i)
		}
	}
	return res
}
