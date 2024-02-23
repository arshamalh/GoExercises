package fibonacci

func FibonacciLoop(num int) int {
	if num <= 1 {
		return num
	}
	n2, n1 := 0, 1
	for i := 2; i < num; i++ {
		n2, n1 = n1, n1+n2
	}
	return n2 + n1
}

func FibonacciRecursive(n int) int {
	if n < 2 {
		return n
	}
	return FibonacciRecursive(n-1) + FibonacciRecursive(n-2)
}

var memoized = make(map[int]int)

func FibonacciRecursiveOptimized(n int) int {
	if n < 2 {
		return n
	}
	if val, ok := memoized[n]; ok {
		return val
	}
	return FibonacciRecursive(n-1) + FibonacciRecursive(n-2)
}

// ** Solving the Fibonacci Sequence with Go Routines ** //
// It may be hard for you, so don't hesitate to see the solution if you get stuck.

func FibonacciRecursiveGoRoutine(n int) int {
	panic("Implement me")
}

func FibonacciRecursiveContinuousGoRoutine(n int) int {
	panic("Implement me")
}
