When it comes to fibonacci, it's always tight with performance of it, so we included concurrent methods and benchmarks as well.

go test -benchmem -run=^$ -bench ^BenchmarkFibonacciLoop$ go_exercises/factorial -v
go test -benchmem -run=^$ -bench ^BenchmarkFibonacciRecursive$ go_exercises/factorial -v
go test -benchmem -run=^$ -bench ^BenchmarkFibonacciRecursiveGoRoutine$ go_exercises/factorial -v
go test -benchmem -run=^$ -bench ^BenchmarkFibonacciRecursiveContinuousGoRoutine$ go_exercises/factorial -v





