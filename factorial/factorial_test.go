package factorial_test

import (
	"go_exercises/factorial"
	"testing"
)

func TestFactorial(t *testing.T) {
	tests := []struct {
		num      uint
		expected uint
	}{
		{0, 1},
		{1, 1},
		{5, 120},
		{6, 720},
		{10, 3628800},
		{16, 20922789888000},
		{20, 2432902008176640000},
	}
	for _, test := range tests {
		actual := factorial.Factorial(test.num)
		if actual != test.expected {
			t.Errorf("Factorial(%d) = %d, expected %d", test.num, actual, test.expected)
		}
	}
}

func BenchmarkFactorial(b *testing.B) {
	for i := 0; i < b.N; i++ {
		factorial.Factorial(30)
	}
}

func BenchmarkFactorialRecursive(b *testing.B) {
	for i := 0; i < b.N; i++ {
		factorial.FactorialRecursive(30)
	}
}
