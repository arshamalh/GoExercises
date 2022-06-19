package fibonacci

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var fibsq = []int{
	0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55,
	89, 144, 233, 377, 610, 987, 1597, 2584, 4181, 6765, 10946,
}

func TestFibonacciLoop(t *testing.T) {
	for i := range fibsq {
		expected := fibsq[i]
		actual := FibonacciLoop(i)
		assert.Equalf(t, expected, actual, "Expected %v, but got %v:", expected, actual)
	}
}

func TestFibonacciRecursive(t *testing.T) {
	for i := range fibsq {
		expected := fibsq[i]
		actual := FibonacciRecursive(i)
		assert.Equalf(t, expected, actual, "Expected %v, but got %v:", expected, actual)
	}
}

func TestFibonacciRecursiveGoRoutine(t *testing.T) {
	for i := range fibsq {
		expected := fibsq[i]
		actual := FibonacciRecursiveGoRoutine(i)
		assert.Equalf(t, expected, actual, "Expected %v, but got %v:", expected, actual)
	}
}

func TestFibonacciRecursiveContinuousGoRoutine(t *testing.T) {
	for i := range fibsq {
		expected := fibsq[i]
		actual := FibonacciRecursiveContinuousGoRoutine(i)
		assert.Equalf(t, expected, actual, "Expected %v, but got %v:", expected, actual)
	}
}

func BenchmarkFibonacciLoop(b *testing.B) {
	for n := 0; n < b.N; n++ {
		FibonacciLoop(15)
	}
}

func BenchmarkFibonacciRecursive(b *testing.B) {
	for n := 0; n < b.N; n++ {
		FibonacciRecursive(15)
	}
}

func BenchmarkFibonacciRecursiveGoRoutine(b *testing.B) {
	for n := 0; n < b.N; n++ {
		FibonacciRecursiveGoRoutine(15)
	}
}

func BenchmarkFibonacciRecursiveContinuousGoRoutine(b *testing.B) {
	for n := 0; n < b.N; n++ {
		FibonacciRecursiveContinuousGoRoutine(15)
	}
}
