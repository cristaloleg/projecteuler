package projecteuler

import (
	"strconv"
	"testing"
)

var globalBool bool

func TestIsPrime(t *testing.T) {
	tests := []struct {
		num  int
		want bool
	}{
		{2, true},
		{3, true},
		{4, false},
		{5, true},
	}
	for _, test := range tests {
		t.Run(strconv.Itoa(test.num), func(t *testing.T) {
			got := IsPrime(test.num)
			if got != test.want {
				t.Fatalf("want %v, got %v", test.want, got)
			}
		})
	}
}

func BenchmarkIsPrime(b *testing.B) {
	for i := 0; i < b.N; i++ {
		globalBool = IsPrime(1000000009)
	}
}
